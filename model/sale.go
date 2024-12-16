package model

import "time"

type Sale struct {
	ID        uint      `json:"id_sale"`
	ProductID uint      `json:"product_id"`
	DateSale  time.Time `json:"date_sale"`
	Total     float64   `json:"total"`
	Quantity  int       `json:"quantity"`
}
