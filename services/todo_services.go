package services

import (
	"gin-todo/models"

	"gorm.io/gorm"
)

// Create a new Todo in the database
func CreateTodo(db *gorm.DB, todo *models.Todo) error {
	result := db.Create(todo)
	return result.Error
}

// Get all Todos from the database
func GetTodos(db *gorm.DB) ([]models.Todo, error) {
	var todos []models.Todo
	result := db.Find(&todos)
	return todos, result.Error
}

// Get a specific Todo by ID from the database
func GetTodoByID(db *gorm.DB, id uint) (*models.Todo, error) {
	var todo models.Todo
	result := db.First(&todo, id)
	return &todo, result.Error
}

// Update a Todo in the database
func UpdateTodoByID(db *gorm.DB, id uint, todo *models.Todo) error {
	var existingTodo models.Todo
	result := db.First(&existingTodo, id)
	if result.Error != nil {
		return result.Error
	}
	// Update the existing Todo with new data
	existingTodo.Title = todo.Title
	existingTodo.Description = todo.Description
	db.Save(&existingTodo)
	return nil
}

// Delete a Todo by ID from the database
func DeleteTodoByID(db *gorm.DB, id uint) error {
	var todo models.Todo
	result := db.First(&todo, id)
	if result.Error != nil {
		return result.Error
	}
	db.Delete(&todo)
	return nil
}
