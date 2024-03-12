package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/nuvotlyuba/Go-yandex/configs"
	"github.com/nuvotlyuba/Go-yandex/internal/models"
	"github.com/nuvotlyuba/Go-yandex/internal/utils"
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

	err := s.dbRepo.CreateNewURL(ctx, &newURL)
	if err != nil {
		return &models.URL{}, err
	}

	//запись в файл
	if configs.FileStoragePath != "" {
		err := s.fileRepo.InsertNewURL(&newURL)
		if err != nil {
			return &models.URL{}, err
		}
	}

	//запись в переменную
	err = s.varRepo.AddNewURL(&newURL)
	if err != nil {
		return &models.URL{}, err
	}

	return &newURL, err
}
