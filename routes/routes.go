package routes

import (
	"gin-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/v1/api")
	{
		// users routes
		api.POST("/login", handlers.Login)
		api.POST("/register", handlers.Register)
		api.GET("/user/:id", handlers.GetUserByID)
		api.GET("/users", handlers.GetUsers)
		api.PUT("/user/:id", handlers.UpdateUser)
		api.DELETE("/user/:id", handlers.DeleteUser)

		// tasks routes
		api.POST("/task", handlers.CreateTask)
		api.GET("/tasks", handlers.GetTasks)
		api.GET("/task/:id", handlers.GetTaskByID)
		api.PUT("/task/:id", handlers.UpdateTask)
		api.DELETE("/task/:id", handlers.DeleteTask)
		api.GET("/tasks/paginated", handlers.GetPaginatedTasks) //http://localhost:8080/v1/api/tasks/paginated?page=1&limit=5
	}
}
