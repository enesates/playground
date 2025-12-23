package main

import "fmt"

type Price struct {
	Amount   float64
	Currency string
}

type Product struct {
	Name string
	Price
}

func main() {
	product := Product{
		Name:  "Jacket",
		Price: Price{Amount: 73.50, Currency: "EUR"},
	}

	fmt.Printf("Product: %s, Price: %.2f %s", product.Name, product.Price.Amount, product.Price.Currency)
}
