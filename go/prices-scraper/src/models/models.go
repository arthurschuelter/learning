package models

import "time"

type Item struct {
	ID       string
	Currency string
	Link     string
	Price    float64
	MaxPrice float64
	MinPrice float64
	Title    string
}

type Product struct {
	ID          int       `db:"id" json:"id"`
	IDProduct   string    `db:"id_product" json:"id_product"`
	ProductName string    `db:"product_name" json:"product_name"`
	URL         string    `db:"url" json:"url"`
	Retailer    string    `db:"retailer" json:"retailer"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

type PriceHistory struct {
	ID        int       `db:"id" json:"id"`
	ProductID int       `db:"product_id" json:"product_id"`
	Price     float64   `db:"price" json:"price"`
	Currency  string    `db:"currency" json:"currency"`
	PriceDate time.Time `db:"price_date" json:"price_date"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
