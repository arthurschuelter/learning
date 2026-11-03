package utils

import (
	"math/rand"
)

func GenerateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

	result := make([]byte, 6)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
