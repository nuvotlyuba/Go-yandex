package store

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nuvotlyuba/Go-yandex/configs"
)

type Store struct {
	db             *pgxpool.Pool
	dbRepository   *DBRepository
	fileRepository *FileRepository
	memRepository  *MemRepository
}

func New(db *pgxpool.Pool) *Store {
	return &Store{
		db: db,
	}
}

var ErrConflict = errors.New("data conflict")
var ErrCreated = errors.New("unable to save data in db")
var ErrQuery = errors.New("unable to exec query in db")

func (s *Store) DBRepo() *DBRepository {
	if s.dbRepository != nil {
		return s.dbRepository
	}

	s.dbRepository = &DBRepository{
		store: s,
	}

	return s.dbRepository
}

func (s *Store) FileRepo() *FileRepository {
	if s.fileRepository != nil {
		return s.fileRepository
	}

	s.fileRepository = &FileRepository{
		FileStoragePath: configs.FileStoragePath,
	}

	return s.fileRepository

}

func (s *Store) MemRepo() *MemRepository {
	if s.memRepository != nil {
		return s.memRepository
	}

	s.memRepository = &MemRepository{
		data: DataURL,
	}

	return s.memRepository
}
