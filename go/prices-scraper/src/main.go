package main

import (
	"items-scraper/src/config"
	"items-scraper/src/models"
	"items-scraper/src/repository"
	"items-scraper/src/scrapers"
	"items-scraper/src/utils"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig()
	db, err := cfg.Connect()
	utils.CheckErr(err)
	productRepo := repository.NewProductRepository(db)
	items := []models.Item{
		{ID: "1", Title: "Switch 2", MinPrice: 2200, MaxPrice: 9999},
		{ID: "2", Title: "Steam Deck", MinPrice: 2200, MaxPrice: 9999},
	}

	scrapers.ScapeAll(items, productRepo)

	defer func() {
		err := db.Close()
		utils.LogErr(err)
	}()
}
