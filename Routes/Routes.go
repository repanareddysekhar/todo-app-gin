package Routes

import (
	"github.com/gin-gonic/gin"
	"todo-app/Controller"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/todos", Controller.GetAllTodo)
		v1.GET("/todo/:id", Controller.GetTodoByID)
		v1.POST("/todo", Controller.AddTodo)
		v1.PUT("/todo", Controller.UpdateTodo)
		v1.DELETE("/todo/:id", Controller.DeleteTodoByID)
	}

	return router
}
