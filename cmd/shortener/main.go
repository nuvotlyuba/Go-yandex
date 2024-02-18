package main

import (
	"fmt"

	"github.com/nuvotlyuba/Go-yandex/configs"
	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver"
)

func main() {
	//переменные окружения
	config := configs.LoadConfig()

	//флаги
	parseFlags()

	//конфигурируем сервер
	cfg := apiserver.NewConfig(config)
	fmt.Println(cfg, "cfg main")
	s := apiserver.New(cfg)
	if err := s.Start(); err != nil {
		panic(err)
	}

}
