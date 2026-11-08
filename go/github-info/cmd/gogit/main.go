package main

import (
	"log"

	"github.com/arthurschuelter/go-git/config"
	"github.com/arthurschuelter/go-git/info"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err := info.GetInfo(cfg.Username, cfg); err != nil {
		log.Fatal(err)
	}
}
