package info

import (
	"errors"
	"strings"
	"testing"

	"github.com/arthurschuelter/go-git/models"
)

func TestCountKLOCAggregatesAndFilters(t *testing.T) {
	repos := []models.Repo{{FullName: "user/one"}, {FullName: "user/fork", Fork: true}, {FullName: "user/private", Private: true}, {FullName: "user/two"}}
	clones, scans := 0, 0
	got, err := countKLOC(repos, func(name string, args ...string) ([]byte, error) {
		if name == "git" {
			clones++
			return nil, nil
		}
		scans++
		if scans == 1 {
			return []byte(`{"header":{},"Go":{"code":1200},"SUM":{"code":1200}}`), nil
		}
		return []byte(`{"header":{},"Go":{"code":300},"Python":{"code":25},"SUM":{"code":325}}`), nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if clones != 2 || scans != 2 || got["Go"] != 1500 || got["Python"] != 25 || len(got) != 2 {
		t.Fatalf("unexpected counts: clones=%d scans=%d totals=%v", clones, scans, got)
	}
}

func TestCountKLOCReportsCloneError(t *testing.T) {
	_, err := countKLOC([]models.Repo{{FullName: "user/repo"}}, func(name string, args ...string) ([]byte, error) { return nil, errors.New("failed") })
	if err == nil || !strings.Contains(err.Error(), "user/repo") {
		t.Fatalf("unexpected error: %v", err)
	}
}
