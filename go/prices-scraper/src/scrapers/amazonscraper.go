package scrapers

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"items-scraper/src/models"
	"items-scraper/src/repository"
	"items-scraper/src/utils"

	"github.com/gocolly/colly"
)

var amazonDomains = []string{
	"www.amazon.com.br",
}

func ScrapeAmazon(itemList []models.Item, repository *repository.ProductRepository) {
	baseURL := "https://www.amazon.com.br/s?k="
	retailer := "Amazon"

	links := []string{}
	for _, item := range itemList {
		url := strings.ReplaceAll(strings.ToLower(item.Title), " ", "+")
		links = append(links, baseURL+url)
	}

	fmt.Printf("Scanning %s for:\n", retailer)
	for _, item := range itemList {
		fmt.Printf("- %s\n", item.Title)
	}

	c := colly.NewCollector(
		colly.AllowedDomains(amazonDomains...),
	)

	c.OnHTML("div[role='listitem']", func(e *colly.HTMLElement) {
		title := extractTitleAmazon(e)
		price := extractPriceAmazon(e)
		l := extractLinkAmazon(e)
		id := extractProductIDAmazon(l)

		item := models.Item{
			Title:    title,
			Currency: "R$",
			Price:    price,
			ID:       id,
			Link:     l,
		}

		product := models.Product{
			IDProduct:   id,
			ProductName: title,
			URL:         l,
			Retailer:    retailer,
		}

		if utils.ValidateItem(item, itemList) {
			err := SaveProductPrice(item, product, repository)
			utils.LogErr(err)
		}
	})

	for i, link := range links {
		fmt.Printf("Scanning %s\n%s\n", itemList[i].Title, link)
		err := c.Visit(link)
		utils.LogErr(err)
	}
}

func extractTitleAmazon(e *colly.HTMLElement) string {
	return e.ChildText("div[data-cy='title-recipe'] h2 span")
}

func extractPriceAmazon(e *colly.HTMLElement) float64 {
	var price float64

	var priceWhole = e.ChildText("div[data-cy='price-recipe'] span.a-price-whole")
	var priceFraction = e.ChildText("div[data-cy='price-recipe'] span.a-price-fraction")
	priceStr := priceWhole + priceFraction
	priceStr = strings.ReplaceAll(priceStr, ".", "")
	priceStr = strings.ReplaceAll(priceStr, ",", ".")

	var err error
	price, err = strconv.ParseFloat(priceStr, 64)
	if err != nil {
		price = 0
	}

	return price
}

func extractLinkAmazon(e *colly.HTMLElement) string {
	baseUrl := "https://www.amazon.com.br"

	link := e.ChildAttr("span[data-component-type='s-product-image'] a.a-link-normal", "href")
	link = utils.UrlDecode(link)

	reURL := regexp.MustCompile(`^([^?]+?)(?:/ref=|$|\?)`)
	if match := reURL.FindStringSubmatch(link); len(match) > 1 {
		cleanURL := match[1]
		link = cleanURL
	}

	return baseUrl + link
}

func extractProductIDAmazon(link string) string {
	reID := regexp.MustCompile(`/dp/([A-Z0-9]+)`)
	productID := ""
	if match := reID.FindStringSubmatch(link); len(match) > 1 {
		productID = match[1]
	}
	return productID
}
