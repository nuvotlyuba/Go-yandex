package store

import (
	"context"

	"github.com/nuvotlyuba/Go-yandex/internal/models"
)


type DBRepo interface {
	CreateNewURL(data *models.URL) error
	GetURL(shortURL string) (*models.URL, error)
}

type DBRepository struct {
	store *Store
}

func (r *DBRepository) CreateNewURL(data *models.URL) error {
	return nil
}

func (r DBRepository) GetURL (shortURL string) (*models.URL, error) {
	return &models.URL{}, nil
}

func (r DBRepository) Ping(ctx context.Context) error {
	if err := r.store.db.Ping(ctx); err != nil {
		return err
	}
	return nil
}