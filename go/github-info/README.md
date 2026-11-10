# gogit

Summarize a GitHub user's public, non-fork repositories, stars, language sizes, and code lines by language.

## Setup

Install Go, Git, and [cloc](https://github.com/AlDanial/cloc#install). Copy `.env.example` to `.env` and set `GITHUB_USERNAME` and `GITHUB_TOKEN`, or provide those variables in your environment. Then run:

```sh
make run
```

To generate a report without cloning repositories or counting code lines, run `make run ARGS=--skip-kloc` (or `./bin/gogit --skip-kloc`). Git and cloc are not required for this option.

The GitHub language ranking measures bytes and uses GitHub's language classifications. KLOC is calculated separately using cloc's code-line counts divided by 1,000. The program shallow-clones each included repository's default branch to a temporary directory and removes it after counting; this can take time and disk space. cloc and GitHub may classify or exclude files differently. Counts describe repository contents, including work by other contributors, rather than lines personally written by the user.

## Cache

The completed report is saved to `.cache/gogit.json` in the current directory, without the token. A report for the configured username is reused for 24 hours from the start of its fetch, without calling GitHub or running Git and cloc. Reports generated with `--skip-kloc` use a separate cache at `.cache/gogit-no-kloc.json`; use `make run ARGS=--force-refresh` to force a refresh. Combine options with `make run ARGS="--skip-kloc --force-refresh"`. Missing, invalid, expired, or different-user caches are refreshed; a failed refresh does not replace the previous cache.

The combined language report uses a curated allowlist of programming and scripting languages. It excludes data, markup, project files, and ambiguous C/C++ headers. GitHub percentages remain shares of all GitHub-reported bytes, including formats hidden from the report. `n/a` means that source did not report that language. Language classifications can differ between GitHub and cloc; JSX counts are grouped with JavaScript.
