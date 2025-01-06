package main

import (
	"budget_manager/database"
	"budget_manager/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	// Open the database
	err := database.Open()
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer database.Close()

	// Set up the Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterCategoryRoutes(router)

	// Start the server
	log.Println("Starting server on port 8000")
	log.Fatal(router.Run(":8000"))
}
