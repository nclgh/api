package handler

import (
	"fmt"
	"sort"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/nclgh/lakawei_api/rpc"
	"github.com/nclgh/lakawei_rpc/client"
	"github.com/nclgh/lakawei_api/utils"
	"github.com/nclgh/lakawei_scaffold/rpc/member"
)

type AddMemberForm struct {
	Code         string `form:"code" binding:"required"`
	Name         string `form:"name" binding:"required"`
	DepartmentId int64  `form:"department_id" binding:"required"`
}

func AddMemberHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "AddMemberHandler")

	form := AddMemberForm{}
	if ok := p.BindAndCheckForm(&form); !ok {
		return
	}
	_, err := rpc.AddMember(&client.RpcRequestCtx{}, form.Code, form.Name, form.DepartmentId)
	if err != nil {
		logrus.Errorf("call ServiceMember.AddMember err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	p.Success(nil, "")
}

func DeleteMemberHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "DeleteMemberHandler")

	form := DeleteForm{}
	if ok := p.BindAndCheckForm(&form); !ok {
		return
	}

	_, err := rpc.DeleteMember(&client.RpcRequestCtx{}, form.Id)
	if err != nil {
		logrus.Errorf("call ServiceMember.DeleteMember err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	p.Success(nil, "")
}

type QueryMemberForm struct {
	QueryPageForm
	DepartmentId int64  `form:"department_id" `
	Code         string `form:"code"`
	Name         string `form:"name"`
}

func QueryMemberHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "QueryMemberHandler")

	form := QueryMemberForm{}
	if ok := p.BindAndCheckForm(&form); !ok {
		return
	}

	rsp, err := rpc.QueryMember(&client.RpcRequestCtx{}, &member.Member{
		Code:         form.Code,
		Name:         form.Name,
		DepartmentId: form.DepartmentId,
	}, form.Page-1, form.Size)
	if err != nil {
		logrus.Errorf("call ServiceMember.QueryMember err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	ds := make([]*member.Member, 0)
	for _, v := range rsp.Members {
		ds = append(ds, v)
	}
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].Name < ds[j].Name
	})

	data := make(map[string]interface{})
	data["member"] = ds
	data["page_info"] = PageInfo{
		CurrentPage: form.Page,
		TotalPages:  rsp.TotalCount / form.Size,
		TotalCount:  rsp.TotalCount,
	}
	p.Success(data, "")
}
