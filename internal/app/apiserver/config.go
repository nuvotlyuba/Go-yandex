package apiserver

import (
	"time"

	"github.com/nuvotlyuba/Go-yandex/config"
)

type APIConfig struct {
	ServerAddress string
	WriteTimeout  time.Duration
	ReadTimeout   time.Duration
}

func NewConfig() *APIConfig {
	return &APIConfig{
		ServerAddress: config.ServerAddress,
		WriteTimeout:  time.Second * time.Duration(config.WriteTimeout),
		ReadTimeout:   time.Second * time.Duration(config.ReadTimeout),
	}
}
