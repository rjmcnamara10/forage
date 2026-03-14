package repository

import (
	"context"
	"github.com/rjmcnamara10/forage/db/sqlc"
)

// ItemCategoryRepository wraps sqlc queries for item categories
type ItemCategoryRepository struct {
	queries *sqlc.Queries
}

func NewItemCategoryRepository(q *sqlc.Queries) *ItemCategoryRepository {
	return &ItemCategoryRepository{queries: q}
}

func (r *ItemCategoryRepository) GetItemCategory(ctx context.Context, id int32) (sqlc.ItemCategory, error) {
	return r.queries.GetItemCategory(ctx, id)
}

func (r *ItemCategoryRepository) GetItemCategoryByName(ctx context.Context, name string) (sqlc.ItemCategory, error) {
	return r.queries.GetItemCategoryByName(ctx, name)
}

func (r *ItemCategoryRepository) ListItemCategories(ctx context.Context) ([]sqlc.ItemCategory, error) {
	return r.queries.ListItemCategories(ctx)
}

func (r *ItemCategoryRepository) CreateItemCategory(ctx context.Context, name string) (sqlc.ItemCategory, error) {
	return r.queries.CreateItemCategory(ctx, name)
}

func (r *ItemCategoryRepository) UpdateItemCategory(ctx context.Context, id int32, name string) error {
	return r.queries.UpdateItemCategory(ctx, sqlc.UpdateItemCategoryParams{ID: id, Name: name})
}

func (r *ItemCategoryRepository) DeleteItemCategory(ctx context.Context, id int32) error {
	return r.queries.DeleteItemCategory(ctx, id)
}

// MealCategoryRepository wraps sqlc queries for meal categories
type MealCategoryRepository struct {
	queries *sqlc.Queries
}

func NewMealCategoryRepository(q *sqlc.Queries) *MealCategoryRepository {
	return &MealCategoryRepository{queries: q}
}

func (r *MealCategoryRepository) GetMealCategory(ctx context.Context, id int32) (sqlc.MealCategory, error) {
	return r.queries.GetMealCategory(ctx, id)
}

func (r *MealCategoryRepository) GetMealCategoryByName(ctx context.Context, name string) (sqlc.MealCategory, error) {
	return r.queries.GetMealCategoryByName(ctx, name)
}

func (r *MealCategoryRepository) ListMealCategories(ctx context.Context) ([]sqlc.MealCategory, error) {
	return r.queries.ListMealCategories(ctx)
}

func (r *MealCategoryRepository) CreateMealCategory(ctx context.Context, name string) (sqlc.MealCategory, error) {
	return r.queries.CreateMealCategory(ctx, name)
}

func (r *MealCategoryRepository) UpdateMealCategory(ctx context.Context, id int32, name string) error {
	return r.queries.UpdateMealCategory(ctx, sqlc.UpdateMealCategoryParams{ID: id, Name: name})
}

func (r *MealCategoryRepository) DeleteMealCategory(ctx context.Context, id int32) error {
	return r.queries.DeleteMealCategory(ctx, id)
}
