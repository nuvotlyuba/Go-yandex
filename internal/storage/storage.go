package storage

import "github.com/nuvotlyuba/Go-yandex/internal/utils"

type UrlItem struct {
	Id      string
	LongUrl string
}

var UrlData []UrlItem

func CreateNewShotUrl(responseString string) string {
	token := utils.GenerateToken(8)
	newUrl := UrlItem{Id: token, LongUrl: responseString}
	UrlData = append(UrlData, newUrl)
	return "http://localhost:8080/" + token
}

func GetItemById(id string) (item UrlItem, isFind bool) {
	for _, item := range UrlData {
		if item.Id == id {
			return item, true
		}
	}
	return UrlItem{}, false
}
