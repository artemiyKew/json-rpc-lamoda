package main

import (
	"log"
	"net/rpc/jsonrpc"

	"github.com/artemiyKew/client/json-rpc-lamoda/product"
)

func main() {
	// Создаем клиентское соединение с сервером JSON-RPC
	client, err := jsonrpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Error connecting to JSON-RPC server:", err)
	}
	defer client.Close()

	if err := product.ReserveProduct(client); err != nil {
		log.Fatal(err)
	}
	if err := product.GetAllReservations(client); err != nil {
		log.Fatal(err)
	}
	if err := product.GetUnreservedProductsByWarehouseID(client); err != nil {
		log.Fatal(err)
	}
	if err := product.CancelReservation(client); err != nil {
		log.Fatal(err)
	}
	if err := product.GetAllReservations(client); err != nil {
		log.Fatal(err)
	}
	if err := product.GetUnreservedProductsByWarehouseID(client); err != nil {
		log.Fatal(err)
	}

}
