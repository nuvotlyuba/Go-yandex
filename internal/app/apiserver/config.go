package apiserver

import (
	"time"

	"github.com/nuvotlyuba/Go-yandex/configs"
)

type APIConfig struct {
	ServerAddress   string
	WriteTimeout    time.Duration
	ReadTimeout     time.Duration
	LogLevel        string
	AppEnv          string
	DataBaseDSN     string
	FileStorageName string
}

func NewConfig(cfg *configs.Config) *APIConfig {
	return &APIConfig{
		ServerAddress:   configs.ServerAddress,
		WriteTimeout:    time.Second * time.Duration(cfg.WriteTimeout),
		ReadTimeout:     time.Second * time.Duration(cfg.ReadTimeout),
		LogLevel:        cfg.LogLevel,
		AppEnv:          cfg.AppEnv,
		DataBaseDSN:     configs.DataBaseDSN,
		FileStorageName: configs.FileStoragePath,
	}
}
