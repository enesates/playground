package main

import (
	"log"
	"time"
)

func main() {
	SetupDB()
	defer CloseConnection()

	product := Product{
		ID:        16,
		Name:      "Harry Potter",
		Category:  "Book",
		Price:     22.50,
		Stock:     20,
		Tags:      []string{"Children", "Fiction"},
		CreatedAt: time.Now(),
	}

	productId := AddProduct(product)

	log.Println("Added Product")
	GetProductByID(productId)

	GetProductByCategory("Book")
}
