package models

type ProductDTO struct {
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Description string  `json:"description"`
	CategoryID  string  `json:"category_id" binding:"required"`
}
