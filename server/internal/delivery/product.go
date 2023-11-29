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
}

type ProductsOutput struct {
	Products []*ProductOutput `json:"products"`
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

func (r *ProductRoutes) GetAll(args string, reply *ProductsOutput) error {
	products, err := r.productService.GetProducts(r.ctx)
	if err != nil {
		return err
	}

	for _, product := range products {
		reply.Products = append(reply.Products, &ProductOutput{
			ID:         product.ID,
			Name:       product.Name,
			UniqueCode: product.UniqueCode,
		})
	}
	return nil
}

type ProductGetUnreservedByWarehouseID struct {
	WarehouseID int `json:"warehouse_id"`
}

func (r *ProductRoutes) GetUnreservedProductsByWarehouseID(args ProductGetUnreservedByWarehouseID, reply *ProductsOutput) error {
	products, err := r.productService.GetUnreservedProductsByWarehouseID(r.ctx, args.WarehouseID)
	if err != nil {
		return err
	}

	for _, product := range products {
		reply.Products = append(reply.Products, &ProductOutput{
			ID:         product.ID,
			Name:       product.Name,
			UniqueCode: product.UniqueCode,
		})
	}
	return nil
}

type ProductReserve struct {
	UniqueCode []string `json:"unique_code"`
}

func (r *ProductRoutes) CreateReserve(args ProductReserve, reply *string) error {
	if err := r.productService.ReserveProduct(r.ctx, args.UniqueCode); err != nil {
		return err
	}
	*reply = fmt.Sprintf("product reserved, unique code: %s", args.UniqueCode)
	return nil
}

func (r *ProductRoutes) CancelReserve(args ProductReserve, reply *string) error {
	if err := r.productService.CancelReservationProduct(r.ctx, args.UniqueCode); err != nil {
		return err
	}
	*reply = fmt.Sprintf("cancel reservation for product with unique code: %s", args.UniqueCode)
	return nil
}

type ProductIsReserved struct {
	UniqueCode string `json:"unique_code"`
}

func (r *ProductRoutes) IsReserved(args ProductIsReserved, reply *string) error {
	ok, err := r.productService.IsProductReserved(r.ctx, args.UniqueCode)
	if err != nil {
		return err
	}
	if !ok {
		*reply = fmt.Sprintf("product: %s is not reserved", args.UniqueCode)
		return nil
	}
	*reply = fmt.Sprintf("cancel reservation for product with unique code: %s", args.UniqueCode)
	return nil
}
