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

type DeleteSessionRequest struct {
	UserId int64
}

type DeleteSessionResponse struct {
	IsSuccess bool
	Code      common.RspCode
	Msg       string
}
