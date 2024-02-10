package apiserver

type APIConfig struct {
	ServerAddress string
}

func NewConfig(addr string) *APIConfig {
	return &APIConfig{
		ServerAddress: addr,
	}
}
