package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type ProductCreateInput struct {
	Name        string `json:"name"`
	Size        string `json:"size"`
	UniqueCode  string `json:"unique_code"`
	Quantity    int    `json:"quantity"`
	WarehouseID int    `json:"warehouse_id"`
}

type ProductReserve struct {
	UniqueCode string `json:"unique_code"`
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

type WarehouseGetByIDInput struct {
	ID int `json:"warehouse_id"`
}

type WarehouseGetByNameInput struct {
	Name string `json:"name"`
}

func main() {
	// Создаем клиентское соединение с сервером JSON-RPC
	client, err := jsonrpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Error connecting to JSON-RPC server:", err)
	}
	defer client.Close()

	// Пример вызова метода Create
	// createInput := WarehouseCreateInput{
	// 	Name:         "Warehouse6",
	// 	Availability: true,
	// }
	// var createOutput WarehouseOutput
	// err = client.Call("WarehouseRoutes.Create", createInput, &createOutput)
	// if err != nil {
	// 	log.Fatal("Error calling Create method:", err)
	// }
	// fmt.Printf("Create method response: %+v\n", createOutput)

	// Пример вызова метода GetByID
	getByIDInput := WarehouseGetByIDInput{
		ID: 1,
	}
	var getByIDOutput WarehouseOutput
	err = client.Call("WarehouseRoutes.GetByID", getByIDInput, &getByIDOutput)
	if err != nil {
		log.Fatal("Error calling GetByID method:", err)
	}
	fmt.Printf("GetByID method response: %+v\n", getByIDOutput)

	// Пример вызова метода GetByName
	getByNameInput := WarehouseGetByNameInput{
		Name: "Warehouse1",
	}
	var getByNameOutput WarehouseOutput
	err = client.Call("WarehouseRoutes.GetByName", getByNameInput, &getByNameOutput)
	if err != nil {
		log.Fatal("Error calling GetByName method:", err)
	}
	fmt.Printf("GetByName method response: %+v\n", getByNameOutput)

	productInput := ProductCreateInput{
		Name:        "Sample Product",
		Size:        "Medium",
		UniqueCode:  "jacketmediumdd",
		Quantity:    10,
		WarehouseID: 1,
	}
	var productCreateOutput string
	err = client.Call("ProductRoutes.Create", productInput, &productCreateOutput)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Create Response:", productCreateOutput)

	// Пример использования метода CreateReserve
	reserveInput := ProductReserve{
		UniqueCode: "jacketmedium",
	}
	var createReserveOutput string
	err = client.Call("ProductRoutes.CreateReserve", reserveInput, &createReserveOutput)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CreateReserve Response:", createReserveOutput)
}
