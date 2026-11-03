package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	ApiPort  string
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func LoadConfig() *DBConfig {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("[WARN]: .env file not found, using environment variables or defaults")
		panic(err)
	}
	return &DBConfig{
		ApiPort:  getEnv("API_PORT"),
		Host:     getEnv("DB_HOST"),
		Port:     getEnvAsInt("DB_PORT", 5432),
		User:     getEnv("DB_USER"),
		Password: getEnv("DB_PASSWORD"),
		DBName:   getEnv("DB_NAME"),
	}
}

func (c *DBConfig) Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("✓ Database connected successfully!")
	return db, nil

}
func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is required but not set", key)
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Warning: Invalid integer value for %s, using default %d", key, defaultValue)
		return defaultValue
	}
	return value
}
