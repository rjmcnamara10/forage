package repository

import (
	"context"
	"github.com/rjmcnamara10/forage/db/sqlc"
)

// GroceryRepository wraps sqlc queries for grocery list operations
type GroceryRepository struct {
	queries *sqlc.Queries
}

func NewGroceryRepository(q *sqlc.Queries) *GroceryRepository {
	return &GroceryRepository{queries: q}
}

func (r *GroceryRepository) GetMealIngredientsForIds(ctx context.Context, mealIds []int32) ([]sqlc.GetMealIngredientsForIdsRow, error) {
	return r.queries.GetMealIngredientsForIds(ctx, mealIds)
}

func (r *GroceryRepository) GetMealServings(ctx context.Context, id int32) (int32, error) {
	return r.queries.GetMealServings(ctx, id)
}
