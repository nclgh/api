package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/nclgh/lakawei_api/rpc"
	"github.com/nclgh/lakawei_rpc/client"
	"github.com/nclgh/lakawei_api/utils"
	"github.com/nclgh/lakawei_scaffold/rpc/member"
	"sort"
	"fmt"
)

type AddDepartmentForm struct {
	Code string `form:"code" binding:"required"`
	Name string `form:"name" binding:"required"`
}

func AddDepartmentHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "AddDepartmentHandler")

	form := AddDepartmentForm{}
	if ok := p.BindAndCheckForm(&form); !ok {
		return
	}
	_, err := rpc.AddDepartment(&client.RpcRequestCtx{}, form.Code, form.Name)
	if err != nil {
		logrus.Errorf("call ServiceMember.AddDepartment err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	p.Success(nil, "")
}

func DeleteDepartmentHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "DeleteDepartmentHandler")

	form := DeleteForm{}
	if ok := p.BindAndCheckForm(&form); !ok {
		return
	}

	_, err := rpc.DeleteDepartment(&client.RpcRequestCtx{}, form.Id)
	if err != nil {
		logrus.Errorf("call ServiceMember.DeleteDepartment err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	p.Success(nil, "")
}

func QueryDepartmentHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "QueryDepartmentHandler")

	form := QueryPageForm{}
	if ok := p.BindAndCheckForm(&form); !ok {
		return
	}

	rsp, err := rpc.QueryDepartment(&client.RpcRequestCtx{}, &member.Department{}, form.Page-1, form.Size)
	if err != nil {
		logrus.Errorf("call ServiceMember.QueryDepartment err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	ds := make([]*member.Department, 0)
	for _, v := range rsp.Departments {
		ds = append(ds, v)
	}
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].Name < ds[j].Name
	})

	data := make(map[string]interface{})
	data["department"] = ds
	data["page_info"] = PageInfo{
		CurrentPage: form.Page,
		TotalPages:  rsp.TotalCount / form.Size,
		TotalCount:  rsp.TotalCount,
	}
	p.Success(data, "")
}
