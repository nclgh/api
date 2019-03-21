package utils

import (
	"sync"
	"github.com/gin-gonic/gin"
)

const (
	ContextKeyReplyOnce = "reply_once"
)

func PrepareMiddleWare(c *gin.Context) {
	runOnce := &sync.Once{}
	c.Set(ContextKeyReplyOnce, runOnce)
}
