package service

import (
	"context"

	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
	"github.com/nuvotlyuba/Go-yandex/internal/models"
	"github.com/nuvotlyuba/Go-yandex/internal/utils"
	"go.uber.org/zap"
)

func (s Service) CreateURL(longURL string) (*models.URL, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	token := utils.GenerateToken()

	newURL := models.URL{
		ID:          utils.GenerateToken(),
		ShortURL:    utils.GetShortURL(token),
		OriginalURL: longURL,
	}

	storage := utils.SwitchStorage()
	switch storage {
	case "db":
		logger.Info("save URL in", zap.String("storage", storage))
		err := s.dbRepo.CreateNewURL(ctx, &newURL)
		if err != nil {
			return &models.URL{}, err
		}
	case "file":
		logger.Info("save URL in", zap.String("storage", storage))
		err := s.fileRepo.WriteNewURL(&newURL)
		if err != nil {
			return &models.URL{}, err
		}
	case "mem":
		logger.Info("save URL in", zap.String("storage", storage))
		s.memRepo.AddNewURL(&newURL)

	}

	return &newURL, nil
}
