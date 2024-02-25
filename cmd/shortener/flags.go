package main

import (
	"flag"
	"os"

	"github.com/nuvotlyuba/Go-yandex/configs"
)

var serverAddress   string
var baseURL         string
var fileStoragePath string

func parseFlags() {
	flag.StringVar(&serverAddress, "a", "", "Server address host:port")
	flag.StringVar(&baseURL, "b", "", "Base URL host:port")
	flag.StringVar(&fileStoragePath, "f", "empty", "Full file name, for saving JSON data")
	flag.Parse()

	//serverAddress
	if serverAddress != "" {
		configs.ServerAddress = serverAddress
	}
	envServerAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" && envServerAddress != "" {
		configs.ServerAddress = envServerAddress
	}

	//baseURL
	if baseURL != "" {
		configs.BaseURL = baseURL
	}
	envBaseURL := os.Getenv("BASE_URL")
	if baseURL == "" && envBaseURL != "" {
		configs.BaseURL = envBaseURL
	}

	//fileStoragePath
	envFileStoragePath := os.Getenv("FILE_STORAGE_PATH")
	if envFileStoragePath != "" {
		configs.FileStoragePath  = envFileStoragePath
	}

	if envFileStoragePath == "" && fileStoragePath != "empty" {
		configs.FileStoragePath = fileStoragePath
	}

}
