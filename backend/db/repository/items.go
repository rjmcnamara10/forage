package repository

import (
	"context"
	"database/sql"
	"github.com/rjmcnamara10/forage/db/sqlc"
)

// ItemRepository wraps sqlc queries for items
type ItemRepository struct {
	queries *sqlc.Queries
}

func NewItemRepository(q *sqlc.Queries) *ItemRepository {
	return &ItemRepository{queries: q}
}

func (r *ItemRepository) GetItem(ctx context.Context, id int32) (sqlc.Item, error) {
	return r.queries.GetItem(ctx, id)
}

func (r *ItemRepository) GetItemByName(ctx context.Context, name string) (sqlc.Item, error) {
	return r.queries.GetItemByName(ctx, name)
}

func (r *ItemRepository) ListItems(ctx context.Context, limit int32, offset int32) ([]sqlc.Item, error) {
	return r.queries.ListItems(ctx, sqlc.ListItemsParams{Limit: limit, Offset: offset})
}

func (r *ItemRepository) ListItemsCount(ctx context.Context) (int64, error) {
	return r.queries.ListItemsCount(ctx)
}

func (r *ItemRepository) CreateItem(ctx context.Context, name string) (sqlc.Item, error) {
	return r.queries.CreateItem(ctx, name)
}

func (r *ItemRepository) UpdateItem(ctx context.Context, id int32, name string) error {
	return r.queries.UpdateItem(ctx, sqlc.UpdateItemParams{ID: id, Name: name})
}

func (r *ItemRepository) DeleteItem(ctx context.Context, id int32) error {
	return r.queries.DeleteItem(ctx, id)
}

func (r *ItemRepository) GetItemCategories(ctx context.Context, itemID int32) ([]sqlc.ItemCategory, error) {
	return r.queries.GetItemCategories(ctx, itemID)
}

func (r *ItemRepository) AddItemCategory(ctx context.Context, itemID int32, categoryID int32) error {
	return r.queries.AddItemCategory(ctx, sqlc.AddItemCategoryParams{ItemID: itemID, CategoryID: categoryID})
}

func (r *ItemRepository) RemoveItemCategory(ctx context.Context, itemID int32, categoryID int32) error {
	return r.queries.RemoveItemCategory(ctx, sqlc.RemoveItemCategoryParams{ItemID: itemID, CategoryID: categoryID})
}

func (r *ItemRepository) RemoveAllItemCategories(ctx context.Context, itemID int32) error {
	return r.queries.RemoveAllItemCategories(ctx, itemID)
}

func (r *ItemRepository) ListItemsByCategory(ctx context.Context, categoryID int32, limit int32, offset int32) ([]sqlc.Item, error) {
	return r.queries.ListItemsByCategory(ctx, sqlc.ListItemsByCategoryParams{CategoryID: categoryID, Limit: limit, Offset: offset})
}

func (r *ItemRepository) ListItemsByCategoryCount(ctx context.Context, categoryID int32) (int64, error) {
	return r.queries.ListItemsByCategoryCount(ctx, categoryID)
}

func (r *ItemRepository) SearchItems(ctx context.Context, pattern string, limit int32, offset int32) ([]sqlc.Item, error) {
	return r.queries.SearchItems(ctx, sqlc.SearchItemsParams{Name: "%" + pattern + "%", Limit: limit, Offset: offset})
}

func (r *ItemRepository) SearchItemsCount(ctx context.Context, pattern string) (int64, error) {
	return r.queries.SearchItemsCount(ctx, "%"+pattern+"%")
}

// InventoryRepository wraps sqlc queries for inventory items
type InventoryRepository struct {
	queries *sqlc.Queries
}

func NewInventoryRepository(q *sqlc.Queries) *InventoryRepository {
	return &InventoryRepository{queries: q}
}

func (r *InventoryRepository) GetInventoryItem(ctx context.Context, id int32) (sqlc.InventoryItem, error) {
	return r.queries.GetInventoryItem(ctx, id)
}

func (r *InventoryRepository) GetInventoryItemByItemId(ctx context.Context, itemID int32) (sqlc.InventoryItem, error) {
	return r.queries.GetInventoryItemByItemId(ctx, itemID)
}

func (r *InventoryRepository) ListInventoryItems(ctx context.Context, limit int32, offset int32) ([]sqlc.InventoryItem, error) {
	return r.queries.ListInventoryItems(ctx, sqlc.ListInventoryItemsParams{Limit: limit, Offset: offset})
}

func (r *InventoryRepository) ListInventoryItemsCount(ctx context.Context) (int64, error) {
	return r.queries.ListInventoryItemsCount(ctx)
}

func (r *InventoryRepository) CreateInventoryItem(ctx context.Context, itemID int32, unitID sql.NullInt32, storedAmount sql.NullString) (sqlc.InventoryItem, error) {
	return r.queries.CreateInventoryItem(ctx, sqlc.CreateInventoryItemParams{ItemID: itemID, UnitID: unitID, StoredAmount: storedAmount})
}

func (r *InventoryRepository) UpdateInventoryItem(ctx context.Context, itemID int32, unitID sql.NullInt32, storedAmount sql.NullString) error {
	return r.queries.UpdateInventoryItem(ctx, sqlc.UpdateInventoryItemParams{UnitID: unitID, StoredAmount: storedAmount, ItemID: itemID})
}

func (r *InventoryRepository) UpdateInventoryStoredAmount(ctx context.Context, itemID int32, storedAmount sql.NullString) error {
	return r.queries.UpdateInventoryStoredAmount(ctx, sqlc.UpdateInventoryStoredAmountParams{StoredAmount: storedAmount, ItemID: itemID})
}

func (r *InventoryRepository) DeleteInventoryItem(ctx context.Context, itemID int32) error {
	return r.queries.DeleteInventoryItem(ctx, itemID)
}
