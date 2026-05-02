package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rjmcnamara10/forage/db/repository"
)

// Request/Response DTOs
type CreateMealRequest struct {
	Name           string `json:"name" binding:"required"`
	MealCategoryID int32  `json:"meal_category_id" binding:"required"`
	Servings       int32  `json:"servings" binding:"required"`
}

type UpdateMealRequest struct {
	Name           string `json:"name" binding:"required"`
	MealCategoryID int32  `json:"meal_category_id" binding:"required"`
	Servings       int32  `json:"servings" binding:"required"`
}

type MealResponse struct {
	ID             int32  `json:"id"`
	Name           string `json:"name"`
	MealCategoryID int32  `json:"meal_category_id"`
	Servings       int32  `json:"servings"`
}

// GET /meals
func getMeals(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit := int32(50)
		offset := int32(0)

		if l := c.Query("limit"); l != "" {
			if parsed, err := strconv.ParseInt(l, 10, 32); err == nil {
				limit = int32(parsed)
			}
		}

		if o := c.Query("offset"); o != "" {
			if parsed, err := strconv.ParseInt(o, 10, 32); err == nil {
				offset = int32(parsed)
			}
		}

		meals, err := repos.Meals.ListMeals(c.Request.Context(), limit, offset)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch meals"})
			return
		}

		response := make([]MealResponse, len(meals))
		for i, meal := range meals {
			response[i] = MealResponse{
				ID:             meal.ID,
				Name:           meal.Name,
				MealCategoryID: meal.MealCategoryID,
				Servings:       meal.Servings,
			}
		}

		c.JSON(200, response)
	}
}

// GET /meals/:id
func getMeal(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid meal ID"})
			return
		}

		meal, err := repos.Meals.GetMeal(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(404, gin.H{"error": "Meal not found"})
			return
		}

		c.JSON(200, MealResponse{
			ID:             meal.ID,
			Name:           meal.Name,
			MealCategoryID: meal.MealCategoryID,
			Servings:       meal.Servings,
		})
	}
}

// POST /meals
func createMeal(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateMealRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		meal, err := repos.Meals.CreateMeal(c.Request.Context(), req.Name, req.MealCategoryID, req.Servings)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create meal"})
			return
		}

		c.JSON(201, MealResponse{
			ID:             meal.ID,
			Name:           meal.Name,
			MealCategoryID: meal.MealCategoryID,
			Servings:       meal.Servings,
		})
	}
}

// PUT /meals/:id
func updateMeal(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid meal ID"})
			return
		}

		var req UpdateMealRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		meal, err := repos.Meals.UpdateMeal(c.Request.Context(), int32(id), req.Name, req.MealCategoryID, req.Servings)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to update meal"})
			return
		}

		c.JSON(200, MealResponse{
			ID:             meal.ID,
			Name:           meal.Name,
			MealCategoryID: meal.MealCategoryID,
			Servings:       meal.Servings,
		})
	}
}

// DELETE /meals/:id
func deleteMeal(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid meal ID"})
			return
		}

		err = repos.Meals.DeleteMeal(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete meal"})
			return
		}

		c.JSON(204, nil)
	}
}
