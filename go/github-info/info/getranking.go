package info

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arthurschuelter/go-git/models"
	"github.com/arthurschuelter/go-git/utils"
)

func GetLanguageData(repos []models.Repo, client *http.Client, token string) ([]models.Ranking, error) {
	fmt.Println("Language Ranking:")
	LanguagesMap := make(map[string]int)
	for _, r := range repos {
		if r.Fork || r.Private {
			continue
		}
		// fmt.Println(r.FullName)
		link := r.LanguagesURL

		req, err := utils.ConfigRequest(link, token)
		if err != nil {
			return nil, err
		}
		body, _, err := utils.MakeRequestAndRead(client, req)

		if err != nil {
			return nil, fmt.Errorf("get languages for %s: %w", r.FullName, err)
		}

		languages, err := ReadLanguage(body)

		if err != nil {
			return nil, fmt.Errorf("get languages for %s: %w", r.FullName, err)
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
	return ranking, nil
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
