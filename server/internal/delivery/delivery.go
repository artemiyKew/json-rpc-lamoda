package delivery

import (
	"context"

	"github.com/artemiyKew/json-rpc-lamoda/internal/service"
)

type Routes struct {
	ProductRoutes   *ProductRoutes
	WarehouseRoutes *WarehouseRoutes
}

func NewRoutes(ctx context.Context, services *service.Services) *Routes {
	return &Routes{
		ProductRoutes:   NewProductRoutes(ctx, services.Product),
		WarehouseRoutes: NewWarehouseRoutes(ctx, services.Warehouse),
	}
}
