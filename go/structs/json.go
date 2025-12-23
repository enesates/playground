package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
}

type OrderItem struct {
	ProductID   string  `json:"productId"`
	ProductName string  `json:"productName"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Order struct {
	OrderID         string      `json:"orderId"`
	CustomerID      string      `json:"customerId"`
	OrderDate       time.Time   `json:"orderDate"`
	Status          string      `json:"status"`
	Items           []OrderItem `json:"items"`
	SubTotal        float64     `json:"subtotal"`
	Tax             float64     `json:"tax"`
	ShippingCost    float64     `json:"shippingCost"`
	TotalAmount     float64     `json:"totalAmount"`
	ShippingAddress Address     `json:"shippingAddress"`
	PaymentMethod   string      `json:"paymentMethod"`
}

type Report struct {
	CustomerId string
	OrderId    string
	Address    Address
}

func main() {
	jsonData, err := os.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	order := Order{}
	if err := json.Unmarshal(jsonData, &order); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Order: %+v\n", order)

	report := Report{}
	report.CustomerId = order.CustomerID
	report.OrderId = order.OrderID
	report.Address = order.ShippingAddress

	if reportJson, err := json.MarshalIndent(report, "", "\t"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Report: %+v\n", string(reportJson))
	}
}
