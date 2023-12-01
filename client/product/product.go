package product

import (
	"fmt"
	"net/rpc"

	"github.com/sirupsen/logrus"
)

type ProductReserve struct {
	UniqueCodes []string `json:"unique_codes"`
}

func ReserveProduct(client *rpc.Client) error {
	reserveInput := ProductReserve{
		UniqueCodes: []string{"Code3", "Code3", "Code3", "Code3", "Code3"},
	}
	var createReserveOutput string
	err := client.Call("ProductRoutes.CreateReserve", reserveInput, &createReserveOutput)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("CreateReserve Response: %+v\n\n", createReserveOutput)
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

func GetAllReservations(client *rpc.Client) error {
	reserveInput := GetAllReservationsInput{}
	var reservationsOutput ProductReservationOutput
	err := client.Call("ProductRoutes.GetAllReservations", reserveInput, &reservationsOutput)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("GetAllReservations Response: %+v\n\n", reservationsOutput)
	return nil
}

func CancelReservation(client *rpc.Client) error {
	reserveInput := ProductReserve{
		UniqueCodes: []string{"Code3"},
	}
	var cancelReserveOutput string
	err := client.Call("ProductRoutes.CancelReservation", reserveInput, &cancelReserveOutput)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("CancelReservation Response: %+v\n\n", cancelReserveOutput)
	return nil
}

type ProductGetUnreservedByWarehouseID struct {
	WarehouseID int `json:"warehouse_id"`
}

type ProductOutput struct {
	ID         int    `json:"product_id"`
	UniqueCode string `json:"unique_code"`
	Quantity   int    `json:"quantity"`
}

type ProductsOutput struct {
	Products []ProductOutput `json:"products"`
}

func GetUnreservedProductsByWarehouseID(client *rpc.Client) error {
	input := ProductGetUnreservedByWarehouseID{
		WarehouseID: 2,
	}
	var productsOutput ProductsOutput
	err := client.Call("ProductRoutes.GetUnreservedProductsByWarehouseID", input, &productsOutput)
	if err != nil {
		logrus.Info(err)
	}
	fmt.Printf("GetUnreservedProductsByWarehouseID Response: %+v\n", productsOutput)
	return nil
}
