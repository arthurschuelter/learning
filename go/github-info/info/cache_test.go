package info

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/arthurschuelter/go-git/config"
	"github.com/arthurschuelter/go-git/models"
)

func TestCacheFreshnessAndIdentity(t *testing.T) {
	now := time.Date(2026, 1, 2, 12, 0, 0, 0, time.UTC)
	path := filepath.Join(t.TempDir(), "cache", "report.json")
	report := &reportCache{Version: cacheVersion, Username: "Example", FetchedAt: now, Ranking: []models.Ranking{}, Lines: map[string]int64{"Go": 1200}}
	if err := saveCache(path, report); err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name, username string
		now            time.Time
		want           bool
	}{
		{"fresh", "example", now.Add(23*time.Hour + 59*time.Minute), true},
		{"expired", "Example", now.Add(24 * time.Hour), false},
		{"other user", "someone", now, false},
		{"future timestamp", "Example", now.Add(-time.Second), false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := loadCache(path, tc.username, tc.now)
			if ok != tc.want {
				t.Fatalf("cache hit = %v, want %v", ok, tc.want)
			}
			if ok && got.Lines["Go"] != 1200 {
				t.Fatalf("unexpected lines: %v", got.Lines)
			}
		})
	}
	if err := os.WriteFile(path, []byte(`{"version":1`), 0600); err != nil {
		t.Fatal(err)
	}
	if _, ok := loadCache(path, "Example", now); ok {
		t.Fatal("corrupt cache should miss")
	}
}

func TestGetInfoUsesFreshCacheWithoutDependencies(t *testing.T) {
	dir := t.TempDir()
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(oldDir); err != nil {
			t.Error(err)
		}
	})
	t.Setenv("PATH", "")
	report := &reportCache{Version: cacheVersion, Username: "example", FetchedAt: time.Now(), User: models.User{Login: "example"}, Ranking: []models.Ranking{}, Lines: map[string]int64{}}
	if err := saveCache(filepath.Join(".cache", "gogit.json"), report); err != nil {
		t.Fatal(err)
	}
	if err := GetInfo("example", &config.Config{Token: "invalid"}); err != nil {
		t.Fatalf("fresh cache should avoid API and cloc: %v", err)
	}
}

func TestGetInfoUsesCacheWithoutKLOC(t *testing.T) {
	dir := t.TempDir()
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(oldDir); err != nil {
			t.Error(err)
		}
	})
	t.Setenv("PATH", "")
	report := &reportCache{Version: cacheVersion, Username: "example", FetchedAt: time.Now(), Ranking: []models.Ranking{}}
	if err := saveCache(filepath.Join(".cache", "gogit-no-kloc.json"), report); err != nil {
		t.Fatal(err)
	}
	if err := GetInfoWithOptions("example", &config.Config{Token: "invalid"}, true); err != nil {
		t.Fatalf("fast cache should avoid API and cloc: %v", err)
	}
	if _, err := os.Stat(filepath.Join(".cache", "gogit.json")); !os.IsNotExist(err) {
		t.Fatalf("fast report should not create a full report cache: %v", err)
	}
}

type failingTransport struct{}

func (failingTransport) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, errors.New("refresh attempted")
}

func TestForceRefreshBypassesCache(t *testing.T) {
	dir := t.TempDir()
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(oldDir); err != nil {
			t.Error(err)
		}
	})
	oldTransport := http.DefaultTransport
	http.DefaultTransport = failingTransport{}
	t.Cleanup(func() { http.DefaultTransport = oldTransport })
	for _, tc := range []struct {
		name string
		skip bool
	}{{"gogit.json", false}, {"gogit-no-kloc.json", true}} {
		t.Run(tc.name, func(t *testing.T) {
			path := filepath.Join(".cache", tc.name)
			report := &reportCache{Version: cacheVersion, Username: "example", FetchedAt: time.Now(), Ranking: []models.Ranking{}, Lines: map[string]int64{}}
			if err := saveCache(path, report); err != nil {
				t.Fatal(err)
			}
			before, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			if err := GetInfoWithRefresh("example", &config.Config{}, tc.skip, true); err == nil {
				t.Fatal("forced refresh should attempt the API")
			}
			after, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			if string(before) != string(after) {
				t.Fatal("failed refresh replaced the cache")
			}
		})
	}
}
