package scrapers

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"items-scraper/src/models"
	"items-scraper/src/repository"
	"items-scraper/src/utils"

	"github.com/gocolly/colly"
)

var domains = []string{
	"www.mercadolivre.com.br",
	"lista.mercadolivre.com.br",
}

func ScrapeMercadoLivre(itemList []models.Item, repository *repository.ProductRepository) {
	baseURL := "https://lista.mercadolivre.com.br/"
	retailer := "Mercado Livre"

	links := []string{}
	for _, item := range itemList {
		url := strings.ReplaceAll(strings.ToLower(item.Title), " ", "-")
		links = append(links, baseURL+url)
	}

	fmt.Printf("Scanning %s for:\n", retailer)
	for _, item := range itemList {
		fmt.Printf("- %s\n", item.Title)
	}

	c := colly.NewCollector(
		colly.AllowedDomains(domains...),
	)

	c.OnHTML("div.ui-search-result__wrapper", func(e *colly.HTMLElement) {
		title := extractTitleMeli(e)
		price := extractPriceMeli(e)
		l := extractLinkMeli(e)
		id, err := extractProductIDMeli(l)
		if err != nil {
			fmt.Printf("[ERROR] Error extracting id: %v\n", err)
			fmt.Printf("[ERROR] %s: %.2f\n", title, price)
			return
		}

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
		fmt.Printf("Scanning %s\n", itemList[i].Title)
		err := c.Visit(link)
		utils.LogErr(err)
	}
}

func extractTitleMeli(e *colly.HTMLElement) string {
	return e.ChildText("h3.poly-component__title-wrapper")
}

func extractPriceMeli(e *colly.HTMLElement) float64 {
	var price float64
	e.ForEach("div.poly-price__current span.andes-money-amount__fraction", func(i int, el *colly.HTMLElement) {
		if i == 0 {
			priceStr := strings.ReplaceAll(el.Text, ".", "")
			var err error
			price, err = strconv.ParseFloat(priceStr, 64)
			if err != nil {
				price = 0
			}

		}
	})
	return price
}

func extractLinkMeli(e *colly.HTMLElement) string {
	return e.ChildAttr("h3.poly-component__title-wrapper a.poly-component__title", "href")
}

func extractProductIDMeli(rawURL string) (id string, err error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	query := u.Query()
	if wid := query.Get("wid"); wid != "" {
		return wid, nil
	}

	upRegex := regexp.MustCompile(`/up/([A-Z0-9]+)`)
	if matches := upRegex.FindStringSubmatch(u.Path); len(matches) > 1 {
		return matches[1], nil
	}

	pRegex := regexp.MustCompile(`/p/([A-Z0-9]+)`)
	if matches := pRegex.FindStringSubmatch(u.Path); len(matches) > 1 {
		return matches[1], nil
	}

	mlbRegex := regexp.MustCompile(`/MLB-([0-9]+)-`)
	if matches := mlbRegex.FindStringSubmatch(u.Path); len(matches) > 1 {
		return "MLB" + matches[1], nil
	}

	return "", fmt.Errorf("no MLB ID found in URL")
}
