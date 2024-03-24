package utils

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"strings"

	"github.com/nuvotlyuba/Go-yandex/configs"
	"github.com/nuvotlyuba/Go-yandex/internal/models"
)

func GenerateToken() string {
	b := make([]byte, 4) //equals 8 characters
	rand.Read(b)
	s := hex.EncodeToString(b)
	return s
}

func GetShortURL(id string) string {
	return configs.BaseURL + "/" + id
}

func GetDirsFromPath(path string) string {
	sl := strings.Split(path, "/")
	sl = sl[:len(sl)-1]
	st := strings.Join(sl, "/")
	return st
}

func SwitchStorage() string {
	if configs.DataBaseDSN != "" || os.Getenv("DATABASE_DSN") != "" {
		return "db"
	}
	if configs.FileStoragePath != "" || os.Getenv("FILE_STORAGE_PATH") != "" {
		return "file"
	}
	return "mem"
}

func ToURL(data models.RequestBatch) []*models.URL {
	var result []*models.URL
	for _, item := range data {

		tmp := &models.URL{
			ID:          item.CorrelationID,
			ShortURL:    GetShortURL(item.CorrelationID),
			OriginalURL: item.OriginalURL,
		}
		result = append(result, tmp)
	}
	return result
}

func ToResult(data []*models.URL) models.ResponseBatch {
	result := make([]models.ResponseItem, 0)
	for _, v := range data {
		item := models.ResponseItem{
			CorrelationID: v.ID,
			ShortURL:      v.ShortURL,
		}
		result = append(result, item)
	}
	return result
}
