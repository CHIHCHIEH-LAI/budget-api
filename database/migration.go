package database

import (
	"log"

	"budget_manager/models"
)

// Migrate runs database migrations
func Migrate() {
	if err := DB.AutoMigrate(&models.Category{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
