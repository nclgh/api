package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nclgh/lakawei_api/utils"
	"github.com/nclgh/lakawei_api/handler"
)

func main() {
	router := gin.Default()
	router.Use(utils.PrepareMiddleWare)
	handler.SetUpRouter(router)
	router.Run(":8080")
}
