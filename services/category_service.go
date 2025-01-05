package services

import (
	"budget_manager/database"
	"budget_manager/models"
)

// GetAllCategories retrieves all categories
func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := database.DB.Find(&categories).Error
	return categories, err
}

// CreateCategory adds a new category
func CreateCategory(category models.Category) error {
	return database.DB.Create(&category).Error
}

// UpdateCategory updates an existing category
func UpdateCategory(id uint, updatedData models.Category) (models.Category, error) {
	var category models.Category
	err := database.DB.First(&category, id).Error
	if err != nil {
		return category, err
	}
	category.Budget = updatedData.Budget
	category.Used = updatedData.Used
	err = database.DB.Save(&category).Error
	return category, err
}

// DeleteCategory removes a category by ID
func DeleteCategory(id uint) error {
	return database.DB.Delete(&models.Category{}, id).Error
}
