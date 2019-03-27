package rpc

import (
	"fmt"
	"github.com/nclgh/lakawei_rpc/client"
	"github.com/nclgh/lakawei_scaffold/rpc/passport"
	"github.com/nclgh/lakawei_scaffold/rpc/common"
)

func CreateSession(ctx *client.RpcRequestCtx, userId int64) (*passport.CreateSessionResponse, error) {
	req := passport.CreateSessionRequest{
		UserId: userId,
	}
	rsp := &passport.CreateSessionResponse{}
	err := GetPassportClient().Call(&client.RpcRequestCtx{}, "CreateSession", req, rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Code != common.CodeSuccess {
		return rsp, fmt.Errorf("call CreateSession failed. code: %v, msg: %v", rsp.Code, rsp.Msg)
	}
	return rsp, nil
}
