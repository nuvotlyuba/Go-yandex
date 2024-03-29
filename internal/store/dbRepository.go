package store

import (
	"context"
	"fmt"
	"time"

	"github.com/nuvotlyuba/Go-yandex/internal/models"
)

type DBRepo interface {
	CreateNewURL(data *models.URL) error
	GetURL(shortURL string) (*models.URL, error)
	CreateBatchURL(ctx context.Context, data []*models.URL)
}

type DBRepository struct {
	store *Store
}

func (r *DBRepository) CreateNewURL(ctx context.Context, data *models.URL) (string, error) {

	result, err := r.store.db.Exec(ctx,
		`INSERT INTO shortener (id, short_url, original_url, created) VALUES ($1, $2, $3, $4)`,
		data.ID, data.ShortURL, data.OriginalURL, time.Now())

	if result.RowsAffected() == 0 {
		row := r.store.db.QueryRow(ctx, "SELECT short_url FROM shortener WHERE original_url = $1", data.OriginalURL)
		var shortURL string

		err := row.Scan(&shortURL)
		if err != nil {
			return "", fmt.Errorf("error in dbRepository: CreateNewURL -> %v", err)
		}

		return shortURL, ErrConflict
	}

	if err != nil {
		return "", fmt.Errorf("error in dbRepository: CreateNewURL -> %v", ErrCreated)
	}

	return "", nil
}

func (r *DBRepository) GetURL(ctx context.Context, shortURL string) (string, error) {
	row := r.store.db.QueryRow(ctx, "SELECT original_url FROM shortener WHERE short_url = $1", shortURL)
	var originalURL string

	err := row.Scan(&originalURL)
	if err != nil {
		return "", fmt.Errorf("error in dbRepository: GetURL -> %v", ErrQuery)
	}

	return originalURL, nil
}

func (r *DBRepository) Ping(ctx context.Context) error {
	if err := r.store.db.Ping(ctx); err != nil {
		return fmt.Errorf("error in dbRepository: Ping -> %v", err)
	}
	return nil
}

func (r *DBRepository) CreateBatchURL(ctx context.Context, data []*models.URL) error {
	tx, err := r.store.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error in dbRepository: CreateBatchURL begin transaction -> %v", err)
	}
	for _, item := range data {
		_, err := tx.Exec(ctx,
			`INSERT INTO shortener (id, short_url, original_url, created) VALUES ($1, $2, $3, $4)`,
			item.ID, item.ShortURL, item.OriginalURL, time.Now())
		if err != nil {
			tx.Rollback(ctx)
			return ErrCreated
		}

	}
	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("error in dbRepository: CreateBatchURL commit -> %v", err)
	}

	return nil
}
