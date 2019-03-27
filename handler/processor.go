package handler

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/nclgh/lakawei_api/utils"
	"net/http"
	"github.com/sirupsen/logrus"
)

type Processor struct {
	Ctx         *gin.Context
	HandlerName string
	BeginTime   time.Time
}

func NewProcessor(ginCtx *gin.Context, handlerName string) (*Processor) {
	return &Processor{
		Ctx:         ginCtx,
		HandlerName: handlerName,
		BeginTime:   time.Now(),
	}
}

func (p *Processor) AbortWithMsg(code int, msg string) {
	rsp := utils.NewCommonResponse(code)
	if msg != "" {
		rsp.Msg = msg
	}
	utils.ReplyOnce(p.Ctx, http.StatusOK, rsp)
}

func (p *Processor) Success(data map[string]interface{}, msg string) {
	rsp := utils.NewCommonResponse(0)
	if msg != "" {
		rsp.Msg = msg
	}
	rsp.Data = data
	utils.ReplyOnce(p.Ctx, http.StatusOK, rsp)
}

func (p *Processor) BindAndCheckForm(form interface{}) bool {
	if err := p.Ctx.ShouldBind(form); err != nil {
		logrus.Infof("%v param error. error: %v", p.HandlerName, err)
		p.AbortWithMsg(utils.CodePARAMERR, "")
		return false
	}
	return true
}
