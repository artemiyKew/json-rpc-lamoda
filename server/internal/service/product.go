package service

import (
	"context"
	"errors"

	"github.com/artemiyKew/json-rpc-lamoda/internal/repo"
	"github.com/artemiyKew/json-rpc-lamoda/internal/types"
)

type ProductService struct {
	productPGDBRepo repo.Product
	productRDBRepo  repo.ProductRedis
	shippingRepo    repo.Shipping
}

func NewProductService(productRepo repo.Product, productRDBRepo repo.ProductRedis, shippingRepo repo.Shipping) *ProductService {
	return &ProductService{
		productPGDBRepo: productRepo,
		productRDBRepo:  productRDBRepo,
		shippingRepo:    shippingRepo,
	}
}

// Postgres
func (s *ProductService) CreateProduct(ctx context.Context, input ProductCreateInput) error {
	if input.Name == "" {
		return errors.New("name is required")
	}

	if input.Size == "" {
		return errors.New("size is required")
	}

	if input.UniqueCode == "" {
		return errors.New("uniqueCode is required")
	}

	if err := s.productPGDBRepo.CreateProduct(ctx, types.Product{
		Name:       input.Name,
		Size:       input.Size,
		UniqueCode: input.UniqueCode,
	}); err != nil {
		return err
	}
	return nil
}

func (s *ProductService) GetProducts(ctx context.Context) ([]*ProductOutput, error) {
	output := make([]*ProductOutput, 0)
	products, err := s.productPGDBRepo.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		output = append(output, &ProductOutput{
			ID:         product.ID,
			Name:       product.Name,
			Size:       product.Size,
			UniqueCode: product.UniqueCode,
		})
	}
	return output, nil
}

func (s *ProductService) GetUnreservedProductsByWarehouseID(ctx context.Context, warehouseID int) ([]*ProductOutput, error) {
	output := make([]*ProductOutput, 0)

	products, err := s.productPGDBRepo.GetUnreservedProductsByWarehouseID(ctx, warehouseID)
	if err != nil {
		return nil, err
	}
	for _, product := range products {
		quantity, err := s.shippingRepo.GetQuantityByWarehouseIDUniqueCode(ctx, warehouseID, product.UniqueCode)
		if err != nil {
			return nil, err
		}
		output = append(output, &ProductOutput{
			ID:         product.ID,
			Name:       product.Name,
			Size:       product.Size,
			UniqueCode: product.UniqueCode,
			Quantity:   quantity,
		})
	}
	return output, nil
}

// Redis
func (s *ProductService) ReserveProduct(ctx context.Context, uniqueCodes []string) error {
	uniqueCodesMap := make(map[string]int, 0)
	for _, uniqueCode := range uniqueCodes {
		uniqueCodesMap[uniqueCode]++
	}
	for uniqueCode, countProduct := range uniqueCodesMap {
		productsQuantityInWarehouse, err := s.shippingRepo.GetQuantityByUniqueCode(ctx, uniqueCode)
		if err != nil {
			return err
		}
		if productsQuantityInWarehouse >= countProduct {
			if err := s.productRDBRepo.ReserveProduct(ctx, uniqueCode, countProduct); err != nil {
				return err
			}
		} else {
			return errors.New("there are less goods in stock than you ordered")
		}
	}
	return nil
}

func (s *ProductService) CancelReservationProduct(ctx context.Context, uniqueCodes []string) error {
	uniqueCodesMap := make(map[string]int, 0)
	for _, uniqueCode := range uniqueCodes {
		uniqueCodesMap[uniqueCode]++
	}
	for uniqueCode, countProduct := range uniqueCodesMap {

		if err := s.productRDBRepo.CancelReservationProduct(ctx, uniqueCode, countProduct); err != nil {
			return err
		}
	}
	return nil
}

func (s *ProductService) IsProductReserved(ctx context.Context, uniqueCode string) (bool, error) {
	ok, err := s.productRDBRepo.IsProductReserved(ctx, uniqueCode)
	if err != nil {
		return false, err
	}
	return ok, nil
}
