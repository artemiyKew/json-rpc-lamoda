package postgres

import (
	"context"

	"github.com/artemiyKew/json-rpc-lamoda/internal/types"
	"github.com/sirupsen/logrus"
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
		logrus.Fatalf("ShippingRepo:CreateShipping %s", err)

		return err
	}
	return nil
}

func (r *ShippingRepo) GetUnreservedProductsByWarehouseID(ctx context.Context, warehouseID int) ([]types.Shipping, error) {
	query := `SELECT * FROM Shipping WHERE warehouse_id = $1 AND quantity > 0`
	p := make([]types.Shipping, 0)
	if err := r.db.SelectContext(ctx, &p, query, warehouseID); err != nil {
		logrus.Fatalf("ShippingRepo:GetUnreservedProductsByWarehouseID %s", err)

		return nil, err
	}

	return p, nil
}

func (r *ShippingRepo) ReserveProduct(ctx context.Context, uniqueCode string, quantity *int) (types.Shipping, error) {
	query := `SELECT quantity FROM Shipping WHERE unique_code = $1 AND quantity > 0 LIMIT 1`
	var q int

	var shipping types.Shipping
	if err := r.db.QueryRowContext(ctx, query, uniqueCode).Scan(&q); err != nil {
		logrus.Fatalf("ShippingRepo:ReserveProduct1 %s", err)
		return types.Shipping{}, err
	}

	if q >= *quantity {
		query = `UPDATE Shipping SET quantity = quantity - $1 WHERE unique_code = $2 AND quantity = $3 RETURNING *`
		if err := r.db.GetContext(ctx, &shipping, query, *quantity, uniqueCode, q); err != nil {
			logrus.Fatalf("ShippingRepo:ReserveProduct2 %s", err)
			return types.Shipping{}, err
		}

		shipping.Quantity = *quantity
		*quantity -= *quantity
	} else {
		query = `UPDATE Shipping SET quantity = 0 WHERE unique_code = $1 AND quantity = $2 RETURNING *`
		if err := r.db.GetContext(ctx, &shipping, query, uniqueCode, q); err != nil {
			logrus.Fatalf("ShippingRepo:ReserveProduct3 %s", err)
			return types.Shipping{}, err
		}

		*quantity = *quantity - q
		shipping.Quantity = q

	}

	return shipping, nil
}

func (r *ShippingRepo) CancelReservationProduct(ctx context.Context, uniqueCode string, warehouseID, quantity int) error {
	query := `UPDATE Shipping SET quantity = quantity + $1 WHERE unique_code = $2 AND warehouse_id = $3`
	if err := r.db.QueryRowContext(ctx, query, quantity, uniqueCode, warehouseID).Err(); err != nil {
		logrus.Fatalf("ShippingRepo:CancelReservationProduct %s", err)
		return err
	}
	return nil
}
