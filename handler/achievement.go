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
)

type Achievement struct {
	Id                     int64  `json:"id"`
	DeviceId               int64  `json:"device_id"`
	MemberId               int64  `json:"member_id"`
	DepartmemtId           int64  `json:"departmemt_id"`
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
		DeviceId:               v.DeviceId,
		MemberId:               v.MemberId,
		DepartmemtId:           v.DepartmentId,
		AchievementDate:        v.AchievementDate.Unix(),
		AchievementDescription: v.AchievementDescription,
		AchievementRemark:      v.AchievementRemark,
		PatentDescription:      v.PatentDescription,
		PaperDescription:       v.PaperDescription,
		CompetitionDescription: v.CompetitionDescription,
	}
}

func batchTranAchievement(dm map[int64]*device.Achievement) []*Achievement {
	ds := make([]*Achievement, 0)
	for _, v := range dm {
		ds = append(ds, tranAchievement(v))
	}
	return ds
}

type AddAchievementForm struct {
	DeviceId               int64  `form:"device_id" binding:"required"`
	MemberId               int64  `form:"member_id" binding:"required"`
	DepartmentId           int64  `form:"department_id" binding:"required"`
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
		DeviceId:               form.DeviceId,
		MemberId:               form.MemberId,
		DepartmentId:           form.DepartmentId,
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
	DeviceId             int64 `form:"device_id"`
	MemberId             int64 `form:"member_id"`
	DepartmentId         int64 `form:"department_id"`
	AchievementDateStart int64 `form:"achievement_date_start"`
	AchievementDateEnd   int64 `form:"achievement_date_end"`
}

func QueryAchievementHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "QueryAchievementHandler")

	form := QueryAchievementForm{}
	if ok := p.BindAndCheckForm(&form); !ok {
		return
	}
	tf := make([]*device.TimeFilter, 0)
	if form.AchievementDateStart != 0 || form.AchievementDateEnd != 0 {
		tf = append(tf, &device.TimeFilter{
			Field: "achievement_date",
			Start: time.Unix(form.AchievementDateStart, 0),
			End:   time.Unix(form.AchievementDateEnd, 0),
		})
	}
	rsp, err := rpc.QueryAchievement(&client.RpcRequestCtx{}, &device.Achievement{
		DeviceId:     form.DeviceId,
		MemberId:     form.MemberId,
		DepartmentId: form.DepartmentId,
	}, form.Page-1, form.Size, tf)
	if err != nil {
		logrus.Errorf("call ServiceAchievement.QueryAchievement err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	ds := batchTranAchievement(rsp.Achievements)
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].Id >= ds[j].Id
	})

	data := make(map[string]interface{})
	data["achievement"] = ds
	data["page_info"] = PageInfo{
		CurrentPage: form.Page,
		TotalPages:  rsp.TotalCount / form.Size,
		TotalCount:  rsp.TotalCount,
	}
	p.Success(data, "")
}
