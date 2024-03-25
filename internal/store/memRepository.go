package store

import (
	"github.com/nuvotlyuba/Go-yandex/internal/models"
)

type MemRepository struct {
	data []*models.URL
}

var DataURL []*models.URL

type MemRepo interface {
	AddNewURL(data *models.URL)
	FindURL(shortURL string) string
	AddBatchURL(data models.BatchURL)
}

func (r *MemRepository) AddNewURL(data *models.URL) {
	r.data = append(r.data, data)
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

func (r *MemRepository) AddBatchURL(data models.BatchURL) {
	r.data = append(r.data, data...)
}

func (r *MemRepository) GetAllURLs() (*[]models.URL, error) {
	var res []models.URL
	for _, v := range r.data {
		res = append(res, *v)
	}
	if len(res) == 0 {
		return nil, ErrNoContent
	}

	return &res, nil
}
