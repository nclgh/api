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
	g.POST("/user/login/", LoginHandler)
}

func setLoginRouter(g *gin.RouterGroup) {
	g.GET("/ping/", func(c *gin.Context) {
		utils.ReplyOnce(c, 200, "pong")
	})
	g.POST("/user/logout/", LogoutHandler)
	g.POST("/department/add/", AddDepartmentHandler)
	g.POST("/department/delete/", DeleteDepartmentHandler)
	g.GET("/department/query/", QueryDepartmentHandler)
	g.POST("/department/member/add/", AddMemberHandler)
	g.POST("/department/member/delete/", DeleteMemberHandler)
	g.GET("/department/member/query/", QueryMemberHandler)
	g.POST("/device/manufacturer/add/", AddManufacturerHandler)
	g.GET("/device/manufacturer/query/", QueryManufacturerHandler)
	g.POST("/device/document/add/", AddDeviceHandler)
	g.POST("/device/document/delete/", DeleteDeviceHandler)
	g.GET("/device/document/query/", QueryDeviceHandler)
	g.POST("/device/achievement/add/", AddAchievementHandler)
	g.POST("/device/achievement/delete/", DeleteAchievementHandler)
	g.GET("/device/achievement/query/", QueryAchievementHandler)
	g.POST("/device/rent/lend/", AddRentHandler)
	g.POST("/device/rent/return/", ReturnRentHandler)
	g.GET("/device/rent/query/", QueryRentHandler)
}
