package main

import (
	"flag"
	"os"

	"github.com/nuvotlyuba/Go-yandex/configs"
	"github.com/nuvotlyuba/Go-yandex/internal/utils"
)

var serverAddress string
var baseURL string
var fileStoragePath string
var dataBaseDSN string

func parseFlags(cfg *configs.Config) {
	flag.StringVar(&serverAddress, "a", "", "Server address host:port")
	flag.StringVar(&baseURL, "b", "", "Base URL host:port")
	flag.StringVar(&fileStoragePath, "f", "", "Full file name, for saving JSON data")
	flag.StringVar(&dataBaseDSN, "d", "", "Data sourse name for postgresDB")
	flag.Parse()

	//serverAddress
	if serverAddress != "" {
		configs.ServerAddress = serverAddress
	}
	envServerAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" && envServerAddress != "" {
		configs.ServerAddress = envServerAddress
	}
	//нет флага, нет переменной окружения -> пишем по-умолчанию
	if serverAddress == "" && envServerAddress == "" {
		configs.ServerAddress = cfg.ServerAddress
	}


	//baseURL
	if baseURL != "" {
		configs.BaseURL = baseURL
	}
	envBaseURL := os.Getenv("BASE_URL")
	if baseURL == "" && envBaseURL != "" {
		configs.BaseURL = envBaseURL
	}
	//нет флага, нет переменной окружения -> пишем по-умолчанию
	if baseURL == "" && envBaseURL == "" {
		configs.BaseURL = cfg.BaseURL
	}

	//fileStoragePath
	//создаем папку
	envFileStoragePath := os.Getenv("FILE_STORAGE_PATH")
	if envFileStoragePath != "" {
		configs.FileStoragePath = envFileStoragePath
		os.MkdirAll(utils.GetDirsFromPath(configs.FileStoragePath), 0777)
	}

	//создаем папку из флага
	if envFileStoragePath == "" && fileStoragePath != "" {
		configs.FileStoragePath = fileStoragePath
		os.MkdirAll(utils.GetDirsFromPath(configs.FileStoragePath), 0777)

	}

	//не создаем папку
	if envFileStoragePath == "" && fileStoragePath != "" {
		configs.FileStoragePath = fileStoragePath
	}


	//dataBaseDSN
	if dataBaseDSN != "" {
		configs.DataBaseDSN = dataBaseDSN
	}
	envDataBaseDSN := os.Getenv("DATABASE_DSN")
	if envDataBaseDSN != "" {
		configs.DataBaseDSN = envDataBaseDSN
	}
	if  dataBaseDSN == "" && envDataBaseDSN == "" {
		configs.DataBaseDSN = cfg.DataBaseDSN
	}
}
