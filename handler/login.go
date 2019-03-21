package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nclgh/lakawei_rpc/client"
	"github.com/nclgh/lakawei_scaffold/rpc/passport"
	"github.com/sirupsen/logrus"
	"github.com/nclgh/lakawei_api/utils"
)

var (
	passportCli *client.RpcClient
)

func init() {
	cli, err := client.Init("ServicePassport")
	if err != nil {
		panic(err)
	}
	passportCli = cli
}

func LoginHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "LoginHandler")
	req := passport.CreateSessionRequest{
		UserId: 666,
	}
	rsp := passport.CreateSessionResponse{}
	err := passportCli.Call(&client.RpcRequestCtx{}, "CreateSession", req, &rsp)
	if err != nil {
		logrus.Errorf("call ServicePassport.CreateSession err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, "")
		return
	}
	ctx.SetCookie(utils.SessionKey, rsp.SessionId, utils.SessionLife, "/", utils.ServerDomain, false, true)
	p.Success(nil, "")
}
