package delivery

import (
	"context"
	"fmt"

	"github.com/artemiyKew/json-rpc-lamoda/internal/service"
)

type ProductRoutes struct {
	productService service.Product
	ctx            context.Context
}

func NewProductRoutes(ctx context.Context, productService service.Product) *ProductRoutes {
	return &ProductRoutes{
		productService: productService,
		ctx:            ctx,
	}
}

type ProductOutput struct {
	ID         int    `json:"product_id"`
	Name       string `json:"name"`
	Size       string `json:"size"`
	UniqueCode string `json:"unique_code"`
	Quantity   int    `json:"quantity"`
}

type ProductsOutput struct {
	Products []ProductOutput `json:"products"`
}

type ProductCreateInput struct {
	Name       string `json:"name"`
	Size       string `json:"size"`
	UniqueCode string `json:"unique_code"`
}

func (r *ProductRoutes) Create(args ProductCreateInput, reply *string) error {
	err := r.productService.CreateProduct(r.ctx, service.ProductCreateInput{
		Name:       args.Name,
		Size:       args.Size,
		UniqueCode: args.UniqueCode,
	})
	if err != nil {
		return err
	}

	*reply = fmt.Sprintf("product created: %s, %s, %s", args.Name, args.Size, args.UniqueCode)

	return nil
}

type ProductGetUnreservedByWarehouseID struct {
	WarehouseID int `json:"warehouse_id"`
}

func (r *ProductRoutes) GetUnreservedProductsByWarehouseID(args ProductGetUnreservedByWarehouseID, reply *ProductsOutput) error {
	output, err := r.productService.GetUnreservedProductsByWarehouseID(r.ctx, args.WarehouseID)
	if err != nil {
		return err
	}
	products := make([]ProductOutput, len(output))
	for i, product := range output {
		products[i] = ProductOutput{
			ID:         product.ID,
			Name:       product.Name,
			Size:       product.Size,
			UniqueCode: product.UniqueCode,
			Quantity:   product.Quantity,
		}
	}
	reply.Products = products
	return nil
}

type ProductReserve struct {
	UniqueCodes []string `json:"unique_codes"`
}

func (r *ProductRoutes) CreateReserve(args ProductReserve, reply *string) error {
	if err := r.productService.CreateReserve(r.ctx, args.UniqueCodes); err != nil {
		return err
	}
	*reply = "product reserved"
	return nil
}

func (r *ProductRoutes) CancelReservation(args ProductReserve, reply *string) error {
	if err := r.productService.CancelReservation(r.ctx, args.UniqueCodes); err != nil {
		return err
	}
	*reply = "cancel reservation"
	return nil
}

type GetAllReservationsInput struct{}

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

func (r *ProductRoutes) GetAllReservations(args GetAllReservationsInput, reply *ProductReservationOutput) error {
	output, err := r.productService.GetAllReservations(r.ctx)
	if err != nil {
		return err
	}
	reservations := make([]ProductReservation, len(output.ProductReservations))
	for i, reservation := range output.ProductReservations {
		reservations[i] = ProductReservation{
			ID:          reservation.ID,
			WarehouseID: reservation.WarehouseID,
			UniqueCode:  reservation.UniqueCode,
			Quantity:    reservation.Quantity,
			Status:      reservation.Status,
		}
	}
	reply.ProductReservations = reservations
	return nil
}
