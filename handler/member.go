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
	Code           string `form:"code" binding:"required"`
	Name           string `form:"name" binding:"required"`
	DepartmentCode string `form:"department_code" binding:"required"`
}

func AddMemberHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "AddMemberHandler")

	form := AddMemberForm{}
	if ok := p.BindAndCheckForm(&form); !ok {
		return
	}
	_, err := rpc.AddMember(&client.RpcRequestCtx{}, form.Code, form.Name, form.DepartmentCode)
	if err != nil {
		logrus.Errorf("call ServiceMember.AddMember err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	p.Success(nil, "")
}

func DeleteMemberHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "DeleteMemberHandler")

	form := DeleteByCodeForm{}
	if ok := p.BindAndCheckForm(&form); !ok {
		return
	}

	_, err := rpc.DeleteMember(&client.RpcRequestCtx{}, form.Code)
	if err != nil {
		logrus.Errorf("call ServiceMember.DeleteMember err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	p.Success(nil, "")
}

type QueryMemberForm struct {
	QueryPageForm
	DepartmentCode string `form:"department_code"`
	Code           string `form:"code"`
	Name           string `form:"name"`
}

func QueryMemberHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "QueryMemberHandler")

	form := QueryMemberForm{}
	if ok := p.BindAndCheckForm(&form); !ok {
		return
	}

	rsp, err := rpc.QueryMember(&client.RpcRequestCtx{}, &member.Member{
		Code:           form.Code,
		Name:           form.Name,
		DepartmentCode: form.DepartmentCode,
	}, form.Page-1, form.Size)
	if err != nil {
		logrus.Errorf("call ServiceMember.QueryMember err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	ms := make([]*member.Member, 0)
	departmentCodeMap := map[string]bool{}
	for _, v := range rsp.Members {
		ms = append(ms, v)
		departmentCodeMap[v.DepartmentCode] = true
	}

	sort.Slice(ms, func(i, j int) bool {
		return ms[i].Name < ms[j].Name
	})

	data := make(map[string]interface{})

	// department
	err = GetRspDepartment(data, departmentCodeMap)
	if err != nil {
		logrus.Errorf("member get department err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, fmt.Sprintf("%v", err))
		return
	}
	data["member"] = ms
	data["page_info"] = PageInfo{
		CurrentPage: form.Page,
		TotalPages:  rsp.TotalCount / form.Size,
		TotalCount:  rsp.TotalCount,
	}
	p.Success(data, "")
}

func GetRspMember(data map[string]interface{}, codes map[string]bool, deptCodes map[string]bool) error {
	memberCodes := make([]string, 0)
	for k, _ := range codes {
		memberCodes = append(memberCodes, k)
	}
	depRsp, err := rpc.GetMemberByCode(&client.RpcRequestCtx{}, memberCodes)
	if err != nil {
		return err
	}
	data["member"] = depRsp.Members
	for _, v := range depRsp.Members {
		deptCodes[v.DepartmentCode] = true
	}
	return nil
}
