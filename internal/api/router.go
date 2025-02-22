package api

import (
    "github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Welcome to the API",
			})
		})
		RegisterV1Routes(api)
	}
}
func RegisterV1Routes(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	{
		RegisterUserRoutes(v1)
        
        // Other feature-specific route registration
        // e.g., RegisterStreamRoutes(v1)
    }
}
