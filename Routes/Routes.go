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
		v1.POST("/todo", Controller.AddTodo)
	}

	return router
}
