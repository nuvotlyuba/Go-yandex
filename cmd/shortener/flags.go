package main

import (
	"flag"
	"os"

	"github.com/nuvotlyuba/Go-yandex/config"
)

var serverAddress string
var baseURL string

func parseFlags() {
	flag.StringVar(&serverAddress, "a", "", "Server address host:port")
	flag.StringVar(&baseURL,       "b", "", "Base URL host:port")
	flag.Parse()

	if serverAddress != "" {
		config.ServerAddress = serverAddress
	}

	if baseURL != "" {
		config.BaseURL = baseURL
	}

	envServerAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" &&  envServerAddress != ""  {
		config.ServerAddress = envServerAddress
	}

	envBaseURL := os.Getenv("BASE_URL")
	if baseURL == "" &&  envBaseURL != "" {
		config.BaseURL = envBaseURL
	}
}
