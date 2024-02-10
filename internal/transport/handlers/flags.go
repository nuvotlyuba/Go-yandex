package handlers

import (
	"os"

	"github.com/nuvotlyuba/Go-yandex/config"
)

// func parseBaseUrl() string{
// 	var baseUrl string
// 	flag.StringVar(&baseUrl, "b", "", "Base url for short links")
// 	fmt.Println(baseUrl, "baseUrl")

//  	envBaseUrl := os.Getenv("BASE_URL")
//   	if envBaseUrl != "" {
// 		return envBaseUrl
// 	}

// 	if envBaseUrl == "" && baseUrl != "" {
// 		return baseUrl
// 	}

// 	return config.ServerAddress

// }
func parseBaseUrl(flagBaseUrl string) string {
	if flagBaseUrl != "" {
		return flagBaseUrl
	}

	envBaseUrl := os.Getenv("BASE_URL")

	if flagBaseUrl == "" && envBaseUrl != "" {
		return envBaseUrl
	}

	return config.ServerAddress

}
