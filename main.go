package main

import (
	"encoding/gob"
	"github.com/nclgh/lakawei_gin"
	"github.com/nclgh/lakawei_api/rpc"
	"github.com/nclgh/lakawei_api/utils"
	"github.com/nclgh/lakawei_api/handler"
	"github.com/nclgh/lakawei_scaffold/rpc/common"
)

func init()  {
	gob.Register(common.TimeFilter{})
}

func initCommon() {
	rpc.Init()
}

func main() {
	//test()
	initCommon()
	gin := lakawei_gin.Init()
	gin.Engine.Use(utils.RequestReport)
	gin.Engine.Use(utils.PrepareMiddleWare)
	gin.Engine.Use(utils.AllowCrossOrigin)
	handler.SetUpRouter(gin.Engine)
	lakawei_gin.Run()
}
