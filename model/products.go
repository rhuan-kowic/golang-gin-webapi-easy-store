package model

type Product struct {
	ID            uint    `json:"id_product"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	StockQuantity int     `json:"stock_quantity"`
	MinimumStock  int     `json:"minimum_stock"`
}
