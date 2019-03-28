package handler

import (
	"fmt"
	"sort"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/nclgh/lakawei_api/rpc"
	"github.com/nclgh/lakawei_api/utils"
	"github.com/nclgh/lakawei_rpc/client"
	"github.com/nclgh/lakawei_scaffold/rpc/device"
	"github.com/nclgh/lakawei_scaffold/rpc/common"
)

type Achievement struct {
	Id                     int64  `json:"id"`
	DeviceCode             string `json:"device_code"`
	MemberCode             string `json:"member_code"`
	DepartmentCode         string `json:"department_code"`
	AchievementDate        int64  `json:"achievement_date"`
	AchievementDescription string `json:"achievement_description"`
	AchievementRemark      string `json:"achievement_remark"`
	PatentDescription      string `json:"patent_description"`
	PaperDescription       string `json:"paper_description"`
	CompetitionDescription string `json:"competition_description"`
}

func tranAchievement(v *device.Achievement) *Achievement {
	return &Achievement{
		Id:                     v.Id,
		DeviceCode:             v.DeviceCode,
		MemberCode:             v.MemberCode,
		DepartmentCode:         v.DepartmentCode,
		AchievementDate:        v.AchievementDate.Unix(),
		AchievementDescription: v.AchievementDescription,
		AchievementRemark:      v.AchievementRemark,
		PatentDescription:      v.PatentDescription,
		PaperDescription:       v.PaperDescription,
		CompetitionDescription: v.CompetitionDescription,
	}
}

func batchTranAchievement(dm map[int64]*device.Achievement) (achis []*Achievement, devCode map[string]bool, memCode map[string]bool, deptCode map[string]bool) {
	achis = make([]*Achievement, 0)
	devCode = make(map[string]bool)
	memCode = make(map[string]bool)
	deptCode = make(map[string]bool)

	for _, v := range dm {
		achis = append(achis, tranAchievement(v))
		devCode[v.DeviceCode] = true
		memCode[v.MemberCode] = true
		deptCode[v.DepartmentCode] = true
	}
	return achis, devCode, memCode, deptCode
}

type AddAchievementForm struct {
	DeviceCode             string `form:"device_code" binding:"required"`
	MemberCode             string `form:"member_code" binding:"required"`
	DepartmentCode         string `form:"department_code" binding:"required"`
	AchievementDate        int64  `form:"achievement_date" binding:"required"`
	AchievementDescription string `form:"achievement_description" binding:"required"`
	AchievementRemark      string `form:"achievement_remark" binding:"required"`
	PatentDescription      string `form:"patent_description" binding:"required"`
	PaperDescription       string `form:"paper_description" binding:"required"`
	CompetitionDescription string `form:"competition_description" binding:"required"`
}

func AddAchievementHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "AddAchievementHandler")

	form := AddAchievementForm{}
	if ok := p.BindAndCheckForm(&form); !ok {
		return
	}
	_, err := rpc.AddAchievement(&client.RpcRequestCtx{}, device.Achievement{
		DeviceCode:             form.DeviceCode,
		MemberCode:             form.MemberCode,
		DepartmentCode:         form.DepartmentCode,
		AchievementDate:        time.Unix(form.AchievementDate, 0),
		AchievementDescription: form.AchievementDescription,
		AchievementRemark:      form.AchievementRemark,
		PatentDescription:      form.PatentDescription,
		PaperDescription:       form.PaperDescription,
		CompetitionDescription: form.CompetitionDescription,
	})
	if err != nil {
		logrus.Errorf("call ServiceAchievement.AddAchievement err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	p.Success(nil, "")
}

func DeleteAchievementHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "DeleteAchievementHandler")

	form := DeleteForm{}
	if ok := p.BindAndCheckForm(&form); !ok {
		return
	}

	_, err := rpc.DeleteAchievement(&client.RpcRequestCtx{}, form.Id)
	if err != nil {
		logrus.Errorf("call ServiceAchievement.DeleteAchievement err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	p.Success(nil, "")
}

type QueryAchievementForm struct {
	QueryPageForm
	CommonDeviceQueryForm
	MemberCode           string `form:"member_code"`
	DepartmentCode       string `form:"department_code"`
	AchievementDateStart int64  `form:"achievement_date_start"`
	AchievementDateEnd   int64  `form:"achievement_date_end"`
}

func QueryAchievementHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "QueryAchievementHandler")

	form := QueryAchievementForm{}
	if ok := p.BindAndCheckForm(&form); !ok {
		return
	}
	filter := &common.Filter{}
	// 成就时间 查询
	if form.AchievementDateStart != 0 || form.AchievementDateEnd != 0 {
		filter.TF = append(filter.TF, &common.TimeFilter{
			Field: "achievement_date",
			Start: time.Unix(form.AchievementDateStart, 0),
			End:   time.Unix(form.AchievementDateEnd, 0),
		})
	}
	// 设备相关查询
	if form.ConditionExist() {
		rspDevice, err := rpc.QueryDevice(&client.RpcRequestCtx{}, &device.Device{
			Code:           form.CommonDeviceQueryForm.Code,
			Name:           form.CommonDeviceQueryForm.Name,
			Model:          form.CommonDeviceQueryForm.Model,
			Brand:          form.CommonDeviceQueryForm.Brand,
			TagCode:        form.CommonDeviceQueryForm.TagCode,
			DepartmentCode: form.CommonDeviceQueryForm.DepartmentCode,
		}, 0, QueryAllCnt, &common.Filter{})
		if err != nil {
			logrus.Errorf("call ServiceDevice.QueryDevice err: %v", err)
			p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
			return
		}
		deviceIds := make([]string, 0)
		for _, v := range rspDevice.Devices {
			deviceIds = append(deviceIds, v.Code)
		}
		filter.IF = append(filter.IF, &common.InFilter{
			Field:     "device_code",
			Condition: deviceIds,
		})
	}

	rsp, err := rpc.QueryAchievement(&client.RpcRequestCtx{}, &device.Achievement{
		MemberCode:     form.MemberCode,
		DepartmentCode: form.DepartmentCode,
	}, form.Page-1, form.Size, filter)
	if err != nil {
		logrus.Errorf("call ServiceDevice.QueryAchievement err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	achis, devCodes, memCodes, deptCodes := batchTranAchievement(rsp.Achievements)
	sort.Slice(achis, func(i, j int) bool {
		return achis[i].Id >= achis[j].Id
	})

	data := make(map[string]interface{})

	err = GetRspMember(data, memCodes, deptCodes)
	if err != nil {
		logrus.Errorf("achievement get member err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	err = GetRspDepartment(data, deptCodes)
	if err != nil {
		logrus.Errorf("achievement get department err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	err = GetRspDevice(data, devCodes)
	if err != nil {
		logrus.Errorf("achievement get device err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}

	data["achievement"] = achis
	data["page_info"] = PageInfo{
		CurrentPage: form.Page,
		TotalPages:  rsp.TotalCount / form.Size,
		TotalCount:  rsp.TotalCount,
	}
	p.Success(data, "")
}
