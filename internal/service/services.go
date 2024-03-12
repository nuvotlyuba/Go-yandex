package service

import (
	"context"

	"github.com/nuvotlyuba/Go-yandex/internal/models"
	"github.com/nuvotlyuba/Go-yandex/internal/store"
)

type Service struct {
	dbRepo   *store.DBRepository
	fileRepo *store.FileRepository
	varRepo  *store.VarRepository
}

type Serv interface {
	CreateNewURL(longURL string) (*models.URL, error)
	FindURL(token string) (*models.URL, error)
	PingDB(ctx context.Context) error
}

func New(store *store.Store) *Service {
	return &Service{
		dbRepo:   store.DBRepo(),
		fileRepo: store.FileRepo(),
		varRepo:  store.VarRepo(),
	}
}
