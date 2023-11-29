package repo

import (
	"context"

	"github.com/artemiyKew/json-rpc-lamoda/internal/repo/postgres"
	"github.com/artemiyKew/json-rpc-lamoda/internal/repo/redis"
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
	GetUnreservedProductsByWarehouseID(context.Context, int) ([]types.Product, error)
	ReserveProduct(context.Context, string, int) error
	CancelReservationProduct(context.Context, string, int, int) error
}

type ProductRedis interface {
	ReserveProduct(context.Context, string, int, int) error
	CancelReservationProduct(context.Context, string, int) error
	IsProductReserved(context.Context, string) (bool, error)
}

type Shipping interface {
	CreateShipping(context.Context, types.Shipping) error
	GetQuantityByWarehouseIDUniqueCode(context.Context, int, string) (int, error)
	GetQuantityByUniqueCode(context.Context, string) (int, error)
}

type Repositories struct {
	Warehouse
	Product
	ProductRedis
	Shipping
}

func NewRepositories(rdb *redis.RedisDB, db *postgres.PostgresDB) *Repositories {
	return &Repositories{
		Warehouse:    postgres.NewWarehouseRepo(db),
		Product:      postgres.NewProductRepo(db),
		ProductRedis: redis.NewProductRepo(rdb),
		Shipping:     postgres.NewShippingRepo(db),
	}
}
