package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(router *gin.Engine) {
	apiGroup := router.Group("/api")
	apiGroup.GET("/ping/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	apiGroup.POST("/user/login/", LoginHandler)
}
