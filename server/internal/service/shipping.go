package service

import (
	"context"

	"github.com/artemiyKew/json-rpc-lamoda/internal/repo"
)

type ShippingService struct {
	shippingRepo repo.Shipping
}

func NewShippingService(shippingRepo repo.Shipping) *ShippingService {
	return &ShippingService{
		shippingRepo: shippingRepo,
	}
}

func (s *ShippingService) CreateShipping(ctx context.Context, shipping ShippingCreateInput) error {
	return nil
}
