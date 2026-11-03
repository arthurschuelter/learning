package utils

import (
	"fmt"
	"items-scraper/src/models"
	"log"
	"net/url"
	"sort"
	"strings"
)

func SortList(list []models.Item) []models.Item {
	sort.Slice(list, func(i, j int) bool {
		return list[i].Price < list[j].Price
	})
	return list
}

func UrlDecode(s string) string {
	decoded, err := url.QueryUnescape(s)
	if err != nil {
		log.Fatal(err)
	}
	return decoded
}

func ValidateItem(item models.Item, compareList []models.Item) bool {
	title := strings.ToUpper(item.Title)
	title = strings.ReplaceAll(title, "™", "")

	for _, compare := range compareList {
		if strings.Contains(title, strings.ToUpper(compare.Title)) && item.Price >= compare.MinPrice && item.Price <= compare.MaxPrice && item.ID != "" {
			return true
		}
	}

	return false
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func LogErr(err error) {
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
	}
}
