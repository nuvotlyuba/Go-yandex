package service

import (
	"context"

	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
	"github.com/nuvotlyuba/Go-yandex/internal/utils"
	"go.uber.org/zap"
)

func (s Service) FindURL(token string) (string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shortURL := utils.GetShortURL(token)

	var data string
	var err error

	storage := utils.SwitchStorage()
	switch storage {
	case "db":
		logger.Info("get URL from", zap.String("storage", storage))
		data, err = s.dbRepo.GetURL(ctx, shortURL)
		if err != nil {
			return data, err
		}
	case "file":
		logger.Info("get URL from", zap.String("storage", storage))
		data, err = s.fileRepo.ReadURL(shortURL)
		if err != nil {
			return data, err
		}
		return data, nil
	case "mem":
		logger.Info("get URL from", zap.String("storage", storage))
		data = s.varRepo.FindURL(shortURL)
		return data, nil
	}

	return data, nil

}
