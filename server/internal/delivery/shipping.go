package delivery

import (
	"context"

	"github.com/artemiyKew/json-rpc-lamoda/internal/service"
)

type ShippingRoutes struct {
	shippingService service.Shipping
	ctx             context.Context
}

func NewShippingRoutes(ctx context.Context, shippingService service.Shipping) *ShippingRoutes {
	return &ShippingRoutes{
		shippingService: shippingService,
		ctx:             ctx,
	}
}
