package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/nclgh/lakawei_api/rpc"
	"github.com/nclgh/lakawei_api/conf"
	"github.com/nclgh/lakawei_api/utils"
	"github.com/nclgh/lakawei_rpc/client"
)

func LoginHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "LoginHandler")
	rsp, err := rpc.CreateSession(&client.RpcRequestCtx{}, 666)
	if err != nil {
		logrus.Errorf("call ServicePassport.CreateSession err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, "")
		return
	}
	ctx.SetCookie(utils.SessionKey, rsp.SessionId, utils.SessionLife, "/", conf.GetDomain(), false, true)
	p.Success(nil, "")
}
