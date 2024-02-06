package apiserver

import (
	"net/http"

	"github.com/nuvotlyuba/Go-yandex/internal/transport/handlers"
)


type APIServer struct {
	config *Config
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

func( s *APIServer) Start() error {
	addr := s.config.BindAddr

	return http.ListenAndServe(addr, handlers.BasicRouter())
}
