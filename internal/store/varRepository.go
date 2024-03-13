package store

import (
	"github.com/nuvotlyuba/Go-yandex/internal/models"
)

type MemRepository struct {
	data []*models.URL
}

var DataURL []*models.URL

type VarRepo interface {
	AddNewURL(data *models.URL) error
	FindURL(shortURL string) (*models.URL, error)
}

func (r *MemRepository) AddNewURL(data *models.URL) error {
	r.data = append(r.data, data)
	return nil
}

func (r *MemRepository) FindURL(shortURL string) string {
	var data *models.URL
	for _, v := range r.data {
		if v.ShortURL == shortURL {
			data = v
		}
	}
	return data.OriginalURL
}
