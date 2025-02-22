package ui

import (
	"github.com/gin-gonic/gin"
)

func RegisterUIRoutes(router *gin.Engine) {
	router.GET("/", IndexHandler)
	gui := router.Group("/ui")
	{
		gui.GET("/", IndexHandler)
	}
}
