package service

import (
	"context"

	"github.com/nuvotlyuba/Go-yandex/internal/models"
	"github.com/nuvotlyuba/Go-yandex/internal/store"
)

type Service struct {
	dbRepo   *store.DBRepository
	fileRepo *store.FileRepository
	memRepo  *store.MemRepository
}

type Serv interface {
	CreateURL(longURL string) (*models.URL, error)
	FindURL(token string) (*models.URL, error)
	PingDB(ctx context.Context) error
	CreateBatchURL()
	GetAllURLs() ([]*models.URL, error)
}

func New(store *store.Store) *Service {
	return &Service{
		dbRepo:   store.DBRepo(),
		fileRepo: store.FileRepo(),
		memRepo:  store.MemRepo(),
	}
}
