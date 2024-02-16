package main

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/nuvotlyuba/Go-yandex/config"
	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver"
)

func main() {
	//переменные окружения
	cfg := &config.Config{}
	err := env.Parse(cfg)
	if err != nil {
		log.Fatal(err)
	}
	//флаги
	parseFlags()

	config := apiserver.NewConfig()
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		panic(err)
	}

}
