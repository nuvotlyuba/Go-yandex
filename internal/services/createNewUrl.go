package services

import (
	"github.com/google/uuid"
	"github.com/nuvotlyuba/Go-yandex/internal/models"
	"github.com/nuvotlyuba/Go-yandex/internal/repository"
	"github.com/nuvotlyuba/Go-yandex/internal/utils"
)

func (s Service) CreateNewURL(longUrl string) (*models.URL, error ) {
	token := utils.GenerateToken(8)

	newURL := models.URL{
		UUID: uuid.New(),
		ShortURL: utils.GetShortURL(token),
		OriginalUrl: longUrl,
	}

	repo := new(repository.Repo)
	err := repo.InsertNewUrl(&newURL)

	return &newURL, err
}

