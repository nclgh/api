package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/nclgh/lakawei_api/rpc"
	"github.com/nclgh/lakawei_api/conf"
	"github.com/nclgh/lakawei_api/utils"
	"github.com/nclgh/lakawei_rpc/client"
)

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func LoginHandler(ctx *gin.Context) {
	p := NewProcessor(ctx, "LoginHandler")

	form := LoginForm{}
	if ok := p.BindAndCheckForm(&form); !ok {
		return
	}

	cRsp, err := rpc.CheckUserIdentity(&client.RpcRequestCtx{}, form.Username, form.Password)
	if err != nil {
		logrus.Errorf("call ServiceMember.CheckUserIdentity err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, "")
		return
	}

	if cRsp.UserId <= 0 {
		logrus.Infof("identity check failed. username: %v, msg: %v", form.Username, cRsp.Msg)
		p.AbortWithMsg(utils.CodeLoginRequire, cRsp.Msg)
		return
	}

	rsp, err := rpc.CreateSession(&client.RpcRequestCtx{}, cRsp.UserId)
	if err != nil {
		logrus.Errorf("call ServicePassport.CreateSession err: %v", err)
		p.AbortWithMsg(utils.CodeFailed, "")
		return
	}
	ctx.SetCookie(utils.SessionKey, rsp.SessionId, utils.SessionLife, "/", conf.GetDomain(), false, true)
	// TODO
	ctx.SetCookie(utils.SessionKey, rsp.SessionId, utils.SessionLife, "/", "127.0.0.1", false, true)
	p.Success(nil, "")
}
