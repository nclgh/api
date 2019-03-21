package main

import (
	"github.com/nclgh/lakawei_gin"
	"github.com/nclgh/lakawei_api/utils"
	"github.com/nclgh/lakawei_api/handler"
)

func main() {
	gin := lakawei_gin.Init()
	gin.Engine.Use(utils.PrepareMiddleWare)
	handler.SetUpRouter(gin.Engine)
	lakawei_gin.Run()
}
