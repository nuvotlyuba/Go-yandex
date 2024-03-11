package store


type Config struct {
	DataBaseDSN string
	FileStoragePath string
}


func NewConfig() *Config {
	return &Config{

	}
}
