package Repository

import (
	"todo-app/Config"
	"todo-app/Model"
)

func GetAllTodos(todo *[]Model.Todo) (err error) {
	return Config.DB.Find(todo).Error
}

func CreateATodo(todo *Model.Todo) (err error) {
	return Config.DB.Create(todo).Error
}
