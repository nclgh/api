package handler

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/nclgh/lakawei_api/utils"
	"net/http"
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
		rsp.Message = msg
	}
	utils.ReplyOnce(p.Ctx, http.StatusOK, rsp)
}

func (p *Processor) Success(data map[string]interface{}, msg string) {
	rsp := utils.NewCommonResponse(0)
	if msg != "" {
		rsp.Message = msg
	}
	rsp.Data = data
	utils.ReplyOnce(p.Ctx, http.StatusOK, rsp)
}
