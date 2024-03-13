package configs

import (
	"log"

	"github.com/caarlos0/env/v10"
)

type Stage string

const (
	Development Stage = "development"
	Production  Stage = "production"
)

var BaseURL = "http://localhost:8080"
var ServerAddress = ":8080"
var FileStoragePath = ""
var DataBaseDSN = "postgres://postgres:user@localhost:5432/shortener"

type Config struct {
	AppEnv          string `env:"APP_ENV"            envDefault:"development"`
	BaseURL         string `env:"BASE_URL"           envDefault:"localhost:8080"`
	ServerAddress   string `env:"SERVER_ADDRESS"     envDefault:":8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"  envDefault:"/tmp/short-url-db.json"`
	LogLevel        string `env:"LOG_LEVEL"          envDefault:"debug"`
	WriteTimeout    int    `env:"WRITE_TIMEOUT"      envDefault:"10"`
	ReadTimeout     int    `env:"READ_TIMEOUT"       envDefault:"10"`
	DBPort          int    `env:"DB_PORT"            envDefault:"5432"`
	DBUser          string `env:"DB_USER"            envDefault:"user"`
	DBPassword      string `env:"DB_PASSWORD"        envDefault:"password"`
	DBHost          string `env:"DB_HOST"            envDefault:"db"`
	DataBaseDSN     string `env:"DATABASE_DSN"       envDefault:"postgres://postgres:user@localhost:5432/shortener"`
}

func LoadConfig() *Config {
	// curDir, _ := os.Getwd()

	// if err := godotenv.Load(curDir + "/.env"); err != nil {
	// 	log.Printf("unable to load .env file: %v", err)
	// }

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("unable to parse environment variables: %v", err)
	}

	return &cfg
}
