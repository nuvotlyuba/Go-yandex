package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/nuvotlyuba/Go-yandex/config"
	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver"
)

var (
	netAddr = new(config.NetAddress)
	urlAddr = new(config.UrlAddress)
)
func init() {
	_= flag.Value(netAddr)
	_=flag.Value(urlAddr)
	flag.Var(netAddr, "a", "Net address host:port")
}




func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	config.Set(netAddr.Host, strconv.Itoa(netAddr.Port))


	s := apiserver.New(config.Get())
	if err := s.Start(); err != nil {
		panic(err)
	}
	log.Fatal("Starting server ...")



}
