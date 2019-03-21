package rpc

import (
	"github.com/nclgh/lakawei_rpc/client"
	"github.com/nclgh/lakawei_scaffold/rpc/passport"
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
	return rsp, nil
}
