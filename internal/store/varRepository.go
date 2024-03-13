package store

import (
	"github.com/nuvotlyuba/Go-yandex/internal/models"
)

type VarRepository struct{}

var DataURL []*models.URL

type VarRepo interface {
	AddNewURL(data *models.URL) error
	FindURL(shortURL string) (*models.URL, error)
}

func (r *VarRepository) AddNewURL(data *models.URL) error {
	DataURL = append(DataURL, data)
	return nil
}

func (r *VarRepository) FindURL(shortURL string) string {
	var data *models.URL
	for _, v := range DataURL {
		if v.ShortURL == shortURL {
			data = v
		}
	}
	return data.OriginalURL
}
