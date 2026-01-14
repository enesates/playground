package dtos

type StockDTO struct {
	IncerementBy int    `json:"increment_by" binding:"required"`
	Reason       string `json:"reason"`
}
