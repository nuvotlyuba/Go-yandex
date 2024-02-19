package main

import (
	"flag"
	"os"

	"github.com/nuvotlyuba/Go-yandex/configs"
)

var serverAddress string
var baseURL string

func parseFlags() {
	flag.StringVar(&serverAddress, "a", "", "Server address host:port")
	flag.StringVar(&baseURL, "b", "", "Base URL host:port")
	flag.Parse()

	if serverAddress != "" {
		configs.ServerAddress = serverAddress
	}

	if baseURL != "" {
		configs.BaseURL = baseURL
	}

	envServerAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" && envServerAddress != "" {
		configs.ServerAddress = envServerAddress
	}

	envBaseURL := os.Getenv("BASE_URL")

	if baseURL == "" && envBaseURL != "" {
		configs.BaseURL = envBaseURL
	}
}
