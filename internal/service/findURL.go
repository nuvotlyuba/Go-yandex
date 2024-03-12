package service

import (
	"context"

	"github.com/nuvotlyuba/Go-yandex/internal/utils"
)

func (s Service) FindURL(token string) (string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shortURL := utils.GetShortURL(token)
	data, err := s.dbRepo.GetURL(ctx, shortURL)

	if err != nil {
		return data, err
	}

	//чтение из файла
	// if configs.FileStoragePath != "" {
	// 	data, err := s.fileRepo.ReadURL(shortenURL)
	// 	if err != nil {
	// 		return &models.URL{}, err
	// 	}
	// 	return data, nil
	// }

	// //чтение из переменной
	// data, err = s.varRepo.FindURL(shortenURL)
	// if err != nil {
	// 	return &models.URL{}, err
	// }

	return data, nil

}
