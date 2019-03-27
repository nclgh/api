package utils

import (
	"sync"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/nclgh/lakawei_api/rpc"
	"github.com/nclgh/lakawei_rpc/client"
	"github.com/nclgh/lakawei_scaffold/rpc/passport"
)

const (
	ContextKeyReplyOnce = "reply_once"

	ContextKeyAuth = "auth"
)

type Auth struct {
	Session string
	UserId  int64
}

func PrepareMiddleWare(c *gin.Context) {
	runOnce := &sync.Once{}
	c.Set(ContextKeyReplyOnce, runOnce)
}

func ApiLoginRequireMiddleWare(c *gin.Context) {
	sid, err := c.Cookie(SessionKey)
	if err != nil {
		logrus.Infof("ApiLoginRequireMiddleWare get empty sid")
		ReplyOnce(c, 200, NewCommonResponse(CodeLoginRequire))
		return
	}
	// TODO 开个后门测试
	if sid == "dec84587-53f0-4a6a-9ffc-7d712b2a5140" {
		auth := &Auth{
			Session: sid,
			UserId:  1,
		}
		c.Set(ContextKeyAuth, auth)
		return
	}
	var rspGetSession passport.GetSessionResponse
	err = rpc.GetPassportClient().Call(&client.RpcRequestCtx{}, "GetSession", passport.GetSessionRequest{SessionId: sid}, &rspGetSession)
	if err != nil || rspGetSession.Code != 0 {
		logrus.Errorf("ApiLoginRequireMiddleWare get session err: %v, rsp code: %v, userId: %v", err, rspGetSession.Code, rspGetSession.UserId)
		ReplyOnce(c, 200, NewCommonResponse(CodeFailed))
		return
	}
	if rspGetSession.UserId <= 0 {
		logrus.Infof("ApiLoginRequireMiddleWare get invalid sid: %v", sid)
		ReplyOnce(c, 200, NewCommonResponse(CodeLoginRequire))
		return
	}
	auth := &Auth{
		Session: sid,
		UserId:  rspGetSession.UserId,
	}
	c.Set(ContextKeyAuth, auth)
}
