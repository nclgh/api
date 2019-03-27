package utils

import (
	"sync"
	"errors"
	"github.com/gin-gonic/gin"
)

func ReplyOnce(ctx *gin.Context, statusCode int, jsonBodyObj interface{}) {
	run, exist := ctx.Get(ContextKeyReplyOnce)
	if exist == true {
		runOnce := run.(*sync.Once)
		runOnce.Do(
			func() {
				ctx.AbortWithStatusJSON(statusCode, jsonBodyObj)
			})
	} else {
		ctx.AbortWithStatusJSON(statusCode, jsonBodyObj)
	}
}

func GetAuth(ctx *gin.Context) (*Auth, error) {
	vauth, exist := ctx.Get(ContextKeyAuth)
	if exist == false {
		return nil, errors.New("auth context not exist")
	}
	auth := vauth.(*Auth)
	return auth, nil
}
