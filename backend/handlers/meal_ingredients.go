package handlers

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rjmcnamara10/forage/db/repository"
	"github.com/rjmcnamara10/forage/db/sqlc"
)

// Request/Response DTOs
type CreateMealIngredientRequest struct {
	MealID                int32   `json:"meal_id" binding:"required"`
	IngredientID          int32   `json:"ingredient_id" binding:"required"`
	UnitID                *int32  `json:"unit_id"`
	RequiredAmount        *string `json:"required_amount"`
	IsSeasoning           bool    `json:"is_seasoning"`
	DepletesSlowly        bool    `json:"depletes_slowly"`
	AlternateIngredientID *int32  `json:"alternate_ingredient_id"`
}

type UpdateMealIngredientRequest struct {
	IngredientID          int32   `json:"ingredient_id" binding:"required"`
	UnitID                *int32  `json:"unit_id"`
	RequiredAmount        *string `json:"required_amount"`
	IsSeasoning           bool    `json:"is_seasoning"`
	DepletesSlowly        bool    `json:"depletes_slowly"`
	AlternateIngredientID *int32  `json:"alternate_ingredient_id"`
}

type MealIngredientResponse struct {
	ID                    int32   `json:"id"`
	MealID                int32   `json:"meal_id"`
	IngredientID          int32   `json:"ingredient_id"`
	UnitID                *int32  `json:"unit_id"`
	RequiredAmount        *string `json:"required_amount"`
	IsSeasoning           bool    `json:"is_seasoning"`
	DepletesSlowly        bool    `json:"depletes_slowly"`
	AlternateIngredientID *int32  `json:"alternate_ingredient_id"`
}

// GET /meal-ingredients/:id
func getMealIngredient(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid meal ingredient ID"})
			return
		}

		ingredient, err := repos.MealIngredients.GetMealIngredient(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(404, gin.H{"error": "Meal ingredient not found"})
			return
		}

		resp := mealIngredientToResponse(ingredient)
		c.JSON(200, resp)
	}
}

// GET /meal-ingredients?meal_id=123
func getMealIngredients(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		mealIDStr := c.Query("meal_id")
		if mealIDStr == "" {
			c.JSON(400, gin.H{"error": "meal_id query parameter is required"})
			return
		}

		mealID, err := strconv.ParseInt(mealIDStr, 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid meal ID"})
			return
		}

		// Verify meal exists
		_, err = repos.Meals.GetMeal(c.Request.Context(), int32(mealID))
		if err != nil {
			c.JSON(404, gin.H{"error": "Meal not found"})
			return
		}

		ingredients, err := repos.MealIngredients.ListMealIngredients(c.Request.Context(), int32(mealID))
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch meal ingredients"})
			return
		}

		response := make([]MealIngredientResponse, len(ingredients))
		for i, ingredient := range ingredients {
			response[i] = mealIngredientToResponse(ingredient)
		}

		c.JSON(200, response)
	}
}

// POST /meal-ingredients
func createMealIngredient(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateMealIngredientRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		// Validate that meal exists
		_, err := repos.Meals.GetMeal(c.Request.Context(), req.MealID)
		if err != nil {
			c.JSON(400, gin.H{"error": "Meal does not exist"})
			return
		}

		// Validate that ingredient item exists
		_, err = repos.Items.GetItem(c.Request.Context(), req.IngredientID)
		if err != nil {
			c.JSON(400, gin.H{"error": "Ingredient item does not exist"})
			return
		}

		// Convert optional fields to sql.Null types
		unitID := sql.NullInt32{Valid: false}
		if req.UnitID != nil {
			unitID = sql.NullInt32{Int32: *req.UnitID, Valid: true}
		}

		requiredAmount := sql.NullString{Valid: false}
		if req.RequiredAmount != nil {
			requiredAmount = sql.NullString{String: *req.RequiredAmount, Valid: true}
		}

		altIngredientID := sql.NullInt32{Valid: false}
		if req.AlternateIngredientID != nil {
			altIngredientID = sql.NullInt32{Int32: *req.AlternateIngredientID, Valid: true}
		}

		ingredient, err := repos.MealIngredients.CreateMealIngredient(
			c.Request.Context(),
			req.MealID,
			req.IngredientID,
			unitID,
			requiredAmount,
			req.IsSeasoning,
			req.DepletesSlowly,
			altIngredientID,
		)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create meal ingredient"})
			return
		}

		resp := mealIngredientToResponse(ingredient)
		c.JSON(201, resp)
	}
}

// PUT /meal-ingredients/:id
func updateMealIngredient(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid meal ingredient ID"})
			return
		}

		var req UpdateMealIngredientRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		// Validate that ingredient item exists
		_, err = repos.Items.GetItem(c.Request.Context(), req.IngredientID)
		if err != nil {
			c.JSON(400, gin.H{"error": "Ingredient item does not exist"})
			return
		}

		// Convert optional fields to sql.Null types
		unitID := sql.NullInt32{Valid: false}
		if req.UnitID != nil {
			unitID = sql.NullInt32{Int32: *req.UnitID, Valid: true}
		}

		requiredAmount := sql.NullString{Valid: false}
		if req.RequiredAmount != nil {
			requiredAmount = sql.NullString{String: *req.RequiredAmount, Valid: true}
		}

		altIngredientID := sql.NullInt32{Valid: false}
		if req.AlternateIngredientID != nil {
			altIngredientID = sql.NullInt32{Int32: *req.AlternateIngredientID, Valid: true}
		}

		ingredient, err := repos.MealIngredients.UpdateMealIngredient(
			c.Request.Context(),
			int32(id),
			req.IngredientID,
			unitID,
			requiredAmount,
			req.IsSeasoning,
			req.DepletesSlowly,
			altIngredientID,
		)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to update meal ingredient"})
			return
		}

		resp := mealIngredientToResponse(ingredient)
		c.JSON(200, resp)
	}
}

// DELETE /meal-ingredients/:id
func deleteMealIngredient(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid meal ingredient ID"})
			return
		}

		err = repos.MealIngredients.DeleteMealIngredient(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete meal ingredient"})
			return
		}

		c.JSON(204, nil)
	}
}

// Helper function to convert SQLC MealIngredient to response DTO
func mealIngredientToResponse(ingredient sqlc.MealIngredient) MealIngredientResponse {
	resp := MealIngredientResponse{
		ID:                    ingredient.ID,
		MealID:                ingredient.MealID,
		IngredientID:          ingredient.IngredientID,
		UnitID:                nil,
		RequiredAmount:        nil,
		IsSeasoning:           ingredient.IsSeasoning,
		DepletesSlowly:        ingredient.DepletesSlowly,
		AlternateIngredientID: nil,
	}

	if ingredient.UnitID.Valid {
		resp.UnitID = &ingredient.UnitID.Int32
	}
	if ingredient.RequiredAmount.Valid {
		resp.RequiredAmount = &ingredient.RequiredAmount.String
	}
	if ingredient.AlternateIngredientID.Valid {
		resp.AlternateIngredientID = &ingredient.AlternateIngredientID.Int32
	}

	return resp
}
