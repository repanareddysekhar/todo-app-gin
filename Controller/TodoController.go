package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func GetTodoByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo Model.Todo
	err := Repository.GetTodoByID(&todo, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateTodo(c *gin.Context) {
	var todo Model.Todo
	c.BindJSON(&todo)

	err := Repository.UpdateTodo(&todo)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodoByID(c *gin.Context) {
	idParam := c.Params.ByName("id")

	id, err := strconv.Atoi(idParam)
	err = Repository.DeleteTodoByID(uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "deleted"})
	}
}
