package utils

import (
	"sync"
	"github.com/gin-gonic/gin"
)

func ReplyOnce(ctx *gin.Context, statusCode int, jsonBodyObj interface{}) {
	run, exist := ctx.Get(ContextKeyReplyOnce)
	if exist == true {
		runOnce := run.(*sync.Once)
		runOnce.Do(
			func() {
				ctx.JSON(statusCode, jsonBodyObj)
			})
	} else {
		ctx.JSON(statusCode, jsonBodyObj)
	}
}
