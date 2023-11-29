package service

import (
	"context"

	"github.com/artemiyKew/json-rpc-lamoda/internal/repo"
)

type WarehouseCreateInput struct {
	Name         string
	Availability bool
}

type WarehouseOutput struct {
	ID           int    `json:"warehouse_id"`
	Name         string `json:"name"`
	Availability bool   `json:"availability"`
}

type Warehouse interface {
	CreateWarehouse(context.Context, WarehouseCreateInput) error
	GetWarehouseByID(context.Context, int) (*WarehouseOutput, error)
	GetWarehouseByName(context.Context, string) (*WarehouseOutput, error)
}

type ProductCreateInput struct {
	Name       string
	Size       string
	UniqueCode string
}

type ProductOutput struct {
	ID         int    `json:"product_id"`
	Name       string `json:"name"`
	Size       string `json:"size"`
	UniqueCode string `json:"unique_code"`
	Quantity   int    `json:"quantity"`
}

type Product interface {
	// Postgres
	CreateProduct(context.Context, ProductCreateInput) error
	GetProducts(context.Context) ([]*ProductOutput, error)
	GetUnreservedProductsByWarehouseID(context.Context, int) ([]*ProductOutput, error)

	// Redis
	ReserveProduct(context.Context, []string) error
	CancelReservationProduct(context.Context, []string) error
	IsProductReserved(context.Context, string) (bool, error)
}

type ShippingCreateInput struct {
	UniqueCode  string
	WarehouseID int
	Quantity    int
}

type ShippingOutput struct {
	ID          int    `json:"shipping_id"`
	UniqueCode  string `json:"unique_code"`
	WarehouseID int    `json:"warehouse_id"`
	Quantity    int    `json:"quantity"`
}

type Shipping interface {
	CreateShipping(context.Context, ShippingCreateInput) error
}

type Services struct {
	Warehouse Warehouse
	Product   Product
	Shipping  Shipping
}

type ServicesDependencies struct {
	Repos *repo.Repositories
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		Warehouse: NewWarehouseService(deps.Repos.Warehouse),
		Product:   NewProductService(deps.Repos.Product, deps.Repos.ProductRedis, deps.Repos.Shipping),
		Shipping:  NewShippingService(deps.Repos.Shipping),
	}
}
