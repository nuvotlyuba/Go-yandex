package utils

import (
	"math/rand"

	"github.com/nuvotlyuba/Go-yandex/config"
)

func GenerateToken(length int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func StringURL(baseURL, id string) string {
	if baseURL == "" {
		return config.GetEnv("BASE_URL", config.BaseURL) + "/" + id
	}
	return baseURL + "/" + id
}
