package store

import "github.com/nuvotlyuba/Go-yandex/internal/models"

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

func (r *VarRepository) FindURL(shortURL string) (*models.URL, error) {
	var result *models.URL
	for _, v := range DataURL {
		if v.ShortURL == shortURL {
			result = v
		}
	}
	return result, nil
}
