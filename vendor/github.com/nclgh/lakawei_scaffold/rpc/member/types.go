package member

import (
	"github.com/nclgh/lakawei_scaffold/rpc/common"
)

type CheckUserIdentityRequest struct {
	UserName string
	Password string
}

type CheckUserIdentityResponse struct {
	UserId int64
	Code   common.RspCode
	Msg    string
}

// Department
type Department struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type AddDepartmentRequest struct {
	Code string
	Name string
}

type AddDepartmentResponse struct {
	Code common.RspCode
	Msg  string
}
type DeleteDepartmentRequest struct {
	Id int64
}

type DeleteDepartmentResponse struct {
	Code common.RspCode
	Msg  string
}

type GetDepartmentByIdRequest struct {
	Ids []int64
}

type GetDepartmentByIdResponse struct {
	Departments map[int64]*Department
	Code        common.RspCode
	Msg         string
}

type QueryDepartmentRequest struct {
	Department *Department
	Page       int64
	PageSize   int64
}

type QueryDepartmentResponse struct {
	Departments map[int64]*Department
	TotalCount  int64
	Code        common.RspCode
	Msg         string
}

// Member
type Member struct {
	Id           int64  `json:"id"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	DepartmentId int64  `json:"department_id"`
}

type AddMemberRequest struct {
	Code         string
	Name         string
	DepartmentId int64
}

type AddMemberResponse struct {
	Code common.RspCode
	Msg  string
}

type DeleteMemberRequest struct {
	Id int64
}

type DeleteMemberResponse struct {
	Code common.RspCode
	Msg  string
}

type GetMemberByIdRequest struct {
	Ids []int64
}

type GetMemberByIdResponse struct {
	Members map[int64]*Member
	Code    common.RspCode
	Msg     string
}

type QueryMemberRequest struct {
	Member   *Member
	Page     int64
	PageSize int64
}

type QueryMemberResponse struct {
	Members    map[int64]*Member
	TotalCount int64
	Code       common.RspCode
	Msg        string
}
