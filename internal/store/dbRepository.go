package store

import (
	"context"

	"github.com/nuvotlyuba/Go-yandex/internal/models"
)


type DbRepo interface {
	CreateNewURL(data *models.URL) error
	GetURL(shortURL string) (*models.URL, error)
}

type DbRepository struct {
	store *Store
}

func (r *DbRepository) CreateNewURL(data *models.URL) error {
	return nil
}

func (r DbRepository) GetURL (shortURL string) (*models.URL, error) {
	return &models.URL{}, nil
}

func (r DbRepository) Ping(ctx context.Context) error {
	if err := r.store.db.Ping(ctx); err != nil {
		return err
	}
	return nil
}
