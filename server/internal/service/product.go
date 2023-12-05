package service

import (
	"context"
	"errors"

	"github.com/artemiyKew/json-rpc-lamoda/internal/repo"
	"github.com/artemiyKew/json-rpc-lamoda/internal/types"
)

type ProductService struct {
	productRepo     repo.Product
	shippingRepo    repo.Shipping
	reservationRepo repo.Reservation
}

func NewProductService(productRepo repo.Product, shippingRepo repo.Shipping, reservationRepo repo.Reservation) *ProductService {
	return &ProductService{
		productRepo:     productRepo,
		shippingRepo:    shippingRepo,
		reservationRepo: reservationRepo,
	}
}

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

	if err := s.productRepo.CreateProduct(ctx, types.Product{
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
	products, err := s.productRepo.GetProducts(ctx)
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
	products, err := s.shippingRepo.GetUnreservedProductsByWarehouseID(ctx, warehouseID)
	if err != nil {
		return nil, err
	}
	for _, product := range products {

		output = append(output, &ProductOutput{
			ID:         product.ID,
			UniqueCode: product.UniqueCode,
			Quantity:   product.Quantity,
		})
	}
	return output, nil
}

func (s *ProductService) CreateReserve(ctx context.Context, uniqueCodes []string) error {
	uniqueCodesMap := make(map[string]int, len(uniqueCodes))

	for _, uniqueCode := range uniqueCodes {
		uniqueCodesMap[uniqueCode]++
	}

	for code, count := range uniqueCodesMap {
		if err := s.productRepo.ReserveProduct(ctx, code, count); err != nil {
			return err
		}
		for count > 0 {
			shipping, err := s.shippingRepo.ReserveProduct(ctx, code, &count)
			if err != nil {
				return err
			}

			if err := s.reservationRepo.CreateReservation(ctx, types.Reservation{
				UniqueCode:  shipping.UniqueCode,
				WarehouseID: shipping.WarehouseID,
				Quantity:    shipping.Quantity,
				Status:      types.ReservationStatusReserved,
			}); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *ProductService) CancelReservation(ctx context.Context, uniqueCodes []string) error {
	uniqueCodesMap := make(map[string]int, len(uniqueCodes))

	for _, uniqueCode := range uniqueCodes {
		uniqueCodesMap[uniqueCode]++
	}
	for code := range uniqueCodesMap {
		reservations, err := s.reservationRepo.CancelReservation(ctx, types.Reservation{
			UniqueCode: code,
			Status:     types.ReservationStatusCanceled,
		})
		if err != nil {
			return err
		}
		for _, reservation := range reservations {
			if err := s.shippingRepo.CancelReservationProduct(ctx, reservation.UniqueCode, reservation.WarehouseID, reservation.Quantity); err != nil {
				return err
			}
			if err := s.productRepo.CancelReservationProduct(ctx, reservation.UniqueCode, reservation.Quantity); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *ProductService) GetAllReservations(ctx context.Context) (ProductReservationOutput, error) {
	rs, err := s.reservationRepo.GetAllReservations(ctx)
	if err != nil {
		return ProductReservationOutput{}, err
	}
	reservations := make([]ProductReservation, 0)
	for _, r := range rs {
		reservations = append(reservations, ProductReservation{
			ID:          r.ID,
			WarehouseID: r.WarehouseID,
			Quantity:    r.Quantity,
			UniqueCode:  r.UniqueCode,
			Status:      r.Status,
		})
	}
	return ProductReservationOutput{
		ProductReservations: reservations,
	}, nil
}
