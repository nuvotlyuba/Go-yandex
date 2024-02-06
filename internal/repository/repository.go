package repository

import "github.com/nuvotlyuba/Go-yandex/internal/utils"



type UrlItem struct {
	Id      string
	LongUrl string
}

var UrlData []UrlItem

func CreateNewId(responseString string) string {
	id := utils.GenerateToken(8)
	newUrl := UrlItem{Id: id, LongUrl: responseString}
	UrlData = append(UrlData, newUrl)
	return id
}

func GetItemById(id string) (item UrlItem, isFind bool) {
	for _, item := range UrlData {
		if item.Id == id {
			return item, true
		}
	}
	return UrlItem{}, false
}
