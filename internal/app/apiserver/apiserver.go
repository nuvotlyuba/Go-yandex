package apiserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
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
	server := &http.Server{
		Addr:         s.config.ServerAddress,
		WriteTimeout: s.config.WriteTimeout,
		ReadTimeout:  s.config.ReadTimeout,
		Handler:      service(),
	}

	return server.ListenAndServe()
}

func service() http.Handler {
	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(Logger(), []string{"/ping"}))
	r.Use(middleware.Heartbeat("/ping"))

	return handlers.BasicRouter(r)
}
