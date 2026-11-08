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
	sort.Slice(list, func(i, j int) bool { return list[i].Total > list[j].Total })
	return list
}

func ConfigRequest(url string, token string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github+json")
	return req, nil
}

func MakeRequestAndRead(client *http.Client, req *http.Request) ([]byte, http.Header, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var apiError struct {
			Message string `json:"message"`
		}
		_ = json.Unmarshal(body, &apiError)
		if apiError.Message != "" {
			return nil, nil, fmt.Errorf("GitHub API: %s: %s", resp.Status, apiError.Message)
		}
		return nil, nil, fmt.Errorf("GitHub API: %s", resp.Status)
	}
	return body, resp.Header, nil
}
