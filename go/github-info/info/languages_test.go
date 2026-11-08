package info

import (
	"bytes"
	"strings"
	"testing"

	"github.com/arthurschuelter/go-git/models"
)

func TestCombinedLanguageReport(t *testing.T) {
	ranking := []models.Ranking{
		{Language: "Go", Total: 2048, Ratio: .2},
		{Language: "JavaScript", Total: 1024, Ratio: .1},
		{Language: "JSON", Total: 4096, Ratio: .4},
		{Language: "Shell", Total: 512, Ratio: .05},
		{Language: "Rust", Total: 256, Ratio: .025},
	}
	lines := map[string]int64{"Go": 1200, "JSX": 100, "JavaScript": 200, "Bourne Shell": 50, "Python": 300, "Text": 9999, "C/C++ Header": 500}
	var output bytes.Buffer
	printLanguages(&output, ranking, lines)
	text := output.String()
	for _, want := range []string{
		"Go: 2.00 KiB (20.00%) | 1.200 KLOC",
		"JavaScript: 1.00 KiB (10.00%) | 0.300 KLOC",
		"Shell: 0.50 KiB (5.00%) | 0.050 KLOC",
		"Rust: 0.25 KiB (2.50%) | n/a",
		"Python: n/a | 0.300 KLOC",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("missing %q in %s", want, text)
		}
	}
	for _, excluded := range []string{"JSON:", "Text:", "C/C++ Header:"} {
		if strings.Contains(text, excluded) {
			t.Errorf("unexpected %q in %s", excluded, text)
		}
	}
}
