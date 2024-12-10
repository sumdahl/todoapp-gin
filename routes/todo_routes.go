package routes

import (
	"gin-todo/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupTodoRoutes(router *gin.Engine, db *gorm.DB) {
	// Pass the db to each handler
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Define routes
	router.POST("/todos", controllers.CreateTodoController)
	router.GET("/todos", controllers.GetTodosController)
	router.GET("/todos/:id", controllers.GetTodoByIDController)
	router.PUT("/todos/:id", controllers.UpdateTodoController)
	router.DELETE("/todos/:id", controllers.DeleteTodoController)
}
