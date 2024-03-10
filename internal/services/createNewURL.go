package services

import (
	"github.com/google/uuid"
	"github.com/nuvotlyuba/Go-yandex/internal/models"
	"github.com/nuvotlyuba/Go-yandex/internal/store"
	"github.com/nuvotlyuba/Go-yandex/internal/utils"
)

func (s Service) CreateNewURL(longURL string) (*models.URL, error ) {
	token := utils.GenerateToken(8)

	newURL := models.URL{
		UUID: uuid.New(),
		ShortURL: utils.GetShortURL(token),
		OriginalURL: longURL,
	}

	//запись в базу
	cfg := store.NewConfig()
	r := store.New(cfg)
	err := r.DBRepo().CreateNewURL(&newURL)
	if err != nil {
		return &models.URL{}, err
	}

	//запись в файл
	if cfg.FileStoragePath != "" {
		err = r.FileRepo().InsertNewURL(&newURL)
		if err != nil {
			return &models.URL{}, err
		}
	}

	//запись в переменную
	err = r.VarRepo().AddNewURL(&newURL)
	if err != nil {
		return &models.URL{}, err
	}

	return &newURL, err
}

