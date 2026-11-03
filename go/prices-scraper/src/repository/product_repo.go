package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"items-scraper/src/models"
	"time"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) FindByIDProduct(idProduct string) (int, error) {
	query := "SELECT id FROM products WHERE id_product = $1"
	var id int
	err := r.db.QueryRow(query, idProduct).Scan(&id)
	if err == sql.ErrNoRows {
		return -1, nil
	}
	return id, err
}

func (r *ProductRepository) InsertProduct(product models.Product) (int, error) {
	query := `
        INSERT INTO products (id_product, product_name, url, retailer)
        VALUES ($1, $2, $3, $4)
        RETURNING id`

	var id int
	err := r.db.QueryRow(query,
		product.IDProduct,
		product.ProductName,
		product.URL,
		product.Retailer,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to insert product: %w", err)
	}

	fmt.Printf("✓ Inserted product ID: %d\n", id)
	return id, nil
}

func (r *ProductRepository) InsertPriceHistory(ph models.PriceHistory) (int, error) {
	if !r.canInsertPriceHistory(ph) {
		return 0, errors.New("duplicate: price already recorded today")
	}

	query := `
        INSERT INTO price_history (product_id, price, currency, price_date)
        VALUES ($1, $2, $3, $4)
        RETURNING id`

	var id int
	err := r.db.QueryRow(query,
		ph.ProductID,
		ph.Price,
		ph.Currency,
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to insert price history: %w", err)
	}

	fmt.Printf("✓ Inserted price history ID: %d\n", id)
	return id, nil
}

func (r *ProductRepository) CanInsertProduct(p models.Product, id int) bool {
	return id == -1 && p.IDProduct != ""
}

func (r *ProductRepository) canInsertPriceHistory(ph models.PriceHistory) bool {
	query := `
        SELECT id FROM price_history 
        WHERE product_id = $1 AND price_date = CURRENT_DATE`

	var id int
	err := r.db.QueryRow(query, ph.ProductID).Scan(&id)
	return err == sql.ErrNoRows
}
