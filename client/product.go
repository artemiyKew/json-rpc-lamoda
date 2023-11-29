package main

import (
	"math/rand"
	"time"
)

// ProductCreateInput структура для создания продукта
type ProductCreateInput struct {
	Name        string
	Size        string
	UniqueCode  string
	Quantity    int
	WarehouseID int
}

var (
	names        = []string{"T-shirt", "sweatshirt", "Jacket"}
	sizes        = []string{"XS", "S", "M", "L", "XL"}
	warehouseIDs = []int{1, 2}
)

// Генератор случайной строки заданной длины
func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// Генератор случайного ProductCreateInput
func generateRandomProduct() ProductCreateInput {
	return ProductCreateInput{
		Name:        names[rand.Intn(len(names))],
		Size:        sizes[rand.Intn(len(sizes))],
		UniqueCode:  generateRandomString(10),
		Quantity:    rand.Intn(10) + 1,
		WarehouseID: warehouseIDs[rand.Intn(len(warehouseIDs))],
	}
}
