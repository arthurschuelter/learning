package info

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/arthurschuelter/go-git/models"
	"github.com/arthurschuelter/go-git/utils"
)

func GetRepoData(rawURL string, client *http.Client, token string) ([]models.Repo, error) {
	first, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	query := first.Query()
	query.Set("per_page", "100")
	first.RawQuery = query.Encode()
	next := first.String()
	var repos []models.Repo
	seen := make(map[string]bool)
	for next != "" {
		if seen[next] {
			return nil, fmt.Errorf("repository pagination loop")
		}
		seen[next] = true
		req, err := utils.ConfigRequest(next, token)
		if err != nil {
			return nil, err
		}
		body, header, err := utils.MakeRequestAndRead(client, req)
		if err != nil {
			return nil, fmt.Errorf("get repositories: %w", err)
		}
		page, err := ReadRepos(body)
		if err != nil {
			return nil, fmt.Errorf("decode repositories: %w", err)
		}
		repos = append(repos, page...)
		next = ""
		for _, link := range strings.Split(header.Get("Link"), ",") {
			parts := strings.Split(link, ";")
			if len(parts) < 2 {
				continue
			}
			for _, param := range parts[1:] {
				if strings.TrimSpace(param) != `rel="next"` {
					continue
				}
				target := strings.TrimSpace(parts[0])
				if !strings.HasPrefix(target, "<") || !strings.HasSuffix(target, ">") {
					return nil, fmt.Errorf("invalid pagination link")
				}
				u, err := req.URL.Parse(target[1 : len(target)-1])
				if err != nil || u.Scheme != first.Scheme || u.Host != first.Host {
					return nil, fmt.Errorf("invalid pagination URL")
				}
				next = u.String()
			}
		}
	}
	return repos, nil
}

func ReadRepos(data []byte) ([]models.Repo, error) {
	var repos []models.Repo
	err := json.Unmarshal(data, &repos)
	return repos, err
}

func PrintRepoData(repos []models.Repo) {
	count := 0
	for _, r := range repos {
		if !r.Fork && !r.Private {
			count++
		}
	}
	// fmt.Printf("%d repositories\n", count)
	stars := 0
	for _, r := range repos {
		if r.Fork || r.Private {
			continue
		}
		// fmt.Printf("%s\n", r.FullName)
		stars += r.Stars
	}
	fmt.Printf("Stars: %d ⭐\n", stars)
}
