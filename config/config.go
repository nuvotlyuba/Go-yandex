package config

var (
	BaseURL       = "http://localhost:8080"
	ServerAddress = ":8080"
	WriteTimeout  = 10
	ReadTimeout   = 10
)

type Config struct {
	BaseURL       string `env:"BASE_URL"       envDefault:"http://localhost:8080"`
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	WriteTimeout  int    `env:"WRITE_TIMEOUT"  envDefault:"10"`
	ReadTimeout   int    `env:"READ_TIMEOUT"  envDefault:"10"`
}

func NewConfig() *Config {
	return &Config{}
}
