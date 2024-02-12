package apiserver

import (
	"net/http"

	"github.com/nuvotlyuba/Go-yandex/internal/transport/handlers"
)

type APIServer struct {
	config *APIConfig
}

func New(config *APIConfig) *APIServer {
	return &APIServer{
		config: config,
	}
}

func (s *APIServer) Start() error {
	addr := s.config.ServerAddress

	return http.ListenAndServe(addr, handlers.BasicRouter())
}
