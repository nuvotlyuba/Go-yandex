package store

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)


type Store struct{
	config             *Config
	db                 *pgxpool.Pool
	dbRepository       *DBRepository
	fileRepository     *FileRepository
	varRepository      *VarRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) OpenPostgres(ctx context.Context, dataBaseDSN string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	fmt.Println("$$$$$$$$$$$$$")

	dbpool, err := pgxpool.New(ctx, s.config.DataBaseDSN)
	if err != nil {
		return err
	}

	if err := dbpool.Ping(ctx); err != nil {
		return err
	}

	s.db = dbpool
	fmt.Println(s.db,  "s.db", dbpool, "dfugbju")

	return nil
}

func (s *Store) ClosePostgres() {
	s.db.Close()
}


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
		FileStoragePath: s.config.FileStoragePath,
	}

	return s.fileRepository

}


func (s *Store) VarRepo() *VarRepository {
	if s.varRepository != nil {
		return s.varRepository
	}

	s.varRepository = &VarRepository{}

	return s.varRepository
}
