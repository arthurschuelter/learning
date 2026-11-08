package info

import (
	"fmt"
	"io"
	"sort"

	"github.com/arthurschuelter/go-git/models"
)

// The allowlist keeps data, markup, and project formats out of the report.
// Extend it when support for another programming language is needed.
var programmingLanguages = map[string]bool{
	"Ada": true, "Assembly": true, "Awk": true, "C": true, "C++": true,
	"C#": true, "Clojure": true, "COBOL": true, "CoffeeScript": true,
	"Crystal": true, "CUDA": true, "D": true, "Dart": true, "Elixir": true,
	"Elm": true, "Erlang": true, "F#": true, "Fortran": true, "GDScript": true,
	"GLSL": true, "Go": true, "Groovy": true, "Haskell": true, "Java": true,
	"JavaScript": true, "Julia": true, "Kotlin": true, "Lisp": true, "Lua": true,
	"MATLAB": true, "Objective-C": true, "OCaml": true, "Pascal": true,
	"Perl": true, "PHP": true, "PowerShell": true, "Prolog": true, "Python": true,
	"R": true, "Ruby": true, "Rust": true, "Scala": true, "Shell": true,
	"Batchfile": true, "SQL": true, "Swift": true, "TypeScript": true,
	"Visual Basic": true, "Zig": true, "Godot Shaders": true, "Logos": true,
	"CMake": true, "Makefile": true,
}

var languageAliases = map[string]string{
	"Bourne Shell": "Shell", "Bourne Again Shell": "Shell", "DOS Batch": "Batchfile",
	"make": "Makefile", "JSX": "JavaScript", "C/C++ Header": "", // Header language is ambiguous.
}

func canonicalLanguage(language string) string {
	if alias, ok := languageAliases[language]; ok {
		language = alias
	}
	if !programmingLanguages[language] {
		return ""
	}
	return language
}

type languageReport struct {
	name     string
	bytes    int64
	ratio    float64
	lines    int64
	hasBytes bool
	hasLines bool
}

func mergeLanguages(ranking []models.Ranking, lines map[string]int64) []languageReport {
	byName := make(map[string]*languageReport)
	get := func(name string) *languageReport {
		if byName[name] == nil {
			byName[name] = &languageReport{name: name}
		}
		return byName[name]
	}
	for _, item := range ranking {
		name := canonicalLanguage(item.Language)
		if name == "" {
			continue
		}
		row := get(name)
		row.bytes += int64(item.Total)
		row.ratio += float64(item.Ratio)
		row.hasBytes = true
	}
	for name, count := range lines {
		name = canonicalLanguage(name)
		if name == "" {
			continue
		}
		row := get(name)
		row.lines += count
		row.hasLines = true
	}
	rows := make([]languageReport, 0, len(byName))
	for _, row := range byName {
		rows = append(rows, *row)
	}
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].bytes != rows[j].bytes {
			return rows[i].bytes > rows[j].bytes
		}
		if rows[i].lines != rows[j].lines {
			return rows[i].lines > rows[j].lines
		}
		return rows[i].name < rows[j].name
	})
	return rows
}

func printLanguages(w io.Writer, ranking []models.Ranking, lines map[string]int64) {
	fmt.Fprintln(w, "Programming languages (GitHub byte share; cloc code lines):")
	rows := mergeLanguages(ranking, lines)
	if len(rows) == 0 {
		fmt.Fprintln(w, "  No recognized programming languages")
		return
	}
	for _, row := range rows {
		bytes := "n/a"
		if row.hasBytes {
			bytes = fmt.Sprintf("%.2f KiB (%.2f%%)", float64(row.bytes)/1024, row.ratio*100)
		}
		kloc := "n/a"
		if row.hasLines {
			kloc = fmt.Sprintf("%.3f KLOC", float64(row.lines)/1000)
		}
		fmt.Fprintf(w, "  %s: %s | %s\n", row.name, bytes, kloc)
	}
}
