package main

import (
	"github.com/arthurschuelter/go-git/config"
	"github.com/arthurschuelter/go-git/info"
)

func main() {
	username := "arthurschuelter"
	cfg := config.LoadConfig()

	info.GetInfo(username, cfg)
}
