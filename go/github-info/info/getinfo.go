package info

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/arthurschuelter/go-git/config"
)

func GetInfo(username string, config *config.Config) error {
	return GetInfoWithOptions(username, config, false)
}

// GetInfoWithOptions generates a report, optionally skipping the KLOC calculation.
func GetInfoWithOptions(username string, config *config.Config, skipKLOC bool) error {
	return GetInfoWithRefresh(username, config, skipKLOC, false)
}

// GetInfoWithRefresh optionally skips KLOC and bypasses the cached report.
func GetInfoWithRefresh(username string, config *config.Config, skipKLOC, forceRefresh bool) error {
	cacheName := "gogit.json"
	if skipKLOC {
		cacheName = "gogit-no-kloc.json"
	}
	cachePath := filepath.Join(".cache", cacheName)
	now := time.Now()
	if report, ok := loadCache(cachePath, username, now); !forceRefresh && ok && (skipKLOC || report.Lines != nil) {
		fmt.Printf("Cached report (fetched %s)\n", report.FetchedAt.Local().Format(time.RFC3339))
		printReport(report)
		return nil
	}

	userURL := fmt.Sprintf("https://api.github.com/users/%s", url.PathEscape(username))
	client := &http.Client{Timeout: 30 * time.Second}
	user, err := GetUserData(userURL, client, config.Token)
	if err != nil {
		return err
	}
	repos, err := GetRepoData(userURL+"/repos", client, config.Token)
	if err != nil {
		return err
	}
	ranking, err := GetLanguageData(repos, client, config.Token)
	if err != nil {
		return err
	}
	var kloc map[string]int64
	if !skipKLOC {
		kloc, err = GetKLOC(repos)
		if err != nil {
			return err
		}
	}

	report := &reportCache{Version: cacheVersion, Username: username, FetchedAt: now.UTC(), User: user, Repos: repos, Ranking: ranking, Lines: kloc}
	if err := saveCache(cachePath, report); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: could not save cache: %v\n", err)
	}
	printReport(report)
	return nil
}

func printReport(report *reportCache) {
	report.User.PrintData()
	PrintRepoData(report.Repos)
	printLanguages(os.Stdout, report.Ranking, report.Lines)
}
