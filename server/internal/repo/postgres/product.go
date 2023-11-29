package postgres

import (
	"context"
	"errors"

	"github.com/artemiyKew/json-rpc-lamoda/internal/types"
)

type ProductRepo struct {
	*PostgresDB
}

func NewProductRepo(db *PostgresDB) *ProductRepo {
	return &ProductRepo{db}
}

func (r *ProductRepo) CreateProduct(ctx context.Context, p types.Product) error {
	query := `INSERT INTO Product (name, size, unique_code, quantity, warehouse_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	if err := r.db.QueryRowContext(ctx, query, p.Name, p.Size, p.UniqueCode).
		Scan(&p.ID); err != nil {
		return err
	}
	return nil
}

func (r *ProductRepo) GetProducts(ctx context.Context) (p []types.Product, err error) {
	query := `SELECT * FROM Product`
	p = make([]types.Product, 0)
	if err := r.db.SelectContext(ctx, &p, query); err != nil {
		return p, err
	}

	return p, err
}

func (r *ProductRepo) GetUnreservedProductsByWarehouseID(ctx context.Context, warehouseID int) (p []types.Product, err error) {
	query := `SELECT p.* FROM Product p JOIN Shipping s ON p.unique_code = s.unique_code WHERE s.warehouse_id = $1 AND s.quantity > 0`
	p = make([]types.Product, 0)

	rows, err := r.db.QueryContext(ctx, query, warehouseID)
	if err != nil {
		return p, err
	}

	defer func() {
		err = errors.Join(err, rows.Close())
	}()

	for rows.Next() {
		var product types.Product
		if err := rows.Scan(&product); err != nil {
			return p, err
		}
		p = append(p, product)
	}

	return p, err
}

func (r *ProductRepo) ReserveProduct(ctx context.Context, uniqueCode string, countProducts int) (err error) {
	query := `UPDATE Shipping SET quantity = quantity - $1 WHERE unique_code = $2 AND quantity > 0 AND warehouse_id IN (SELECT id FROM Warehouse WHERE availability = true)`
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			err = errors.Join(err, tx.Rollback())
			return
		}
		err = errors.Join(err, tx.Commit())
	}()

	_, err = tx.ExecContext(ctx, query, countProducts, uniqueCode)

	return err
}

func (r *ProductRepo) CancelReservationProduct(ctx context.Context, uniqueCode string, warehouse int, countProducts int) (err error) {
	query := `UPDATE Shipping SET quantity = quantity + $1 WHERE unique_code = $2 AND warehouse_id = $3 IN (SELECT id FROM Warehouse WHERE availability = true)`
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			err = errors.Join(err, tx.Rollback())
			return
		}
		err = errors.Join(err, tx.Commit())
	}()

	_, err = tx.ExecContext(ctx, query, countProducts, uniqueCode)

	return err
}
