package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Token    string
	Username string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("load .env: %w", err)
	}
	cfg := &Config{Token: os.Getenv("GITHUB_TOKEN"), Username: os.Getenv("GITHUB_USERNAME")}
	if cfg.Token == "" {
		return nil, fmt.Errorf("GITHUB_TOKEN is required")
	}
	if cfg.Username == "" {
		return nil, fmt.Errorf("GITHUB_USERNAME is required")
	}
	return cfg, nil
}
