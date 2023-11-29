package service

import (
	"context"
	"errors"

	"github.com/artemiyKew/json-rpc-lamoda/internal/repo"
	"github.com/artemiyKew/json-rpc-lamoda/internal/types"
)

type WarehouseService struct {
	warehouseRepo repo.Warehouse
}

func NewWarehouseService(warehouseRepo repo.Warehouse) *WarehouseService {
	return &WarehouseService{
		warehouseRepo: warehouseRepo,
	}
}

func (s *WarehouseService) CreateWarehouse(ctx context.Context, input WarehouseCreateInput) error {
	if input.Name == "" {
		return errors.New("warehouse name is required")
	}

	if !input.Availability {
		input.Availability = false
	}

	err := s.warehouseRepo.CreateWarehouse(ctx, types.Warehouse{
		Name:         input.Name,
		Availability: input.Availability,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *WarehouseService) GetWarehouseByID(ctx context.Context, id int) (*WarehouseOutput, error) {
	if id <= 0 {
		return &WarehouseOutput{}, errors.New("invalid id")
	}

	output, err := s.warehouseRepo.GetWarehouseByID(ctx, id)
	if err != nil {
		return &WarehouseOutput{}, err
	}

	return &WarehouseOutput{
		ID:           output.ID,
		Name:         output.Name,
		Availability: output.Availability,
	}, nil
}

func (s *WarehouseService) GetWarehouseByName(ctx context.Context, name string) (*WarehouseOutput, error) {
	if name == "" {
		return &WarehouseOutput{}, errors.New("invalid name")
	}

	output, err := s.warehouseRepo.GetWarehouseByName(ctx, name)
	if err != nil {
		return &WarehouseOutput{}, err
	}

	return &WarehouseOutput{
		ID:           output.ID,
		Name:         output.Name,
		Availability: output.Availability,
	}, nil
}
