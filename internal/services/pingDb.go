package services

import (
	"context"
	"fmt"
	"time"

	"github.com/nuvotlyuba/Go-yandex/internal/store"
)

func (s Service) PingDB(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
fmt.Println("************")
	cfg := store.NewConfig()
	r := store.New(cfg)
	if err := r.DBRepo().Ping(ctx); err!= nil {
		return nil
	}

	return nil
}
