package main

import (
	"flag"
	"log"

	"github.com/nuvotlyuba/Go-yandex/config"
	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver"
)

var (
	netAddr = new(config.NetAddress)
)
func init() {
	flag.Var(netAddr, "a", "Net address host:port")
}

func main() {
	_= flag.Value(netAddr)
	flag.Parse()

	config := apiserver.NewConfig()
	config.Set(netAddr.Host, netAddr.Port)

	s := apiserver.New(config.Get())
	if err := s.Start(); err != nil {
		panic(err)
	}
	log.Fatal("Starting server ...")



}
