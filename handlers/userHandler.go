package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// register a user
func Register(c *gin.Context) {
	fmt.Println("create a user")
}

// login a user
func Login(c *gin.Context) {
	fmt.Println("login a user")
}
