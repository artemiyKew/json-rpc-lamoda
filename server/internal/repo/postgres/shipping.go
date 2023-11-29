package postgres

import (
	"context"
	"errors"

	"github.com/artemiyKew/json-rpc-lamoda/internal/types"
)

type ShippingRepo struct {
	*PostgresDB
}

func NewShippingRepo(db *PostgresDB) *ShippingRepo {
	return &ShippingRepo{db}
}

func (r *ShippingRepo) CreateShipping(ctx context.Context, s types.Shipping) error {
	query := `INSERT INTO Shipping (unique_code, warehouse_id, quantity)`
	if err := r.db.QueryRowContext(ctx, query, s.UniqueCode, s.WarehouseID, s.Quantity).
		Scan(&s.ID); err != nil {
		return err
	}
	return nil
}

func (r *ShippingRepo) GetQuantityByWarehouseIDUniqueCode(ctx context.Context, warehouseID int, uniqueCode string) (int, error) {
	query := `SELECT quantity FROM Shipping WHERE wharehouse_id = $1 AND unique_code = $2`
	var quantity int
	if err := r.db.SelectContext(ctx, &quantity, query, warehouseID, uniqueCode); err != nil {
		return 0, err
	}
	return quantity, nil
}

// Get quantity from all warehouses that avaliable
func (r *ShippingRepo) GetQuantityByUniqueCode(ctx context.Context, uniqueCode string) (quantity int, err error) {
	query := `SELECT quantity FROM Shipping s JOIN Warehouse w ON s.warehouse_id = w.id WHERE s.unique_code = $1 AND w.availability = true`

	rows, err := r.db.QueryContext(ctx, query, uniqueCode)
	if err != nil {
		return quantity, err
	}

	defer func() {
		err = errors.Join(err, rows.Close())
	}()

	for rows.Next() {
		var num int
		if err := rows.Scan(&num); err != nil {
			return quantity, err
		}
		quantity += num
	}
	return quantity, err
}
