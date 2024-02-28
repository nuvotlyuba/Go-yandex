package apiserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nuvotlyuba/Go-yandex/configs"
	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/gzip"
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
		Handler:      service(s),
	}

	return server.ListenAndServe()
}

func service(cfg *APIServer) http.Handler {
	if err := logger.Initialize(cfg.config.LogLevel, cfg.config.AppEnv); err != nil {
		logger.Fatal("Don't initialize logger")
	}

	logger.Info("Server running ...", zap.String("address", configs.ServerAddress))

	r := chi.NewRouter()
	r.Use(logger.LoggerMiddleware)
	r.Use(gzip.GzipMiddleware)
	r.Use(middleware.Heartbeat("/ping"))

	return handlers.BasicRouter(r)
}
