package apiserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nuvotlyuba/Go-yandex/configs"
	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
	"github.com/nuvotlyuba/Go-yandex/internal/transport/handlers"
	"go.uber.org/zap"
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
		Handler:      service(s.config.LogLevel),
	}

	return server.ListenAndServe()
}

func service(logLevel string) http.Handler {
	if err := logger.Initialize(logLevel); err != nil {
		logger.Log.Fatal("Don't initialize logger")
	}
	logger.Log.Info("Server running ...", zap.String("address", configs.ServerAddress))
	r := chi.NewRouter()
	r.Use(logger.RequestLogger)
	r.Use(middleware.Heartbeat("/ping"))

	return handlers.BasicRouter(r)
}
