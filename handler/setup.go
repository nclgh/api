package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nclgh/lakawei_api/utils"
)

func SetUpRouter(router *gin.Engine) {
	SetRouter(router, setNoLoginRouter)
	SetRouter(router, setLoginRouter, utils.ApiLoginRequireMiddleWare)
}

func SetRouter(router *gin.Engine, setRouter func(*gin.RouterGroup), middleWares ...gin.HandlerFunc) {
	apiGroup := router.Group("/api", middleWares...)
	setRouter(apiGroup)
}

func setNoLoginRouter(g *gin.RouterGroup) {
	g.GET("/user/login/", LoginHandler)
}

func setLoginRouter(g *gin.RouterGroup) {
	g.GET("/ping/", func(c *gin.Context) {
		utils.ReplyOnce(c,200,"pong")
	})
}
