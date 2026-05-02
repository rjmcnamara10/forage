package repository

import (
	"context"

	"github.com/rjmcnamara10/forage/db/sqlc"
)

// ShoppingListRepository wraps sqlc queries for shopping lists
type ShoppingListRepository struct {
	queries *sqlc.Queries
}

func NewShoppingListRepository(q *sqlc.Queries) *ShoppingListRepository {
	return &ShoppingListRepository{queries: q}
}

func (r *ShoppingListRepository) GetShoppingList(ctx context.Context, id int32) (sqlc.ShoppingList, error) {
	return r.queries.GetShoppingList(ctx, id)
}

func (r *ShoppingListRepository) ListShoppingLists(ctx context.Context, limit int32, offset int32) ([]sqlc.ShoppingList, error) {
	return r.queries.ListShoppingLists(ctx, sqlc.ListShoppingListsParams{Limit: limit, Offset: offset})
}

func (r *ShoppingListRepository) ListShoppingListsCount(ctx context.Context) (int64, error) {
	return r.queries.ListShoppingListsCount(ctx)
}

func (r *ShoppingListRepository) CreateShoppingList(ctx context.Context, name string) (sqlc.ShoppingList, error) {
	return r.queries.CreateShoppingList(ctx, name)
}

func (r *ShoppingListRepository) UpdateShoppingList(ctx context.Context, id int32, name string) (sqlc.ShoppingList, error) {
	return r.queries.UpdateShoppingList(ctx, sqlc.UpdateShoppingListParams{ID: id, Name: name})
}

func (r *ShoppingListRepository) DeleteShoppingList(ctx context.Context, id int32) error {
	return r.queries.DeleteShoppingList(ctx, id)
}

// ShoppingListItemRepository wraps sqlc queries for shopping list items
type ShoppingListItemRepository struct {
	queries *sqlc.Queries
}

func NewShoppingListItemRepository(q *sqlc.Queries) *ShoppingListItemRepository {
	return &ShoppingListItemRepository{queries: q}
}

func (r *ShoppingListItemRepository) GetShoppingListItem(ctx context.Context, id int32) (sqlc.ShoppingListItem, error) {
	return r.queries.GetShoppingListItem(ctx, id)
}

func (r *ShoppingListItemRepository) ListShoppingListItems(ctx context.Context, shoppingListID int32) ([]sqlc.ShoppingListItem, error) {
	return r.queries.ListShoppingListItems(ctx, shoppingListID)
}

func (r *ShoppingListItemRepository) CreateShoppingListItem(ctx context.Context, params sqlc.CreateShoppingListItemParams) (sqlc.ShoppingListItem, error) {
	return r.queries.CreateShoppingListItem(ctx, params)
}

func (r *ShoppingListItemRepository) UpdateShoppingListItem(ctx context.Context, params sqlc.UpdateShoppingListItemParams) (sqlc.ShoppingListItem, error) {
	return r.queries.UpdateShoppingListItem(ctx, params)
}

func (r *ShoppingListItemRepository) DeleteShoppingListItem(ctx context.Context, id int32) error {
	return r.queries.DeleteShoppingListItem(ctx, id)
}

func (r *ShoppingListItemRepository) DeleteShoppingListItemsByShoppingListId(ctx context.Context, shoppingListID int32) error {
	return r.queries.DeleteShoppingListItemsByShoppingListId(ctx, shoppingListID)
}
