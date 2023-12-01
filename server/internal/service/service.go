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

type ProductReservation struct {
	ID          int    `json:"reservation_id"`
	WarehouseID int    `json:"warehouse_id"`
	UniqueCode  string `json:"unique_code"`
	Quantity    int    `json:"quantity"`
	Status      string `json:"status"`
}

type ProductReservationOutput struct {
	ProductReservations []ProductReservation `json:"product_reservations"`
}

type Product interface {
	CreateProduct(context.Context, ProductCreateInput) error
	GetProducts(context.Context) ([]*ProductOutput, error)
	GetUnreservedProductsByWarehouseID(context.Context, int) ([]*ProductOutput, error)
	CreateReserve(context.Context, []string) error
	CancelReservation(context.Context, []string) error
	GetAllReservations(ctx context.Context) (ProductReservationOutput, error)
}

type Services struct {
	Warehouse Warehouse
	Product   Product
}

type ServicesDependencies struct {
	Repos *repo.Repositories
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		Warehouse: NewWarehouseService(deps.Repos.Warehouse),
		Product:   NewProductService(deps.Repos.Product, deps.Repos.Shipping, deps.Repos.Reservation),
	}
}
