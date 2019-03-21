package passport

import (
	"github.com/nclgh/lakawei_scaffold/rpc/common"
)

type CreateSessionRequest struct {
	UserId int64
}

type CreateSessionResponse struct {
	SessionId string
	Code      common.RspCode
	Msg       string
}

type GetSessionRequest struct {
	SessionId string
}

type GetSessionResponse struct {
	UserId int64
	Code   common.RspCode
	Msg    string
}

func GetCreateSessionResponse(code common.RspCode, msg string) *CreateSessionResponse {
	rsp := &CreateSessionResponse{
		Code: code,
		Msg:  msg,
	}
	if msg == "" {
		dMsg := common.UnknownCodeMsg
		m, ok := common.CodeMsg[code]
		if ok {
			dMsg = m
		}
		rsp.Msg = dMsg
	}
	return rsp
}

func GetGetSessionResponse(code common.RspCode, msg string) *GetSessionResponse {
	rsp := &GetSessionResponse{
		Code: code,
		Msg:  msg,
	}
	if msg == "" {
		dMsg := common.UnknownCodeMsg
		m, ok := common.CodeMsg[code]
		if ok {
			dMsg = m
		}
		rsp.Msg = dMsg
	}
	return rsp
}
