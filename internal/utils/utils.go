package utils

import (
	"math/rand"
	"strings"

	"github.com/nuvotlyuba/Go-yandex/configs"
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
