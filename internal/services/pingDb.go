package services

import (
	"context"
	"time"

	"github.com/nuvotlyuba/Go-yandex/internal/store"
)

func (s Service) PingDb(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cfg := store.NewConfig()
	r := store.New(cfg)
	if err := r.DBRepo().Ping(ctx); err!= nil {
		return nil
	}

	return nil
}
