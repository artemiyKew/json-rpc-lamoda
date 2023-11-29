package postgres

import (
	"context"

	"github.com/artemiyKew/json-rpc-lamoda/internal/types"
)

type WarehouseRepo struct {
	*PostgresDB
}

func NewWarehouseRepo(db *PostgresDB) *WarehouseRepo {
	return &WarehouseRepo{db}
}

func (r *WarehouseRepo) CreateWarehouse(ctx context.Context, wh types.Warehouse) error {
	query := `INSERT INTO Warehouse (name, availability) VALUES ($1, $2) RETURNING id`
	if err := r.db.QueryRowContext(ctx, query, wh.Name, wh.Availability).Scan(&wh.ID); err != nil {
		return err
	}
	return nil
}

func (r *WarehouseRepo) GetWarehouseByID(ctx context.Context, id int) (types.Warehouse, error) {
	query := `SELECT * FROM Warehouse WHERE id = $1`

	var wh types.Warehouse
	if err := r.db.QueryRowxContext(ctx, query, id).StructScan(&wh); err != nil {
		return types.Warehouse{}, err
	}

	return wh, nil
}

func (r *WarehouseRepo) GetWarehouseByName(ctx context.Context, name string) (types.Warehouse, error) {
	query := `SELECT * FROM Warehouse WHERE name = $1`

	var wh types.Warehouse
	if err := r.db.QueryRowxContext(ctx, query, name).StructScan(&wh); err != nil {
		return types.Warehouse{}, err
	}
	return wh, nil
}
