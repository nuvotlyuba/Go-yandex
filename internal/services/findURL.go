package services

import (
	"github.com/nuvotlyuba/Go-yandex/configs"
	"github.com/nuvotlyuba/Go-yandex/internal/models"
	"github.com/nuvotlyuba/Go-yandex/internal/store"
	"github.com/nuvotlyuba/Go-yandex/internal/utils"
)

func (s Service) FindURL(token string) (*models.URL, error) {
	shortenURL := utils.GetShortURL(token)

	cfg := store.NewConfig()
	r := store.New(cfg)
	data, err := r.DBRepo().GetURL(shortenURL)
	if err != nil {
		return &models.URL{}, err
	}

	//чтение из файла
	if configs.FileStoragePath != "" {
		data, err = r.FileRepo().ReadURL(shortenURL)
		if err != nil {
			return &models.URL{}, err
		}
		return data, nil
	}

	//чтение из переменной
	data, err = r.VarRepo().FindURL(shortenURL)
	if err != nil {
		return &models.URL{}, err
	}

	return data, nil

}
