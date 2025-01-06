package services

import (
	"budget_manager/database"
	"budget_manager/models"
	"database/sql"
	"fmt"
)

// GetAllCategories retrieves all categories
func GetAllCategories() ([]models.Category, error) {
	rows, err := database.DB.Query("SELECT id, name, budget, expense FROM categories")
	if err != nil {
		return nil, fmt.Errorf("failed to query categories: %v", err)
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Budget, &category.Expense); err != nil {
			return nil, fmt.Errorf("failed to scan category: %v", err)
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate categories: %v", err)
	}

	return categories, nil
}

// CreateCategory adds a new category
func CreateCategory(category models.CategoryCreate) (models.Category, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		return models.Category{}, fmt.Errorf("failed to start transaction: %v", err)
	}

	if category.Budget < category.Expense {
		tx.Rollback()
		return models.Category{}, fmt.Errorf("budget cannot be less than expense")
	}

	var id uint
	err = tx.QueryRow("SELECT id FROM categories WHERE name = $1 FOR UPDATE", category.Name).Scan(&id)
	if err == nil {
		tx.Rollback()
		return models.Category{}, fmt.Errorf("category already exists")
	} else if err != sql.ErrNoRows {
		tx.Rollback()
		return models.Category{}, fmt.Errorf("failed to query category: %v", err)
	}

	_, err = tx.Exec("INSERT INTO categories (name, budget, expense) VALUES ($1, $2, $3)", category.Name, category.Budget, category.Expense)
	if err != nil {
		tx.Rollback()
		return models.Category{}, fmt.Errorf("failed to insert category: %v", err)
	}

	var categoryCreated models.Category
	err = tx.QueryRow("SELECT id, name, budget, expense FROM categories WHERE name = $1", category.Name).Scan(&categoryCreated.ID, &categoryCreated.Name, &categoryCreated.Budget, &categoryCreated.Expense)
	if err != nil {
		tx.Rollback()
		return models.Category{}, fmt.Errorf("failed to query category: %v", err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return models.Category{}, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return categoryCreated, nil
}

// UpdateCategoryBudget updates an existing category
func UpdateCategoryBudget(id uint, categoryUpdate models.CategoryBudgetUpdate) (models.Category, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		return models.Category{}, fmt.Errorf("failed to start transaction: %v", err)
	}

	var category models.Category
	err = tx.QueryRow("SELECT id, name, budget, expense FROM categories WHERE id = $1 FOR UPDATE", id).Scan(&category.ID, &category.Name, &category.Budget, &category.Expense)
	if err == sql.ErrNoRows {
		tx.Rollback()
		return models.Category{}, fmt.Errorf("category does not exist")
	} else if err != nil {
		tx.Rollback()
		return models.Category{}, fmt.Errorf("failed to query category: %v", err)
	}

	if categoryUpdate.Budget < category.Expense {
		tx.Rollback()
		return models.Category{}, fmt.Errorf("budget cannot be less than expense")
	}

	_, err = tx.Exec("UPDATE categories SET budget = $1 WHERE id = $2", categoryUpdate.Budget, id)
	if err != nil {
		tx.Rollback()
		return models.Category{}, fmt.Errorf("failed to update category: %v", err)
	}

	err = tx.QueryRow("SELECT id, name, budget, expense FROM categories WHERE id = $1", id).Scan(&category.ID, &category.Name, &category.Budget, &category.Expense)
	if err != nil {
		tx.Rollback()
		return models.Category{}, fmt.Errorf("failed to query category: %v", err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return models.Category{}, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return category, nil
}

// UpdateCategoryExpense updates an existing category
func UpdateCategoryExpense(id uint, categoryUpdate models.CategoryExpenseUpdate) (models.Category, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		return models.Category{}, fmt.Errorf("failed to start transaction: %v", err)
	}

	var category models.Category
	err = tx.QueryRow("SELECT id, name, budget, expense FROM categories WHERE id = $1 FOR UPDATE", id).Scan(&category.ID, &category.Name, &category.Budget, &category.Expense)
	if err == sql.ErrNoRows {
		tx.Rollback()
		return models.Category{}, fmt.Errorf("category does not exist")
	} else if err != nil {
		tx.Rollback()
		return models.Category{}, fmt.Errorf("failed to query category: %v", err)
	}

	if category.Budget < categoryUpdate.Expense {
		tx.Rollback()
		return models.Category{}, fmt.Errorf("expense cannot be more than budget")
	}

	_, err = tx.Exec("UPDATE categories SET expense = $1 WHERE id = $2", categoryUpdate.Expense, id)
	if err != nil {
		tx.Rollback()
		return models.Category{}, fmt.Errorf("failed to update category: %v", err)
	}

	err = tx.QueryRow("SELECT id, name, budget, expense FROM categories WHERE id = $1", id).Scan(&category.ID, &category.Name, &category.Budget, &category.Expense)
	if err != nil {
		tx.Rollback()
		return models.Category{}, fmt.Errorf("failed to query category: %v", err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return models.Category{}, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return category, nil
}

// DeleteCategory removes a category by ID
func DeleteCategory(id uint) error {
	tx, err := database.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}

	err = tx.QueryRow("SELECT id FROM categories WHERE id = $1 FOR UPDATE", id).Scan(&id)
	if err == sql.ErrNoRows {
		tx.Rollback()
		return fmt.Errorf("category does not exist")
	} else if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to query category: %v", err)
	}

	_, err = tx.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete category: %v", err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}
