package info

import (
	"fmt"
	"net/http"

	"github.com/arthurschuelter/go-git/config"
)

func GetInfo(username string, config *config.Config) {
	reposURL := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	userURL := fmt.Sprintf("https://api.github.com/users/%s", username)
	client := http.Client{}

	user := GetUserData(userURL, &client, config.Token)
	user.PrintData()

	repos := GetRepoData(reposURL, &client, config.Token)
	PrintRepoData(repos)

	ranking := GetLanguageData(repos, &client, config.Token)
	for i, r := range ranking {
		r.PrintRanking(i)
	}
}
