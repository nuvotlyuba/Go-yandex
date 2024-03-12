package service

import (
	"context"
	"time"

	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
	"go.uber.org/zap"
)

func (s Service) PingDB(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := s.dbRepo.Ping(ctx); err != nil {
		return err
	}
	logger.Info("success", zap.String("in", "SERVICE"))

	return nil
}
