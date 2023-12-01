package main

import (
	"log"
	"net/rpc/jsonrpc"

	"github.com/artemiyKew/client/json-rpc-lamoda/product"
	"github.com/sirupsen/logrus"
)

func main() {
	client, err := jsonrpc.Dial("tcp", ":1234")
	if err != nil {
		logrus.Fatal("Error connecting to JSON-RPC server: ", err)
	}
	defer func() {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if err := product.ReserveProduct(client); err != nil {
		logrus.Fatal(err)
	}
	if err := product.GetAllReservations(client); err != nil {
		logrus.Fatal(err)
	}
	if err := product.GetUnreservedProductsByWarehouseID(client); err != nil {
		logrus.Fatal(err)
	}
	if err := product.CancelReservation(client); err != nil {
		logrus.Fatal(err)
	}
	if err := product.GetAllReservations(client); err != nil {
		logrus.Fatal(err)
	}
	if err := product.GetUnreservedProductsByWarehouseID(client); err != nil {
		logrus.Fatal(err)
	}

}
