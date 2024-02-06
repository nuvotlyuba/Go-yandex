package utils

import (
	"math/rand"
	"strconv"

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

func StringUrl(port int, host, id string) string {
	if host != "" && port != 0 {
		return "http://"+ host + ":" + strconv.Itoa(port) + "/" + id
	}
	return "http://" + config.Host + ":" + strconv.Itoa(config.Port) + "/"+ id
}
