package config

import (
	"os"
)

const (
	BaseUrl       = "http://localhost:8080"
	ServerAddress = "localhost:8080"
)

type Config struct {
	BaseUrl      string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	SeverAddress string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
