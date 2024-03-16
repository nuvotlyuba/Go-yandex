package service

import (
	"context"

	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
	"github.com/nuvotlyuba/Go-yandex/internal/models"
	"github.com/nuvotlyuba/Go-yandex/internal/utils"
	"go.uber.org/zap"
)

func (s *Service) CreateBatchURL(data models.RequestBatch) (models.BatchURL, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	convData := utils.ToURL(data)
	storage := utils.SwitchStorage()
	switch storage {
	case "db":
		logger.Info("save URL in", zap.String("storage", storage))
		err := s.dbRepo.CreateBatchURL(ctx, convData)
		if err != nil {
			return nil, err
		}
	case "file":
		logger.Info("save URL in", zap.String("storage", storage))
		err := s.fileRepo.WriteBatchURL(convData)
		if err != nil {
			return nil, err
		}
	case "mem":
		logger.Info("save URL in", zap.String("storage", storage))
		s.memRepo.AddBatchURL(convData)
	}

	return convData, nil
}
