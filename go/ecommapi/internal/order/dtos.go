package order

import (
	"ecommapi/internal/cart"
)

type OrderDTO struct {
	Items           []cart.CartItemDTO `json:"items" binding:"required"`
	ShippingStreet  string             `json:"shipping_street"`
	ShippingCity    string             `json:"shipping_city"`
	ShippingZip     string             `json:"shipping_zip"`
	ShippingCountry string             `json:"shipping_country"`
}
