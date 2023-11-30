package repo

import (
	"context"

	"github.com/artemiyKew/json-rpc-lamoda/internal/repo/postgres"
	"github.com/artemiyKew/json-rpc-lamoda/internal/types"
)

type Warehouse interface {
	CreateWarehouse(context.Context, types.Warehouse) error
	GetWarehouseByID(context.Context, int) (types.Warehouse, error)
	GetWarehouseByName(context.Context, string) (types.Warehouse, error)
}

type Product interface {
	CreateProduct(context.Context, types.Product) error
	GetProducts(context.Context) ([]types.Product, error)
	ReserveProduct(context.Context, string, int) error
	CancelReservationProduct(context.Context, string, int) error
}

type Shipping interface {
	CreateShipping(context.Context, types.Shipping) error
	ReserveProduct(context.Context, string, *int) (types.Shipping, error)
	GetUnreservedProductsByWarehouseID(context.Context, int) ([]types.Shipping, error)
	CancelReservationProduct(ctx context.Context, uniqueCode string, warehouseId int, quantity int) error
}

type Reservation interface {
	CreateReservation(context.Context, types.Reservation) error
	CancelReservation(context.Context, types.Reservation) ([]types.Reservation, error)

	GetAllReservations(context.Context) ([]types.Reservation, error)
}

type Repositories struct {
	Warehouse
	Product
	Shipping
	Reservation
}

func NewRepositories(db *postgres.PostgresDB) *Repositories {
	return &Repositories{
		Warehouse:   postgres.NewWarehouseRepo(db),
		Product:     postgres.NewProductRepo(db),
		Shipping:    postgres.NewShippingRepo(db),
		Reservation: postgres.NewReservationRepo(db),
	}
}
