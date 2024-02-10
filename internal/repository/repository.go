package repository

import "github.com/nuvotlyuba/Go-yandex/internal/utils"

type URLItem struct {
	ID      string
	LongURL string
}

var URLData []URLItem

func CreateNewID(responseString string) string {
	id := utils.GenerateToken(8)
	newURL := URLItem{ID: id, LongURL: responseString}
	URLData = append(URLData, newURL)
	return id
}

func GetItemByID(id string) (item URLItem, isFind bool) {
	for _, item := range URLData {
		if item.ID == id {
			return item, true
		}
	}
	return URLItem{}, false
}
