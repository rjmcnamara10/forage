package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rjmcnamara10/forage/db/repository"
)

// Request/Response DTOs
type CreateItemCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateItemCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateMealCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateMealCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type CategoryResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

// GET /item-categories
func getItemCategories(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		categories, err := repos.ItemCategories.ListItemCategories(c.Request.Context())
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch item categories"})
			return
		}

		response := make([]CategoryResponse, len(categories))
		for i, cat := range categories {
			response[i] = CategoryResponse{
				ID:   cat.ID,
				Name: cat.Name,
			}
		}

		c.JSON(200, response)
	}
}

// GET /item-categories/:id
func getItemCategory(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid category ID"})
			return
		}

		category, err := repos.ItemCategories.GetItemCategory(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(404, gin.H{"error": "Item category not found"})
			return
		}

		c.JSON(200, CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}
}

// POST /item-categories
func createItemCategory(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateItemCategoryRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		category, err := repos.ItemCategories.CreateItemCategory(c.Request.Context(), req.Name)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create item category"})
			return
		}

		c.JSON(201, CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}
}

// PUT /item-categories/:id
func updateItemCategory(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid category ID"})
			return
		}

		var req UpdateItemCategoryRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		category, err := repos.ItemCategories.UpdateItemCategory(c.Request.Context(), int32(id), req.Name)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to update item category"})
			return
		}

		c.JSON(200, CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}
}

// DELETE /item-categories/:id
func deleteItemCategory(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid category ID"})
			return
		}

		err = repos.ItemCategories.DeleteItemCategory(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete item category"})
			return
		}

		c.JSON(204, nil)
	}
}

// GET /meal-categories
func getMealCategories(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		categories, err := repos.MealCategories.ListMealCategories(c.Request.Context())
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch meal categories"})
			return
		}

		response := make([]CategoryResponse, len(categories))
		for i, cat := range categories {
			response[i] = CategoryResponse{
				ID:   cat.ID,
				Name: cat.Name,
			}
		}

		c.JSON(200, response)
	}
}

// GET /meal-categories/:id
func getMealCategory(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid category ID"})
			return
		}

		category, err := repos.MealCategories.GetMealCategory(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(404, gin.H{"error": "Meal category not found"})
			return
		}

		c.JSON(200, CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}
}

// POST /meal-categories
func createMealCategory(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateMealCategoryRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		category, err := repos.MealCategories.CreateMealCategory(c.Request.Context(), req.Name)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create meal category"})
			return
		}

		c.JSON(201, CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}
}

// PUT /meal-categories/:id
func updateMealCategory(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid category ID"})
			return
		}

		var req UpdateMealCategoryRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		category, err := repos.MealCategories.UpdateMealCategory(c.Request.Context(), int32(id), req.Name)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to update meal category"})
			return
		}

		c.JSON(200, CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}
}

// DELETE /meal-categories/:id
func deleteMealCategory(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid category ID"})
			return
		}

		err = repos.MealCategories.DeleteMealCategory(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete meal category"})
			return
		}

		c.JSON(204, nil)
	}
}
