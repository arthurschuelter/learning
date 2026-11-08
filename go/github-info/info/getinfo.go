package info

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/arthurschuelter/go-git/config"
)

func GetInfo(username string, config *config.Config) error {
	userURL := fmt.Sprintf("https://api.github.com/users/%s", url.PathEscape(username))
	client := &http.Client{Timeout: 30 * time.Second}

	user, err := GetUserData(userURL, client, config.Token)
	if err != nil {
		return err
	}
	user.PrintData()

	repos, err := GetRepoData(userURL+"/repos", client, config.Token)
	if err != nil {
		return err
	}
	PrintRepoData(repos)

	ranking, err := GetLanguageData(repos, client, config.Token)
	if err != nil {
		return err
	}
	for i, r := range ranking {
		r.PrintRanking(i)
	}
	return nil
}
