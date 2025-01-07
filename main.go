package main

import (
	"budget_manager/database"
	"budget_manager/routes"
	"budget_manager/utility"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load environment variables
	dotenv_filepath := ".env"
	utility.LoadDotenvIfExists(dotenv_filepath)

	// Get database connection information from environment variables
	dbConfig := utility.GetDBConfig()

	// Open the database
	err := database.Open(dbConfig)
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
