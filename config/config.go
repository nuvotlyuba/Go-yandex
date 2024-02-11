package config

import (
	"os"
)

var (
	BaseURL       = "http://localhost:8080"
	ServerAddress = "localhost:8080"
)

type Config struct {
	BaseURL       string `env:"BASE_URL"       envDefault:"http://localhost:8080"`
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
}

func NewConfig() *Config {
	return &Config{}
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
