package scrapers

import (
	"fmt"
	"items-scraper/src/models"
	"items-scraper/src/repository"
	"sync"
)

func ScapeAll(items []models.Item, repository *repository.ProductRepository) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		ScrapeMercadoLivre(items, repository)
	}()

	go func() {
		defer wg.Done()
		ScrapeAmazon(items, repository)
	}()

	wg.Wait()
	fmt.Println("✓ All scrapers completed!")

}

func SaveProductPrice(item models.Item, product models.Product, repo *repository.ProductRepository) error {
	// Find product
	productID, err := repo.FindByIDProduct(product.IDProduct)
	if err != nil {
		return fmt.Errorf("error finding product: %w", err)
	}

	// Product doesn't exist, create it
	if productID == -1 {
		productID, err = repo.InsertProduct(product)
		if err != nil {
			return fmt.Errorf("error inserting product: %w", err)
		}
	}

	// Insert price history
	priceHistory := models.PriceHistory{
		ProductID: productID,
		Price:     item.Price,
		Currency:  item.Currency,
	}

	_, err = repo.InsertPriceHistory(priceHistory)
	if err != nil {
		return fmt.Errorf("error inserting price history: %w", err)
	}

	return nil
}
