package apiserver

import "strconv"

const (
	Port = 8080
	Host = "localhost"
)
type Config struct {
	BindAddr string
}

func NewConfig() *Config {
	return &Config {
		BindAddr: Host + ":" + strconv.Itoa(Port),

	}
}

func (c *Config) Set(host, port string)  {
	if host != "" && port != "" {
		c.BindAddr = host + ":" + port
	}
}

func (c Config) Get() *Config {
	return &Config {
		BindAddr: c.BindAddr,
	}
}

