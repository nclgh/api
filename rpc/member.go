package rpc

import (
	"fmt"
	"github.com/nclgh/lakawei_rpc/client"
	"github.com/nclgh/lakawei_scaffold/rpc/member"
	"github.com/nclgh/lakawei_scaffold/rpc/common"
)

func CheckUserIdentity(ctx *client.RpcRequestCtx, username, password string) (*member.CheckUserIdentityResponse, error) {
	req := member.CheckUserIdentityRequest{
		UserName: username,
		Password: password,
	}
	rsp := &member.CheckUserIdentityResponse{}
	err := GetMemberClient().Call(&client.RpcRequestCtx{}, "CheckUserIdentity", req, rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Code != common.CodeSuccess {
		return nil, fmt.Errorf("call CheckUserIdentity failed. code: %v, msg: %v", rsp.Code, rsp.Msg)
	}
	return rsp, nil
}

func AddDepartment(ctx *client.RpcRequestCtx, code, name string) (*member.AddDepartmentResponse, error) {
	req := member.AddDepartmentRequest{
		Code: code,
		Name: name,
	}
	rsp := &member.AddDepartmentResponse{}
	err := GetMemberClient().Call(&client.RpcRequestCtx{}, "AddDepartment", req, rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Code != common.CodeSuccess {
		return nil, fmt.Errorf("call AddDepartment failed. code: %v, msg: %v", rsp.Code, rsp.Msg)
	}
	return rsp, nil
}

func DeleteDepartment(ctx *client.RpcRequestCtx, code string) (*member.DeleteDepartmentResponse, error) {
	req := member.DeleteDepartmentRequest{
		Code: code,
	}
	rsp := &member.DeleteDepartmentResponse{}
	err := GetMemberClient().Call(&client.RpcRequestCtx{}, "DeleteDepartment", req, rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Code != common.CodeSuccess {
		return nil, fmt.Errorf("call DeleteDepartment failed. code: %v, msg: %v", rsp.Code, rsp.Msg)
	}
	return rsp, nil
}

func GetDepartmentByCode(ctx *client.RpcRequestCtx, codes []string) (*member.GetDepartmentByCodeResponse, error) {
	req := member.GetDepartmentByCodeRequest{
		Codes: codes,
	}
	rsp := &member.GetDepartmentByCodeResponse{}
	err := GetMemberClient().Call(&client.RpcRequestCtx{}, "GetDepartmentByCode", req, rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Code != common.CodeSuccess {
		return nil, fmt.Errorf("call GetDepartmentByCode failed. code: %v, msg: %v", rsp.Code, rsp.Msg)
	}
	return rsp, nil
}

func QueryDepartment(ctx *client.RpcRequestCtx, d *member.Department, page, pageSize int64) (*member.QueryDepartmentResponse, error) {
	req := member.QueryDepartmentRequest{
		Department: d,
		Page:       page,
		PageSize:   pageSize,
	}
	rsp := &member.QueryDepartmentResponse{}
	err := GetMemberClient().Call(&client.RpcRequestCtx{}, "QueryDepartment", req, rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Code != common.CodeSuccess {
		return nil, fmt.Errorf("call QueryDepartment failed. code: %v, msg: %v", rsp.Code, rsp.Msg)
	}
	return rsp, nil
}

func AddMember(ctx *client.RpcRequestCtx, code, name string, departmentCode string) (*member.AddMemberResponse, error) {
	req := member.AddMemberRequest{
		Code:           code,
		Name:           name,
		DepartmentCode: departmentCode,
	}
	rsp := &member.AddMemberResponse{}
	err := GetMemberClient().Call(&client.RpcRequestCtx{}, "AddMember", req, rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Code != common.CodeSuccess {
		return nil, fmt.Errorf("call AddMember failed. code: %v, msg: %v", rsp.Code, rsp.Msg)
	}
	return rsp, nil
}

func DeleteMember(ctx *client.RpcRequestCtx, code string) (*member.DeleteMemberResponse, error) {
	req := member.DeleteMemberRequest{
		Code: code,
	}
	rsp := &member.DeleteMemberResponse{}
	err := GetMemberClient().Call(&client.RpcRequestCtx{}, "DeleteMember", req, rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Code != common.CodeSuccess {
		return nil, fmt.Errorf("call DeleteMember failed. code: %v, msg: %v", rsp.Code, rsp.Msg)
	}
	return rsp, nil
}

func GetMemberByCode(ctx *client.RpcRequestCtx, codes []string) (*member.GetMemberByCodeResponse, error) {
	req := member.GetMemberByCodeRequest{
		Codes: codes,
	}
	rsp := &member.GetMemberByCodeResponse{}
	err := GetMemberClient().Call(&client.RpcRequestCtx{}, "GetMemberByCode", req, rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Code != common.CodeSuccess {
		return nil, fmt.Errorf("call GetMemberByCode failed. code: %v, msg: %v", rsp.Code, rsp.Msg)
	}
	return rsp, nil
}

func QueryMember(ctx *client.RpcRequestCtx, m *member.Member, page, pageSize int64) (*member.QueryMemberResponse, error) {
	req := member.QueryMemberRequest{
		Member:   m,
		Page:     page,
		PageSize: pageSize,
	}
	rsp := &member.QueryMemberResponse{}
	err := GetMemberClient().Call(&client.RpcRequestCtx{}, "QueryMember", req, rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Code != common.CodeSuccess {
		return nil, fmt.Errorf("call QueryMember failed. code: %v, msg: %v", rsp.Code, rsp.Msg)
	}
	return rsp, nil
}
