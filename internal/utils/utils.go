package utils

import (
	"math/rand"

	"github.com/nuvotlyuba/Go-yandex/config"
)


func GenerateToken(length int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func StringUrl(baseUrl, id string) string {
	if baseUrl != "" {
		return "http://"+ baseUrl + "/" + id
	}
	return "http://" + config.GetDefaultBaseUrl() + "/" + id
}
