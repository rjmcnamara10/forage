package repository

import (
	"context"
	"github.com/rjmcnamara10/forage/db/sqlc"
)

// UnitRepository wraps sqlc queries for units
type UnitRepository struct {
	queries *sqlc.Queries
}

func NewUnitRepository(q *sqlc.Queries) *UnitRepository {
	return &UnitRepository{queries: q}
}

func (r *UnitRepository) GetUnit(ctx context.Context, id int32) (sqlc.Unit, error) {
	return r.queries.GetUnit(ctx, id)
}

func (r *UnitRepository) GetUnitByName(ctx context.Context, name string) (sqlc.Unit, error) {
	return r.queries.GetUnitByName(ctx, name)
}

func (r *UnitRepository) ListUnits(ctx context.Context) ([]sqlc.Unit, error) {
	return r.queries.ListUnits(ctx)
}

func (r *UnitRepository) CreateUnit(ctx context.Context, name string) (sqlc.Unit, error) {
	return r.queries.CreateUnit(ctx, name)
}

func (r *UnitRepository) UpdateUnit(ctx context.Context, id int32, name string) error {
	return r.queries.UpdateUnit(ctx, sqlc.UpdateUnitParams{ID: id, Name: name})
}

func (r *UnitRepository) DeleteUnit(ctx context.Context, id int32) error {
	return r.queries.DeleteUnit(ctx, id)
}
