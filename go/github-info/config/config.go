package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Token string
}

func LoadConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("[ERROR]: .env file not found")
		panic(err)
	}

	return &Config{
		Token: getEnv("GITHUB_TOKEN"),
	}
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is required but not set", key)
	}
	return value
}
