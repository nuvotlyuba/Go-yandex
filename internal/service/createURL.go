package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
	"github.com/nuvotlyuba/Go-yandex/internal/models"
	"github.com/nuvotlyuba/Go-yandex/internal/utils"
	"go.uber.org/zap"
)

func (s Service) CreateNewURL(longURL string) (*models.URL, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	token := utils.GenerateToken(8)

	newURL := models.URL{
		UUID:        uuid.New(),
		ShortURL:    utils.GetShortURL(token),
		OriginalURL: longURL,
	}

	storage := utils.SwitchStorage()
	switch storage {
	case "db":
		logger.Info("save URL in", zap.String("", storage))
		err := s.dbRepo.CreateNewURL(ctx, &newURL)
		if err != nil {
			return &models.URL{}, err
		}
	case "file":
		logger.Info("save URL in", zap.String("", storage))
		err := s.fileRepo.WriteNewURL(&newURL)
		if err != nil {
			return &models.URL{}, err
		}
	case "mem":
		logger.Info("save URL in", zap.String("", storage))
		err := s.varRepo.AddNewURL(&newURL)
		if err != nil {
			return &models.URL{}, err
		}
	}

	return &newURL, nil
}
