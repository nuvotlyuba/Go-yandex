package repository

import (
	"github.com/nuvotlyuba/Go-yandex/internal/models"
	"github.com/nuvotlyuba/Go-yandex/internal/utils"
)

type Repo struct{}

var data models.URLData

func (r Repo) CreateNewID(responseString string) string {
	id := utils.GenerateToken(8)

	newURL := models.URLItem{ID: id, URL: responseString}
	data = append(data, newURL)
	return id
}

func (r Repo) GetItemByID(id string) (item *models.URLItem, isFind bool) {
	for _, item := range data {
		if item.ID == id {
			return &item, true
		}
	}
	return &models.URLItem{}, false
}
