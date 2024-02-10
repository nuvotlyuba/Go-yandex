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
func parseBaseURL(flagBaseURL string) string {
	if flagBaseURL != "" {
		return flagBaseURL
	}

	envBaseURL := os.Getenv("BASE_URL")

	if flagBaseURL == "" && envBaseURL != "" {
		return envBaseURL
	}

	return config.BaseURL

}
