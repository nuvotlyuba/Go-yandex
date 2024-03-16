package utils

import (
	"math/rand"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/nuvotlyuba/Go-yandex/configs"
	"github.com/nuvotlyuba/Go-yandex/internal/models"
)

func GenerateToken(length int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
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
		token := GenerateToken(8)

		tmp := &models.URL{
			UUID:        uuid.New(),
			ShortURL:    GetShortURL(token),
			OriginalURL: item.OriginalURL,
		}
		result = append(result, tmp)
	}
	return result
}
