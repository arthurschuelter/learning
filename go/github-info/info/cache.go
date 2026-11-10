package info

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/arthurschuelter/go-git/models"
)

const (
	cacheVersion  = 1
	cacheLifetime = 24 * time.Hour
)

// reportCache contains only public report data; it never stores the API token.
type reportCache struct {
	Version   int              `json:"version"`
	Username  string           `json:"username"`
	FetchedAt time.Time        `json:"fetched_at"`
	User      models.User      `json:"user"`
	Repos     []models.Repo    `json:"repositories"`
	Ranking   []models.Ranking `json:"language_ranking"`
	Lines     map[string]int64 `json:"code_lines"`
}

func loadCache(path, username string, now time.Time) (*reportCache, bool) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, false
	}
	var report reportCache
	if json.Unmarshal(data, &report) != nil || report.Version != cacheVersion ||
		!strings.EqualFold(report.Username, username) || report.FetchedAt.IsZero() ||
		report.FetchedAt.After(now) || now.Sub(report.FetchedAt) >= cacheLifetime ||
		report.Ranking == nil {
		return nil, false
	}
	return &report, true
}

func saveCache(path string, report *reportCache) error {
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return err
	}
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	file, err := os.CreateTemp(dir, ".gogit-cache-*")
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())
	if _, err := file.Write(append(data, '\n')); err != nil {
		file.Close()
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	if err := os.Rename(file.Name(), path); err != nil {
		return fmt.Errorf("replace cache: %w", err)
	}
	return nil
}
