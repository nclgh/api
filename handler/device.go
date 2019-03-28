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
)

type Device struct {
	Code             string `json:"code"`
	Name             string `json:"name"`
	Model            string `json:"model"`
	Brand            string `json:"brand"`
	TagCode          string `json:"tag_code"`
	DepartmentCode   string `json:"department_code"`
	ManufacturerId   int64  `json:"manufacturer_id"`
	ManufacturerDate int64  `json:"manufacturer_date"`
	RentStatus       int64  `json:"rent_status"`
	Description      string `json:"description"`
}

func tranDevice(d *device.Device) *Device {
	return &Device{
		Code:             d.Code,
		Name:             d.Name,
		Model:            d.Model,
		Brand:            d.Brand,
		TagCode:          d.TagCode,
		DepartmentCode:   d.DepartmentCode,
		ManufacturerId:   d.ManufacturerId,
		ManufacturerDate: d.ManufacturerDate.Unix(),
		RentStatus:       d.RentStatus,
		Description:      d.Description,
	}
}

func batchTranDevice(dm map[string]*device.Device) (ds []*Device, deptCode map[string]bool, manuId map[int64]bool) {
	ds = make([]*Device, 0)
	deptCode = make(map[string]bool)
	manuId = make(map[int64]bool)
	for _, v := range dm {
		ds = append(ds, tranDevice(v))
		deptCode[v.DepartmentCode] = true
		manuId[v.ManufacturerId] = true
	}
	return ds, deptCode, manuId
}

type AddDeviceForm struct {
	Code             string `form:"code" binding:"required"`
	Name             string `form:"name" binding:"required"`
	Model            string `form:"model" binding:"required"`
	Brand            string `form:"brand" binding:"required"`
	TagCode          string `form:"tag_code" binding:"required"`
	DepartmentCode   string `form:"department_code" binding:"required"`
	ManufacturerId   int64  `form:"manufacturer_id" binding:"required"`
	ManufacturerDate int64  `form:"manufacturer_date" binding:"required"`
	Description      string `form:"description" binding:"required"`
}

func AddDeviceHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "AddDeviceHandler")

	form := AddDeviceForm{}
	if ok := p.BindAndCheckForm(&form); !ok {
		return
	}
	_, err := rpc.AddDevice(&client.RpcRequestCtx{}, device.Device{
		Code:             form.Code,
		Name:             form.Name,
		Model:            form.Model,
		Brand:            form.Brand,
		TagCode:          form.TagCode,
		DepartmentCode:   form.DepartmentCode,
		ManufacturerId:   form.ManufacturerId,
		ManufacturerDate: time.Unix(form.ManufacturerDate, 0),
		Description:      form.Description,
	})
	if err != nil {
		logrus.Errorf("call ServiceDevice.AddDevice err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	p.Success(nil, "")
}

func DeleteDeviceHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "DeleteDeviceHandler")

	form := DeleteByCodeForm{}
	if ok := p.BindAndCheckForm(&form); !ok {
		return
	}

	_, err := rpc.DeleteDevice(&client.RpcRequestCtx{}, form.Code)
	if err != nil {
		logrus.Errorf("call ServiceDevice.DeleteDevice err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	p.Success(nil, "")
}

type QueryDeviceForm struct {
	QueryPageForm
	Code                  string `form:"code"`
	Name                  string `form:"name"`
	Model                 string `form:"model"`
	Brand                 string `form:"brand"`
	TagCode               string `form:"tag_code"`
	DepartmentCode        string `form:"department_code"`
	ManufacturerId        int64  `form:"manufacturer_id"`
	RentStatus            int64  `form:"rent_status"`
	ManufacturerDateStart int64  `form:"manufacturer_date_start"`
	ManufacturerDateEnd   int64  `form:"manufacturer_date_end"`
}

func QueryDeviceHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "QueryDeviceHandler")

	form := QueryDeviceForm{}
	if ok := p.BindAndCheckForm(&form); !ok {
		return
	}
	filter := &common.Filter{}
	if form.ManufacturerDateStart != 0 || form.ManufacturerDateEnd != 0 {
		filter.TF = append(filter.TF, &common.TimeFilter{
			Field: "manufacturer_date",
			Start: time.Unix(form.ManufacturerDateStart, 0),
			End:   time.Unix(form.ManufacturerDateEnd, 0),
		})
	}
	rsp, err := rpc.QueryDevice(&client.RpcRequestCtx{}, &device.Device{
		Code:           form.Code,
		Name:           form.Name,
		Model:          form.Model,
		Brand:          form.Brand,
		TagCode:        form.TagCode,
		DepartmentCode: form.DepartmentCode,
		ManufacturerId: form.ManufacturerId,
		RentStatus:     form.RentStatus,
	}, form.Page-1, form.Size, filter)
	if err != nil {
		logrus.Errorf("call ServiceDevice.QueryDevice err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	des, deCode, manuId := batchTranDevice(rsp.Devices)
	sort.Slice(des, func(i, j int) bool {
		return des[i].Name < des[j].Name
	})

	data := make(map[string]interface{})
	err = GetRspDepartment(data, deCode)
	if err != nil {
		logrus.Errorf("device get department err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	err = GetRspManufacturer(data, manuId)
	if err != nil {
		logrus.Errorf("device get manufacturer err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	data["device"] = des
	data["page_info"] = PageInfo{
		CurrentPage: form.Page,
		TotalPages:  rsp.TotalCount / form.Size,
		TotalCount:  rsp.TotalCount,
	}
	p.Success(data, "")
}

func GetRspDevice(data map[string]interface{}, codes map[string]bool) error {
	deviceCodes := make([]string, 0)
	for k, _ := range codes {
		deviceCodes = append(deviceCodes, k)
	}
	depRsp, err := rpc.GetDeviceByCode(&client.RpcRequestCtx{}, deviceCodes)
	if err != nil {
		return err
	}
	data["device"] = depRsp.Devices
	return nil
}