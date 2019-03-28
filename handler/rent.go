package handler

import (
	"fmt"
	"sort"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/nclgh/lakawei_api/rpc"
	"github.com/nclgh/lakawei_rpc/client"
	"github.com/nclgh/lakawei_api/utils"
	"github.com/nclgh/lakawei_scaffold/rpc/device"
	"github.com/nclgh/lakawei_scaffold/rpc/common"
	"github.com/nclgh/lakawei_scaffold/rpc/member"
)

type Rent struct {
	Id                 int64  `json:"id"`
	DeviceCode         string `json:"device_code"`
	RentStatus         int64  `json:"rent_status"`
	BorrowerMemberCode string `json:"borrower_member_code"`
	BorrowDate         int64  `json:"borrow_date"`
	ExpectReturnDate   int64  `json:"expect_return_date"`
	ReturnerMemberCode string `json:"returner_member_code"`
	RealReturnDate     int64  `json:"real_return_date"`
	BorrowRemark       string `json:"borrow_remark"`
	ReturnRemark       string `json:"return_remark"`
}

func tranRent(v *device.Rent) *Rent {
	return &Rent{
		Id:                 v.Id,
		DeviceCode:         v.DeviceCode,
		RentStatus:         v.RentStatus,
		BorrowerMemberCode: v.ReturnerMemberCode,
		BorrowDate:         v.BorrowDate.Unix(),
		BorrowRemark:       v.BorrowRemark,
		ExpectReturnDate:   v.ExpectReturnDate.Unix(),
		ReturnerMemberCode: v.ReturnerMemberCode,
		RealReturnDate:     v.RealReturnDate.Unix(),
		ReturnRemark:       v.ReturnRemark,
	}
}

func batchTranRent(rm map[int64]*device.Rent) (rs []*Rent, devCode map[string]bool, memCode map[string]bool) {
	rs = make([]*Rent, 0)
	devCode = make(map[string]bool)
	memCode = make(map[string]bool)
	for _, v := range rm {
		rs = append(rs, tranRent(v))
		devCode[v.DeviceCode] = true
		memCode[v.BorrowerMemberCode] = true
		memCode[v.ReturnerMemberCode] = true
	}
	return rs, devCode, memCode
}

type AddRentForm struct {
	DeviceCode         string `form:"device_code" binding:"required"`
	BorrowerMemberCode string `form:"borrower_member_code" binding:"required"`
	BorrowRemark       string `form:"borrow_remark" binding:"required"`
	ExpectReturnDate   int64  `form:"expect_return_date" binding:"required"`
}

func AddRentHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "AddRentHandler")

	form := AddRentForm{}
	if ok := p.BindAndCheckForm(&form); !ok {
		return
	}
	_, err := rpc.AddRent(&client.RpcRequestCtx{}, device.Rent{
		DeviceCode:         form.DeviceCode,
		BorrowerMemberCode: form.BorrowerMemberCode,
		BorrowRemark:       form.BorrowRemark,
		ExpectReturnDate:   time.Unix(form.ExpectReturnDate, 0),
	})
	if err != nil {
		logrus.Errorf("call ServiceRent.AddRent err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	p.Success(nil, "")
}

type ReturnRentForm struct {
	DeviceCode         string `form:"device_code" binding:"required"`
	ReturnerMemberCode string `form:"returner_member_code" binding:"required"`
	ReturnRemark       string `form:"return_remark" binding:"required"`
}

func ReturnRentHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "DeleteRentHandler")

	form := ReturnRentForm{}
	if ok := p.BindAndCheckForm(&form); !ok {
		return
	}

	_, err := rpc.ReturnRent(&client.RpcRequestCtx{}, form.DeviceCode, form.ReturnerMemberCode, form.ReturnRemark)
	if err != nil {
		logrus.Errorf("call ServiceRent.DeleteRent err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	p.Success(nil, "")
}

type RentMemberQueryForm struct {
	BorrowerMemberCode     string `form:"borrower_member_code"`
	BorrowerDepartmentCode string `form:"borrower_department_code"`
	ReturnerMemberCode     string `form:"returner_member_code"`
	ReturnerDepartmentCode string `form:"returner_department_code"`
}

func (r RentMemberQueryForm) ConditionBorrowerExist() bool {
	return r.BorrowerMemberCode != "" || r.BorrowerDepartmentCode != ""
}

func (r RentMemberQueryForm) ConditionReturnerExist() bool {
	return r.ReturnerMemberCode != "" || r.ReturnerDepartmentCode != ""
}

type QueryRentForm struct {
	QueryPageForm
	CommonDeviceQueryForm
	RentMemberQueryForm
	RentStatus      int64 `form:"rent_status"`
	BorrowDateStart int64 `form:"borrow_date_start"`
	BorrowDateEnd   int64 `form:"borrow_date_end"`
	ReturnDateStart int64 `form:"return_date_start"`
	ReturnDateEnd   int64 `form:"return_date_end"`
	ExpectDateStart int64 `form:"expect_date_start"`
	ExpectDateEnd   int64 `form:"expect_date_end"`
}

func QueryRentHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "QueryRentHandler")

	form := QueryRentForm{}
	if ok := p.BindAndCheckForm(&form); !ok {
		return
	}
	filter := &common.Filter{}
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
		deviceCodes := make([]string, 0)
		for _, v := range rspDevice.Devices {
			deviceCodes = append(deviceCodes, v.Code)
		}
		filter.IF = append(filter.IF, &common.InFilter{
			Field:     "device_code",
			Condition: deviceCodes,
		})
	}
	// 人员相关查询
	if form.ConditionBorrowerExist() {
		rspMember, err := rpc.QueryMember(&client.RpcRequestCtx{}, &member.Member{
			Code:           form.RentMemberQueryForm.BorrowerMemberCode,
			DepartmentCode: form.RentMemberQueryForm.BorrowerDepartmentCode,
		}, 0, QueryAllCnt)
		if err != nil {
			logrus.Errorf("call Borrower QueryMember.QueryDevice err: %v", err)
			p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
			return
		}
		borrowerCodes := make([]string, 0)
		for _, v := range rspMember.Members {
			borrowerCodes = append(borrowerCodes, v.Code)
		}
		logrus.Infof("borrowerCodes: %v", borrowerCodes)
		filter.IF = append(filter.IF, &common.InFilter{
			Field:     "borrower_member_code",
			Condition: borrowerCodes,
		})
	}
	if form.ConditionReturnerExist() {
		rspMember, err := rpc.QueryMember(&client.RpcRequestCtx{}, &member.Member{
			Code:           form.RentMemberQueryForm.ReturnerMemberCode,
			DepartmentCode: form.RentMemberQueryForm.ReturnerDepartmentCode,
		}, 0, QueryAllCnt)
		if err != nil {
			logrus.Errorf("call Returner QueryMember.QueryDevice err: %v", err)
			p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
			return
		}
		borrowerCodes := make([]string, 0)
		for _, v := range rspMember.Members {
			borrowerCodes = append(borrowerCodes, v.Code)
		}
		filter.IF = append(filter.IF, &common.InFilter{
			Field:     "returner_member_code",
			Condition: borrowerCodes,
		})
	}
	if form.BorrowDateStart != 0 || form.BorrowDateEnd != 0 {
		filter.TF = append(filter.TF, &common.TimeFilter{
			Field: "borrow_date",
			Start: time.Unix(form.BorrowDateStart, 0),
			End:   time.Unix(form.BorrowDateEnd, 0),
		})
	}
	if form.ReturnDateStart != 0 || form.ReturnDateEnd != 0 {
		filter.TF = append(filter.TF, &common.TimeFilter{
			Field: "real_return_date",
			Start: time.Unix(form.ReturnDateStart, 0),
			End:   time.Unix(form.ReturnDateEnd, 0),
		})
	}
	if form.ExpectDateStart != 0 || form.ExpectDateEnd != 0 {
		filter.TF = append(filter.TF, &common.TimeFilter{
			Field: "expect_return_date",
			Start: time.Unix(form.ExpectDateStart, 0),
			End:   time.Unix(form.ExpectDateEnd, 0),
		})
	}
	rsp, err := rpc.QueryRent(&client.RpcRequestCtx{}, &device.Rent{
		RentStatus: form.RentStatus,
	}, form.Page-1, form.Size, filter)
	if err != nil {
		logrus.Errorf("call ServiceRent.QueryRent err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	ds, devCodes, memCodes := batchTranRent(rsp.Rents)
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].Id < ds[j].Id
	})
	data := make(map[string]interface{})

	deptCode := map[string]bool{}
	err = GetRspMember(data, memCodes, deptCode)
	if err != nil {
		logrus.Errorf("rent get member err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}

	err = GetRspDepartment(data, deptCode)
	if err != nil {
		logrus.Errorf("rent get department err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}

	err = GetRspDevice(data, devCodes)
	if err != nil {
		logrus.Errorf("rent get device err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}

	data["rent"] = ds
	data["page_info"] = PageInfo{
		CurrentPage: form.Page,
		TotalPages:  rsp.TotalCount / form.Size,
		TotalCount:  rsp.TotalCount,
	}
	p.Success(data, "")
}
