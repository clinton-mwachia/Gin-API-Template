package routes

import (
	"gin-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/v1/api")
	{
		api.POST("/login", handlers.Login)
	}
}
