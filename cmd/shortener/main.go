package main

import (
	"context"

	"github.com/nuvotlyuba/Go-yandex/configs"
	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver"
	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
)

func main() {
	//переменные окружения
	config := configs.LoadConfig()

	//флаги
	parseFlags()

	//конфигурируем сервер
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := apiserver.NewConfig(config)
	logger.Info(cfg.DataBaseDSN, "cfg.DataBaseDSN")
	s := apiserver.New(cfg)
	if err := s.Start(ctx); err != nil {
		panic(err)
	}

}
