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

func GetTodoByID(todo *Model.Todo, id string) (err error) {
	return Config.DB.Find(&todo, id).Error
}

func UpdateTodo(todo *Model.Todo) (err error) {
	return Config.DB.Save(&todo).Error
}

func DeleteTodoByID(id uint) (err error) {
	return Config.DB.Delete(&Model.Todo{}, id).Error
}
