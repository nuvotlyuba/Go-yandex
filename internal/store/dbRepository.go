package store

import (
	"context"
	"fmt"

	"github.com/nuvotlyuba/Go-yandex/internal/models"
)


type DBRepo interface {
	CreateNewURL(data *models.URL) error
	GetURL(shortURL string) (*models.URL, error)
}

type DBRepository struct {
	store *Store
}

func (r *DBRepository) CreateNewURL( data *models.URL) error {
	// r.store.db.QueryRow(ctx, "")
	return nil
}

func (r DBRepository) GetURL (shortURL string) (*models.URL, error) {
	return &models.URL{}, nil
}

func (r DBRepository) Ping(ctx context.Context) error {
	fmt.Println(r.store, "&&&&&&&&&&&&&&&&&&&&&&&7")
	if err := r.store.db.Ping(ctx); err != nil {
		return err
	}
	return nil
}
