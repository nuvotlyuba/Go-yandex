package services

import (
	"github.com/nuvotlyuba/Go-yandex/internal/models"
	"github.com/nuvotlyuba/Go-yandex/internal/repository"
)

func (s Service) FindUrl(token string) (*models.URL, error) {

	repo := new(repository.Repo)
	data, err := repo.GetURL(token)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}

	return data, nil

}
