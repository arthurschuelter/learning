package info

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/arthurschuelter/go-git/models"
)

var repoNamePattern = regexp.MustCompile(`^[A-Za-z0-9_.-]+/[A-Za-z0-9_.-]+$`)

type languageLines struct {
	Code int64 `json:"code"`
}

type commandRunner func(name string, args ...string) ([]byte, error)

// GetKLOC counts code lines in shallow clones of public, non-fork repositories.
func GetKLOC(repos []models.Repo) (map[string]int64, error) {
	if _, err := exec.LookPath("git"); err != nil {
		return nil, fmt.Errorf("KLOC requires git: %w", err)
	}
	if _, err := exec.LookPath("cloc"); err != nil {
		return nil, fmt.Errorf("KLOC requires cloc; install it and ensure it is on PATH: %w", err)
	}
	return countKLOC(repos, func(name string, args ...string) ([]byte, error) {
		cmd := exec.Command(name, args...)
		cmd.Env = append(os.Environ(), "GIT_TERMINAL_PROMPT=0")
		return cmd.Output()
	})
}

func countKLOC(repos []models.Repo, run commandRunner) (map[string]int64, error) {
	dir, err := os.MkdirTemp("", "gogit-kloc-*")
	if err != nil {
		return nil, fmt.Errorf("create temporary directory: %w", err)
	}
	defer os.RemoveAll(dir)

	totals := make(map[string]int64)
	for i, repo := range repos {
		if repo.Fork || repo.Private {
			continue
		}
		if !repoNamePattern.MatchString(repo.FullName) {
			return nil, fmt.Errorf("invalid repository name %q", repo.FullName)
		}
		path := filepath.Join(dir, fmt.Sprintf("repo-%d", i))
		if _, err := run("git", "clone", "--quiet", "--depth=1", "--", "https://github.com/"+repo.FullName+".git", path); err != nil {
			return nil, fmt.Errorf("clone %s: %w", repo.FullName, err)
		}
		output, err := run("cloc", "--json", "--quiet", path)
		if err != nil {
			return nil, fmt.Errorf("count lines in %s: %w", repo.FullName, err)
		}
		// cloc may emit no JSON when there are no recognized files.
		if len(output) == 0 {
			continue
		}
		var entries map[string]json.RawMessage
		if err := json.Unmarshal(output, &entries); err != nil {
			return nil, fmt.Errorf("decode cloc output for %s: %w", repo.FullName, err)
		}
		for language, data := range entries {
			if language == "header" || language == "SUM" {
				continue
			}
			var lines languageLines
			if err := json.Unmarshal(data, &lines); err != nil {
				return nil, fmt.Errorf("decode cloc language %s: %w", language, err)
			}
			totals[language] += lines.Code
		}
	}
	return totals, nil
}
