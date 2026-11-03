package info

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arthurschuelter/go-git/models"
	"github.com/arthurschuelter/go-git/utils"
)

func GetLanguageData(repos []models.Repo, client *http.Client, token string) []models.Ranking {
	fmt.Println("Language Ranking:")
	LanguagesMap := make(map[string]int)
	for _, r := range repos {
		// fmt.Println(r.FullName)
		link := r.LanguagesURL

		req := utils.ConfigRequest(link, token)
		body, err := utils.MakeRequestAndRead(client, req)

		if err != nil {
			panic(err)
		}

		languages, err := ReadLanguage(body)

		if err != nil {
			panic(err)
		}

		for key, value := range languages {
			_, ok := LanguagesMap[key]
			if !ok {
				LanguagesMap[key] = value
			} else {
				LanguagesMap[key] += value
			}
		}

	}

	total := GetTotal(LanguagesMap)
	ranking := []models.Ranking{}

	for key, value := range LanguagesMap {
		// fmt.Println(key, value)
		lang := models.Ranking{
			Language: key,
			Total:    value,
			Ratio:    float32(value) / float32(total),
		}
		ranking = append(ranking, lang)
	}

	ranking = utils.SortList(ranking)
	return ranking
}

func ReadLanguage(data []byte) (map[string]int, error) {
	var languagesMap map[string]int
	err := json.Unmarshal(data, &languagesMap)
	return languagesMap, err
}

func GetTotal(LanguagesMap map[string]int) int {
	sum := 0

	for _, value := range LanguagesMap {
		sum += value
	}

	return sum
}
