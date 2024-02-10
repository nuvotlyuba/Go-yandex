package main

import (
	"os"

	"github.com/nuvotlyuba/Go-yandex/config"
)

// var serverAddress string

// func parseServerAddress() string{
// 	flag.StringVar(&serverAddress, "a", "", "Server address host:port")

// 	if envServerAddress := config.GetEnv("SERVER_ADDRESS", config.SeverAddress); envServerAddress != "" {
// 		serverAddress = envServerAddress
// 	}

// 	return serverAddress
// }

func parseServerAddress(flagServerAddress string) string {
	if flagServerAddress != "" {
		return flagServerAddress
	}

	envServerAddress := os.Getenv("SERVER_ADDRESS")

	if flagServerAddress == "" && envServerAddress != "" {
		return envServerAddress
	}

	return config.ServerAddress
}
