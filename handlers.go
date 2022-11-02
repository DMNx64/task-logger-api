package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var tasks = []task {}

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func createTask(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}
	newTask.ID = uuid.New()
	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, tasks)
}
