package handler

import (
	"fmt"
	"sort"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/nclgh/lakawei_api/rpc"
	"github.com/nclgh/lakawei_api/utils"
	"github.com/nclgh/lakawei_rpc/client"
	"github.com/nclgh/lakawei_scaffold/rpc/device"
)

type AddManufacturerForm struct {
	Name string `form:"name" binding:"required"`
}

func AddManufacturerHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "AddManufacturerHandler")

	form := AddManufacturerForm{}
	if ok := p.BindAndCheckForm(&form); !ok {
		return
	}
	_, err := rpc.AddManufacturer(&client.RpcRequestCtx{}, form.Name)
	if err != nil {
		logrus.Errorf("call ServiceManufacturer.AddManufacturer err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	p.Success(nil, "")
}

type QueryManufacturerForm struct {
	QueryPageForm
}

func QueryManufacturerHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "QueryManufacturerHandler")

	form := QueryManufacturerForm{}
	if ok := p.BindAndCheckForm(&form); !ok {
		return
	}

	rsp, err := rpc.QueryManufacturer(&client.RpcRequestCtx{}, &device.Manufacturer{}, form.Page-1, form.Size)
	if err != nil {
		logrus.Errorf("call ServiceManufacturer.QueryManufacturer err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	ds := make([]*device.Manufacturer, 0)
	for _, v := range rsp.Manufacturers {
		ds = append(ds, v)
	}
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].Name < ds[j].Name
	})

	data := make(map[string]interface{})
	data["manufacturer"] = ds
	data["page_info"] = PageInfo{
		CurrentPage: form.Page,
		TotalPages:  rsp.TotalCount / form.Size,
		TotalCount:  rsp.TotalCount,
	}
	p.Success(data, "")
}
