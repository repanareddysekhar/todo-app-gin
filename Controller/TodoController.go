package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-app/Model"
	"todo-app/Repository"
)

func GetAllTodo(c *gin.Context) {
	var todo []Model.Todo

	err := Repository.GetAllTodos(&todo)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func AddTodo(c *gin.Context) {
	var todo Model.Todo
	// Bind the JSON payload to the Todo model
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	err := Repository.CreateATodo(&todo)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
