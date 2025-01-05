package main

import (
	"budget_manager/database"
	"budget_manager/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Specify the database file path
	dbFilePath := "budget.db"

	// Initialize the database
	database.Init(dbFilePath)

	// Run migrations
	database.Migrate()

	// Set up the Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterCategoryRoutes(router)

	// Start the server
	router.Run(":8000")
}
