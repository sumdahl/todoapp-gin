package controllers

import (
	"gin-todo/models"
	"gin-todo/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Create a new Todo
func CreateTodoController(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// Call service to create Todo
	if err := services.CreateTodo(c.MustGet("db").(*gorm.DB), &todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// Get all Todos
func GetTodosController(c *gin.Context) {
	todos, err := services.GetTodos(c.MustGet("db").(*gorm.DB))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// Get a Todo by ID
func GetTodoByIDController(c *gin.Context) {
	todoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	todo, err := services.GetTodoByID(c.MustGet("db").(*gorm.DB), uint(todoID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// Update a Todo by ID
func UpdateTodoController(c *gin.Context) {
	todoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedTodo models.Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// Call service to update Todo
	if err := services.UpdateTodoByID(c.MustGet("db").(*gorm.DB), uint(todoID), &updatedTodo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Todo"})
		return
	}

	c.JSON(http.StatusOK, updatedTodo)
}

// Delete a Todo by ID
func DeleteTodoController(c *gin.Context) {
	todoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Call service to delete Todo
	if err := services.DeleteTodoByID(c.MustGet("db").(*gorm.DB), uint(todoID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
