package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	
	router.GET("/tasks", getTasks)
	router.POST("/tasks", createTask)

	router.Run("0.0.0.0:8080")
}
