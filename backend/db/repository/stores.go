package repository

import (
	"context"

	"github.com/rjmcnamara10/forage/db/sqlc"
)

// StoreRepository wraps sqlc queries for stores
type StoreRepository struct {
	queries *sqlc.Queries
}

func NewStoreRepository(q *sqlc.Queries) *StoreRepository {
	return &StoreRepository{queries: q}
}

func (r *StoreRepository) GetStore(ctx context.Context, id int32) (sqlc.Store, error) {
	return r.queries.GetStore(ctx, id)
}

func (r *StoreRepository) GetStoreByName(ctx context.Context, name string) (sqlc.Store, error) {
	return r.queries.GetStoreByName(ctx, name)
}

func (r *StoreRepository) ListStores(ctx context.Context, limit int32, offset int32, sortOrder string) ([]sqlc.Store, error) {
	if sortOrder == "DESC" {
		return r.queries.ListStoresDesc(ctx, sqlc.ListStoresDescParams{Limit: limit, Offset: offset})
	}
	return r.queries.ListStoresAsc(ctx, sqlc.ListStoresAscParams{Limit: limit, Offset: offset})
}

func (r *StoreRepository) ListStoresCount(ctx context.Context) (int64, error) {
	return r.queries.ListStoresCount(ctx)
}

func (r *StoreRepository) CreateStore(ctx context.Context, name string) (sqlc.Store, error) {
	return r.queries.CreateStore(ctx, name)
}

func (r *StoreRepository) UpdateStore(ctx context.Context, id int32, name string) (sqlc.Store, error) {
	return r.queries.UpdateStore(ctx, sqlc.UpdateStoreParams{ID: id, Name: name})
}

func (r *StoreRepository) DeleteStore(ctx context.Context, id int32) error {
	return r.queries.DeleteStore(ctx, id)
}

// StoreItemRepository wraps sqlc queries for store items
type StoreItemRepository struct {
	queries *sqlc.Queries
}

func NewStoreItemRepository(q *sqlc.Queries) *StoreItemRepository {
	return &StoreItemRepository{queries: q}
}

func (r *StoreItemRepository) GetStoreItem(ctx context.Context, id int32) (sqlc.StoreItem, error) {
	return r.queries.GetStoreItem(ctx, id)
}

func (r *StoreItemRepository) ListStoreItems(ctx context.Context, limit int32, offset int32) ([]sqlc.StoreItem, error) {
	return r.queries.ListStoreItems(ctx, sqlc.ListStoreItemsParams{Limit: limit, Offset: offset})
}

func (r *StoreItemRepository) ListStoreItemsCount(ctx context.Context) (int64, error) {
	return r.queries.ListStoreItemsCount(ctx)
}

func (r *StoreItemRepository) CreateStoreItem(ctx context.Context, params sqlc.CreateStoreItemParams) (sqlc.StoreItem, error) {
	return r.queries.CreateStoreItem(ctx, params)
}

func (r *StoreItemRepository) UpdateStoreItem(ctx context.Context, params sqlc.UpdateStoreItemParams) (sqlc.StoreItem, error) {
	return r.queries.UpdateStoreItem(ctx, params)
}

func (r *StoreItemRepository) DeleteStoreItem(ctx context.Context, id int32) error {
	return r.queries.DeleteStoreItem(ctx, id)
}

func (r *StoreItemRepository) ListStoreItemsByStore(ctx context.Context, storeID int32) ([]sqlc.StoreItem, error) {
	return r.queries.ListStoreItemsByStore(ctx, storeID)
}
