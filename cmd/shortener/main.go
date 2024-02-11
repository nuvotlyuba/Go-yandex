package main

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/nuvotlyuba/Go-yandex/config"
	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver"
)

func main() {
	cfg := &config.Config{}
	err := env.Parse(cfg)
	if err != nil {
		log.Fatal(err)
	}
	parseFlags()
	config := apiserver.NewConfig(config.ServerAddress)

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		panic(err)
	}

}
