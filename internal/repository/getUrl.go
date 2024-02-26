package repository

import (
	"fmt"

	"github.com/nuvotlyuba/Go-yandex/configs"
	"github.com/nuvotlyuba/Go-yandex/internal/models"
	"github.com/nuvotlyuba/Go-yandex/internal/utils"
)


func (r Repo) GetURL(token string)  (*models.URL, error) {
	shortUrl := utils.GetShortURL(token)
	var result *models.URL

	//возвращаем из переменной
	if configs.FileStoragePath == "" {
		for _, v := range DataUrl {
			if v.ShortURL == shortUrl {
				result = v
			}
		}
		return result, nil
	}
	//возвращаем из файла
	fmt.Println(configs.FileStoragePath, "file path in repo")
	rr, err := newURLReader(configs.FileStoragePath)
	if err != nil {
		return nil, err
	}
	rr.Split()
	data, err := rr.ReadURL(shortUrl)
	if err != nil {
		return nil, err
	}

	return data, nil
}
