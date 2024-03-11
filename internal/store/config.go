package store

import "github.com/nuvotlyuba/Go-yandex/configs"


type Config struct {
	DataBaseDSN string
	FileStoragePath string
}


func NewConfig() *Config {
	return &Config{
		DataBaseDSN: configs.DataBaseDSN,
		FileStoragePath: configs.FileStoragePath,
	}
}
