package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	fmt.Println("create a task")
}

func GetTasks(c *gin.Context) {
	fmt.Println("get task")
}
