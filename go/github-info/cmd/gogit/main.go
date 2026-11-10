package main

import (
	"flag"
	"log"

	"github.com/arthurschuelter/go-git/config"
	"github.com/arthurschuelter/go-git/info"
)

func main() {
	skipKLOC := flag.Bool("skip-kloc", false, "skip counting code lines")
	forceRefresh := flag.Bool("force-refresh", false, "refresh the report even if a cached report is available")
	flag.Parse()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err := info.GetInfoWithRefresh(cfg.Username, cfg, *skipKLOC, *forceRefresh); err != nil {
		log.Fatal(err)
	}
}
