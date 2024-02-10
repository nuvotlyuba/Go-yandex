package main

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/nuvotlyuba/Go-yandex/config"
	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver"
)

var serverAddress string

func init() {
	flag.StringVar(&serverAddress, "a", "", "Server address host:port")
}
func main() {
	flag.Parse()
	serverAddress = parseServerAddress(serverAddress)

	cfg := &config.Config{
		// SeverAddress: config.SeverAddress,
	}

	err := env.Parse(cfg)
	if err != nil {
		log.Fatal(err)
	}

	config := apiserver.NewConfig(serverAddress)

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		panic(err)
	}

}
