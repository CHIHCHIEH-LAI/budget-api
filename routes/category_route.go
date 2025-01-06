package routes

import (
	"budget_manager/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterCategoryRoutes registers category-related routes
func RegisterCategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/api/v1/categories")
	{
		categoryRoutes.GET("", controllers.ListCategories)
		categoryRoutes.POST("", controllers.CreateCategory)
		categoryRoutes.PUT("/:id/budget", controllers.UpdateCategoryBudget)
		categoryRoutes.PUT("/:id/expense", controllers.UpdateCategoryExpense)
		categoryRoutes.DELETE("/:id", controllers.DeleteCategory)
	}
}
