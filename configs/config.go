package configs

import (
	"log"

	"github.com/caarlos0/env/v10"
)

var BaseURL  = "http://localhost:8080"
var ServerAddress = ":8080"

type Config struct {
	BaseURL       string `env:"BASE_URL"       envDefault:"localhost:8080"`
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:":8080"`
	WriteTimeout  int    `env:"WRITE_TIMEOUT"  envDefault:"10"`
	ReadTimeout   int    `env:"READ_TIMEOUT"   envDefault:"10"`
	DBPort        int    `env:"DB_PORT"        envDefault:"5432"`
	DBUser        string `env:"DB_USER"        envDefault:"user"`
	DBPassword    string `env:"DB_PASSWORD"    envDefault:"password"`
	DBHost        string `env:"DB_HOST"        envDefault:"db"`
}

func LoadConfig() *Config {
	// curDir, err := os.Getwd()
	// if err != nil {
	// 	log.Println(err)
	// }
	// if err := godotenv.Load(curDir + "/.env"); err != nil {
	// 	log.Fatal("unable to load .env file: ", err)
	// }

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal("unable to parse environment variables: ", err)
	}

	return &cfg
}
