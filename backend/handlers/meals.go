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
		sortBy := "name"
		sortOrder := "ASC"
		query := c.Query("q")
		mealCategoryID := c.Query("meal_category_id")

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

		if sb := c.Query("sort_by"); sb != "" && (sb == "name" || sb == "servings") {
			sortBy = sb
		}

		if so := c.Query("sort_order"); so != "" && (so == "ASC" || so == "DESC") {
			sortOrder = so
		}

		var meals []MealResponse
		var total int64

		if query != "" {
			// Search mode
			searchMeals, err := repos.Meals.SearchMeals(c.Request.Context(), query, limit, offset, sortOrder)
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to fetch meals"})
				return
			}
			meals = make([]MealResponse, len(searchMeals))
			for i, meal := range searchMeals {
				meals[i] = MealResponse{
					ID:             meal.ID,
					Name:           meal.Name,
					MealCategoryID: meal.MealCategoryID,
					Servings:       meal.Servings,
				}
			}
			total, _ = repos.Meals.SearchMealsCount(c.Request.Context(), query)
		} else if mealCategoryID != "" {
			// Category filter mode
			catID, err := strconv.ParseInt(mealCategoryID, 10, 32)
			if err != nil {
				c.JSON(400, gin.H{"error": "Invalid meal category ID"})
				return
			}
			filterMeals, err := repos.Meals.ListMealsByCategory(c.Request.Context(), int32(catID), limit, offset, sortBy, sortOrder)
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to fetch meals"})
				return
			}
			meals = make([]MealResponse, len(filterMeals))
			for i, meal := range filterMeals {
				meals[i] = MealResponse{
					ID:             meal.ID,
					Name:           meal.Name,
					MealCategoryID: meal.MealCategoryID,
					Servings:       meal.Servings,
				}
			}
			total, _ = repos.Meals.ListMealsByCategoryCount(c.Request.Context(), int32(catID))
		} else {
			// List all meals
			listMeals, err := repos.Meals.ListMeals(c.Request.Context(), limit, offset, sortBy, sortOrder)
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to fetch meals"})
				return
			}
			meals = make([]MealResponse, len(listMeals))
			for i, meal := range listMeals {
				meals[i] = MealResponse{
					ID:             meal.ID,
					Name:           meal.Name,
					MealCategoryID: meal.MealCategoryID,
					Servings:       meal.Servings,
				}
			}
			total, _ = repos.Meals.ListMealsCount(c.Request.Context())
		}

		c.JSON(200, PaginatedResponse{
			Data:   meals,
			Total:  total,
			Limit:  limit,
			Offset: offset,
		})
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
