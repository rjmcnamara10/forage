package repository

import (
	"context"
	"database/sql"

	"github.com/rjmcnamara10/forage/db/sqlc"
)

// MealRepository wraps sqlc queries for meals
type MealRepository struct {
	queries *sqlc.Queries
}

func NewMealRepository(q *sqlc.Queries) *MealRepository {
	return &MealRepository{queries: q}
}

func (r *MealRepository) GetMeal(ctx context.Context, id int32) (sqlc.Meal, error) {
	return r.queries.GetMeal(ctx, id)
}

func (r *MealRepository) GetMealByName(ctx context.Context, name string) (sqlc.Meal, error) {
	return r.queries.GetMealByName(ctx, name)
}

func (r *MealRepository) ListMeals(ctx context.Context, limit int32, offset int32) ([]sqlc.Meal, error) {
	return r.queries.ListMeals(ctx, sqlc.ListMealsParams{Limit: limit, Offset: offset})
}

func (r *MealRepository) ListMealsCount(ctx context.Context) (int64, error) {
	return r.queries.ListMealsCount(ctx)
}

func (r *MealRepository) CreateMeal(ctx context.Context, name string, categoryID int32, servings int32) (sqlc.Meal, error) {
	return r.queries.CreateMeal(ctx, sqlc.CreateMealParams{Name: name, MealCategoryID: categoryID, Servings: servings})
}

func (r *MealRepository) UpdateMeal(ctx context.Context, id int32, name string, categoryID int32, servings int32) (sqlc.Meal, error) {
	return r.queries.UpdateMeal(ctx, sqlc.UpdateMealParams{Name: name, MealCategoryID: categoryID, Servings: servings, ID: id})
}

func (r *MealRepository) DeleteMeal(ctx context.Context, id int32) error {
	return r.queries.DeleteMeal(ctx, id)
}

func (r *MealRepository) ListMealsByCategory(ctx context.Context, categoryID int32, limit int32, offset int32) ([]sqlc.Meal, error) {
	return r.queries.ListMealsByCategory(ctx, sqlc.ListMealsByCategoryParams{MealCategoryID: categoryID, Limit: limit, Offset: offset})
}

func (r *MealRepository) ListMealsByCategoryCount(ctx context.Context, categoryID int32) (int64, error) {
	return r.queries.ListMealsByCategoryCount(ctx, categoryID)
}

func (r *MealRepository) SearchMeals(ctx context.Context, pattern string, limit int32, offset int32) ([]sqlc.Meal, error) {
	return r.queries.SearchMeals(ctx, sqlc.SearchMealsParams{Name: "%" + pattern + "%", Limit: limit, Offset: offset})
}

func (r *MealRepository) SearchMealsCount(ctx context.Context, pattern string) (int64, error) {
	return r.queries.SearchMealsCount(ctx, "%"+pattern+"%")
}

// MealIngredientRepository wraps sqlc queries for meal ingredients
type MealIngredientRepository struct {
	queries *sqlc.Queries
}

func NewMealIngredientRepository(q *sqlc.Queries) *MealIngredientRepository {
	return &MealIngredientRepository{queries: q}
}

func (r *MealIngredientRepository) GetMealIngredient(ctx context.Context, id int32) (sqlc.MealIngredient, error) {
	return r.queries.GetMealIngredient(ctx, id)
}

func (r *MealIngredientRepository) ListMealIngredients(ctx context.Context, mealID int32) ([]sqlc.MealIngredient, error) {
	return r.queries.ListMealIngredients(ctx, mealID)
}

func (r *MealIngredientRepository) CreateMealIngredient(ctx context.Context, mealID int32, ingredientID int32, unitID sql.NullInt32, requiredAmount sql.NullString, isSeasoning bool, depletesSlowly bool, altIngredientID sql.NullInt32) (sqlc.MealIngredient, error) {
	return r.queries.CreateMealIngredient(ctx, sqlc.CreateMealIngredientParams{
		MealID:                mealID,
		IngredientID:          ingredientID,
		UnitID:                unitID,
		RequiredAmount:        requiredAmount,
		IsSeasoning:           isSeasoning,
		DepletesSlowly:        depletesSlowly,
		AlternateIngredientID: altIngredientID,
	})
}

func (r *MealIngredientRepository) UpdateMealIngredient(ctx context.Context, id int32, ingredientID int32, unitID sql.NullInt32, requiredAmount sql.NullString, isSeasoning bool, depletesSlowly bool, altIngredientID sql.NullInt32) (sqlc.MealIngredient, error) {
	return r.queries.UpdateMealIngredient(ctx, sqlc.UpdateMealIngredientParams{
		IngredientID:          ingredientID,
		UnitID:                unitID,
		RequiredAmount:        requiredAmount,
		IsSeasoning:           isSeasoning,
		DepletesSlowly:        depletesSlowly,
		AlternateIngredientID: altIngredientID,
		ID:                    id,
	})
}

func (r *MealIngredientRepository) DeleteMealIngredient(ctx context.Context, id int32) error {
	return r.queries.DeleteMealIngredient(ctx, id)
}

func (r *MealIngredientRepository) DeleteMealIngredientsByMealId(ctx context.Context, mealID int32) error {
	return r.queries.DeleteMealIngredientsByMealId(ctx, mealID)
}
