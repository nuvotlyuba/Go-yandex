package apiserver

import (
	"strconv"

	"github.com/nuvotlyuba/Go-yandex/config"
)

type Config struct {
	BindAddr string
}

func NewConfig() *Config {
	return &Config {
		BindAddr: config.Host + ":" + strconv.Itoa(config.Port),

	}
}

func (c *Config) Set(host string, port int)  {
	if host != "" && port != 0 {
		c.BindAddr = host + ":" + strconv.Itoa(port)
	}
}

func (c Config) Get() *Config {
	return &Config {
		BindAddr: c.BindAddr,
	}
}

