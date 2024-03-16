package apiserver

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nuvotlyuba/Go-yandex/configs"
	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/gzip"
	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
	"github.com/nuvotlyuba/Go-yandex/internal/service"
	"github.com/nuvotlyuba/Go-yandex/internal/store"
	"github.com/nuvotlyuba/Go-yandex/internal/transport/handler"
	"go.uber.org/zap"
)

type APIServer struct {
	config *APIConfig
	logger *zap.Logger
	router *chi.Mux
	db     *pgxpool.Pool
}

func New(config *APIConfig) *APIServer {
	return &APIServer{
		config: config,
		router: chi.NewRouter(),
	}
}

type Server interface {
	Start(ctx context.Context) error
	configureLogger() error
	configureDB(ctx context.Context) error
	closePostgres()
	configureRouter() *chi.Mux
	createTables(ctx context.Context) error
}

func (s *APIServer) Start(ctx context.Context) error {
	if err := s.configureLogger(); err != nil {
		s.logger.Fatal("Don't initialize logger")
	}
	s.logger.Info("Server running ...", zap.String("address", s.config.ServerAddress))

	s.router.Use(logger.LoggerMiddleware)
	s.router.Use(gzip.GzipMiddleware)
	s.router.Use(middleware.Heartbeat("/ping"))

	if configs.DataBaseDSN != "" {
		s.logger.Info("Postgres connecting ...")
		s.logger.Info("postgres env " + s.config.DataBaseDSN)

		if err := s.configureDB(ctx); err != nil {
			s.logger.Fatal("Unable to create connection pool.", zap.Error(err))
		}
		defer s.closePostgres()
		s.logger.Info("Successfully connected to postgreSQL pool.")

		err := s.createTables(ctx)
		if err != nil {
			s.logger.Fatal("Unable to create tables in db.", zap.Error(err))
		}
		s.logger.Info("Successfully created tables.")
	}

	store := store.New(s.db)
	service := service.New(store)
	handler := handler.New(service)

	s.configureRouter(handler)

	server := &http.Server{
		Addr:         s.config.ServerAddress,
		WriteTimeout: s.config.WriteTimeout,
		ReadTimeout:  s.config.ReadTimeout,
		Handler:      s.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (s *APIServer) configureLogger() error {

	lvl, err := zap.ParseAtomicLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	var cfg zap.Config
	if configs.Stage(s.config.AppEnv) == configs.Production {
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

func (s *APIServer) configureDB(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, s.config.DataBaseDSN)
	if err != nil {
		return err
	}

	if err := dbpool.Ping(ctx); err != nil {
		return err
	}

	s.db = dbpool

	return nil
}

func (s *APIServer) closePostgres() {
	s.db.Close()
}

func (s *APIServer) createTables(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS public."shortener" (
			"id"           varchar(100) NOT NULL,
			"short_url"    varchar(100) NOT NULL,
			"original_url" varchar(100) NOT NULL,
			"created"      timestamp(0) NOT NULL,
			CONSTRAINT "shorter.id" PRIMARY KEY ("id")
		);
		CREATE INDEX IF NOT EXISTS short_url_index ON public.shortener (short_url);
	`)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	tx.Commit(ctx)

	return nil
}

func (s *APIServer) configureRouter(h *handler.Handler) *chi.Mux {

	s.router.Post("/", h.PostURLHandler)
	s.router.Get("/{id}", h.GetURLHandler)
	s.router.Post("/api/shorten", h.PostURLJsonHandler)
	s.router.Get("/ping", h.GetConnDBHandler)
	s.router.Post("/api/shorten/batch", h.PostURLBatchHandler)

	walkRout(s.router)

	return s.router
}

func walkRout(r *chi.Mux) {
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		logger.Info(fmt.Sprintf("%s %s\n", method, route))
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		logger.Debug(fmt.Sprintf("Logging err: %s\n", err.Error()))
	}
}
