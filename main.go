package main

import (
	"gin-be/routes"
	"gin-be/services"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Connect to the SQLite database
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auto-migrate the Todo model to create the table
	db.AutoMigrate(&services.Todo{})

	// Create a new Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupTodoRoutes(router, db)

	// Start the server
	router.Run(":8080")
}
