package apiserver

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nuvotlyuba/Go-yandex/configs"
	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/gzip"
	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
	"github.com/nuvotlyuba/Go-yandex/internal/store"
	"github.com/nuvotlyuba/Go-yandex/internal/transport/handlers"
	"go.uber.org/zap"
)

type APIServer struct {
	config      *APIConfig
	logger      *zap.Logger
	router      *chi.Mux
	store       *store.Store
	fileWriter  *store.URLRecorder
	fileReader  *store.URLScanner
}

func New(config *APIConfig) *APIServer {
	return &APIServer{
		config: config,
		logger: zap.New(nil),
		router: chi.NewRouter(),
	}
}

func (s *APIServer) Start(ctx context.Context) error {
	if err := s.configureLogger(s.config.LogLevel, s.config.AppEnv); err != nil {
		s.logger.Fatal("Don't initialize logger")
	}
	s.logger.Info("Server running ...", zap.String("address", configs.ServerAddress))

	s.router.Use(logger.LoggerMiddleware)
	s.router.Use(gzip.GzipMiddleware)
	s.router.Use(middleware.Heartbeat("/ping"))

	s.logger.Info("Postgres connecting ...")
	s.logger.Info("postgres env "+s.config.DataBaseDSN)
	if err := s.configureRepository(ctx); err != nil {
		s.logger.Fatal("Unable to create connection pool.", zap.Error(err))
	}
	s.logger.Info("Successfully connected to postgreSQL pool.")


	server := &http.Server{
		Addr:         s.config.ServerAddress,
		WriteTimeout: s.config.WriteTimeout,
		ReadTimeout:  s.config.ReadTimeout,
		Handler:      BasicRouter(s),
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}


func (s *APIServer) configureLogger(level string, appEnv string) error {

	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return err
	}
	var cfg zap.Config
	if configs.Stage(appEnv) == configs.Production {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}
	cfg.Level = lvl

	zl, err := cfg.Build()
	if err != nil {
		return err
	}

	s.logger = zl

	zap.ReplaceGlobals(s.logger)

	return nil
}


func (s *APIServer) configureRepository(ctx context.Context) error {
	r := store.New(s.config.Store)
	if err := r.OpenPostgres(ctx, s.config.DataBaseDSN); err != nil {
		return err
	}

	s.store = r

	defer s.store.ClosePostgres()

	return nil
}

func (s *APIServer) configureFile() error {
	w, err := store.NewURLRecorder(s.config.FileStorageName)
	if err != nil {
		return err
	}
	defer w.Close()

	s.fileWriter = w

	r, err := store.NewURLScanner(s.config.FileStorageName)
	if err != nil {
		return err
	}
	s.fileReader = r

	return nil
}


func BasicRouter(s *APIServer) chi.Router {

	h := handlers.New()
	s.router.Post("/", h.PostURLHandler)
	s.router.Get("/{id}", h.GetURLHandler)
	s.router.Post("/api/shorten", h.PostURLJsonHandler)
	s.router.Get("/ping", h.GetConnDBHandler)

	WalkRout(s.router)

	return s.router
}

func  WalkRout(r *chi.Mux) {
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		logger.Info(fmt.Sprintf("%s %s\n", method, route))
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		logger.Debug(fmt.Sprintf("Logging err: %s\n", err.Error()))
	}
}
