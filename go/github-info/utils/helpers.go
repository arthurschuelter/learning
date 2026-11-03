package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"

	"github.com/arthurschuelter/go-git/models"
)

func SortList(list []models.Ranking) []models.Ranking {
	sort.Slice(list, func(i, j int) bool {
		return list[i].Total > list[j].Total
	})
	return list
}

func PrintRanking(item models.Ranking, i int) {
	fmt.Printf("  (%2d) %s: %d (%.2f%%)\n", i+1, item.Language, item.Total, item.Ratio*100)
}

func ConfigRequest(url string, token string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	return req
}

func MakeRequestAndRead(client *http.Client, req *http.Request) ([]byte, error) {
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func ReadJson[T models.Repo | models.Ranking](data []byte) ([]T, error) {
	var repos []T
	err := json.Unmarshal(data, &repos)
	return repos, err
}
