package store

import (
	"github.com/nuvotlyuba/Go-yandex/configs"
	"github.com/nuvotlyuba/Go-yandex/internal/models"
	"github.com/nuvotlyuba/Go-yandex/internal/utils"
)


func (r FileRepository) GetURL(token string)  (*models.URL, error) {
	shortURL := utils.GetShortURL(token)
	var result *models.URL

	//возвращаем из переменной
	if configs.FileStoragePath == "" {
		for _, v := range DataURL {
			if v.ShortURL == shortURL {
				result = v
			}
		}
		return result, nil
	}

	//возвращаем из файла
	rr, err := NewURLScanner(configs.FileStoragePath)
	if err != nil {
		return nil, err
	}
	rr.Split()
	data, err := rr.ScanURL(shortURL)
	if err != nil {
		return nil, err
	}

	return data, nil
}
