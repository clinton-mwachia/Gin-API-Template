package routes

import (
	"gin-api/handlers"
	"gin-api/helpers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/v1/api")
	{
		// users routes
		api.POST("/login", handlers.Login)
		api.POST("/register", handlers.Register)
		api.GET("/user/:id", handlers.GetUserByID)
		api.PUT("/user/:id", handlers.UpdateUser)
		api.DELETE("/user/:id", handlers.DeleteUser)

		// protected routes
		secured := api.Group("")
		secured.Use(helpers.AuthMiddleware())
		secured.GET("/users", handlers.GetUsers)

		// tasks routes
		api.POST("/task", handlers.CreateTask)
		api.GET("/tasks", handlers.GetTasks)
		api.GET("/task/:id", handlers.GetTaskByID)
		api.PUT("/task/:id", handlers.UpdateTask)
		api.DELETE("/task/:id", handlers.DeleteTask)
		api.GET("/tasks/paginated", handlers.GetPaginatedTasks) //http://localhost:8080/v1/api/tasks/paginated?page=1&limit=5
	}
}
