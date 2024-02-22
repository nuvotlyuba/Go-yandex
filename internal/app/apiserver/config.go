package apiserver

import (
	"time"

	"github.com/nuvotlyuba/Go-yandex/configs"
)

type APIConfig struct {
	ServerAddress string
	WriteTimeout  time.Duration
	ReadTimeout   time.Duration
	LogLevel      string
}

func NewConfig(cfg *configs.Config) *APIConfig {
	return &APIConfig{
		ServerAddress: configs.ServerAddress,
		WriteTimeout:  time.Second * time.Duration(cfg.WriteTimeout),
		ReadTimeout:   time.Second * time.Duration(cfg.ReadTimeout),
		LogLevel:      cfg.LogLevel,
	}
}
