package delivery

import (
	"context"
	"fmt"

	"github.com/artemiyKew/json-rpc-lamoda/internal/service"
)

type WarehouseRoutes struct {
	warehouseService service.Warehouse
	ctx              context.Context
}

func NewWarehouseRoutes(ctx context.Context, warehouseService service.Warehouse) *WarehouseRoutes {
	return &WarehouseRoutes{
		warehouseService: warehouseService,
		ctx:              ctx,
	}
}

type WarehouseOutput struct {
	ID           int    `json:"warehouse_id"`
	Name         string `json:"name"`
	Availability bool   `json:"availability"`
}

type WarehouseCreateInput struct {
	Name         string `json:"name"`
	Availability bool   `json:"availability"`
}

func (r *WarehouseRoutes) Create(args WarehouseCreateInput, reply *string) error {
	err := r.warehouseService.CreateWarehouse(r.ctx, service.WarehouseCreateInput{
		Name:         args.Name,
		Availability: args.Availability,
	})
	if err != nil {
		return err
	}

	*reply = fmt.Sprintf("warehouse created: %s, %v", args.Name, args.Availability)
	return nil
}

type WarehouseGetByIDInput struct {
	ID int `json:"warehouse_id"`
}

func (r *WarehouseRoutes) GetByID(args WarehouseGetByIDInput, reply *WarehouseOutput) error {
	output, err := r.warehouseService.GetWarehouseByID(r.ctx, args.ID)
	if err != nil {
		return err
	}

	*reply = WarehouseOutput{
		ID:           output.ID,
		Name:         output.Name,
		Availability: output.Availability,
	}
	return nil
}

type WarehouseGetByNameInput struct {
	Name string `json:"name"`
}

func (r *WarehouseRoutes) GetByName(args WarehouseGetByNameInput, reply *WarehouseOutput) error {
	output, err := r.warehouseService.GetWarehouseByName(r.ctx, args.Name)
	if err != nil {
		return err
	}

	*reply = WarehouseOutput{
		ID:           output.ID,
		Name:         output.Name,
		Availability: output.Availability,
	}
	return nil
}
