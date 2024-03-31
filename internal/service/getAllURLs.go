package service

import (
	"context"

	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
	"github.com/nuvotlyuba/Go-yandex/internal/models"
	"github.com/nuvotlyuba/Go-yandex/internal/utils"
	"go.uber.org/zap"
)

func (s Service) GetAllURLs() (*[]models.URL, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var data *[]models.URL
	var err error

	storage := utils.SwitchStorage()
	logger.Info("get all URLs from", zap.String("storage", storage))
	switch storage {
	case "db":
		data, err = s.dbRepo.GetAllURLs(ctx)
		if err != nil {
			return nil, err
		}
	case "file":
		data, err = s.fileRepo.ReadAllURLs()
		if err != nil {
			return nil, err
		}
	case "mem":
		data, err = s.memRepo.GetAllURLs()
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}
