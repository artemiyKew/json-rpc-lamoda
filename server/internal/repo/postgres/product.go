package postgres

import (
	"context"
	"errors"

	"github.com/artemiyKew/json-rpc-lamoda/internal/types"
	"github.com/sirupsen/logrus"
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
		logrus.Fatalf("ProductRepo:CreateProduct %s", err)
		return err
	}
	return nil
}

func (r *ProductRepo) GetProducts(ctx context.Context) (p []types.Product, err error) {
	query := `SELECT * FROM Product`
	p = make([]types.Product, 0)
	if err := r.db.SelectContext(ctx, &p, query); err != nil {
		logrus.Fatalf("ProductRepo:GetProducts %s", err)
		return p, err
	}

	return p, err
}

func (r *ProductRepo) ReserveProduct(ctx context.Context, uniqueCode string, countProducts int) (err error) {
	query := `UPDATE Product SET quantity = CASE WHEN quantity >= $1 THEN quantity - $1 ELSE quantity END WHERE quantity > 0 AND unique_code = $2 RETURNING *
	`
	var product types.Product
	if err := r.db.QueryRowxContext(ctx, query, countProducts, uniqueCode).StructScan(&product); err != nil {
		logrus.Fatalf("ProductRepo:ReserveProduct %s", err)
		return err
	}
	if product.Quantity == 0 {
		return errors.New("invalid quantity")
	}

	return err
}

func (r *ProductRepo) CancelReservationProduct(ctx context.Context, uniqueCode string, quantity int) error {
	query := `UPDATE Product SET quantity = quantity + $1 WHERE unique_code = $2`
	if err := r.db.QueryRowxContext(ctx, query, quantity, uniqueCode).Err(); err != nil {
		logrus.Fatalf("ProductRepo:CancelReservationProduct %s", err)
		return err
	}
	return nil
}
