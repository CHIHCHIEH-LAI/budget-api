package controllers

import (
	"budget_manager/models"
	"budget_manager/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListCategories retrieves all categories
func ListCategories(c *gin.Context) {
	categories, err := services.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// CreateCategory creates a new category
func CreateCategory(c *gin.Context) {
	var categoryCreate models.CategoryCreate
	if err := c.ShouldBindJSON(&categoryCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := services.CreateCategory(categoryCreate)
	if err != nil {
		if err.Error() == "Category already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "Category already exists"})
			return
		} else if err.Error() == "Budget cannot be less than expense" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Budget cannot be less than expense"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
			return
		}
	}
	c.JSON(http.StatusCreated, category)
}

// UpdateCategoryBudget updates an existing category
func UpdateCategoryBudget(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var categoryUpdate models.CategoryBudgetUpdate
	if err := c.ShouldBindJSON(&categoryUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category, err := services.UpdateCategoryBudget(uint(id), categoryUpdate)
	if err != nil {
		if err.Error() == "Category not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		} else if err.Error() == "Budget cannot be less than expense" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Budget cannot be less than expense"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
			return
		}
	}
	c.JSON(http.StatusOK, category)
}

// UpdateCategoryExpense updates an existing category
func UpdateCategoryExpense(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var categoryUpdate models.CategoryExpenseUpdate
	if err := c.ShouldBindJSON(&categoryUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category, err := services.UpdateCategoryExpense(uint(id), categoryUpdate)
	if err != nil {
		if err.Error() == "Category not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		} else if err.Error() == "Budget cannot be less than expense" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Budget cannot be less than expense"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
			return
		}
	}
	c.JSON(http.StatusOK, category)
}

// DeleteCategory removes a category
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := services.DeleteCategory(uint(id)); err != nil {
		if err.Error() == "Category not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
			return
		}
	}
	c.Status(http.StatusNoContent)
}
