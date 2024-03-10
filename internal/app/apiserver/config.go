package apiserver

import (
	"time"

	"github.com/nuvotlyuba/Go-yandex/configs"
	"github.com/nuvotlyuba/Go-yandex/internal/store"
)

type APIConfig struct {
	ServerAddress   string
	WriteTimeout    time.Duration
	ReadTimeout     time.Duration
	LogLevel        string
	AppEnv          string
	Store           *store.Config
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
		Store:           store.NewConfig(),
		DataBaseDSN:     configs.DataBaseDSN,
		FileStorageName: configs.FileStoragePath,
	}
}
