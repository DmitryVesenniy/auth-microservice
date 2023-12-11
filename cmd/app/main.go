package main

import (
	"fmt"

	"github.com/DmitryVesenniy/auth-microservice/internal/app"
	"github.com/DmitryVesenniy/auth-microservice/internal/config"
	"github.com/DmitryVesenniy/auth-microservice/pkg/logger"
)

func main() {
	cfg, err := config.ConfigLoad()

	if err != nil {
		panic(fmt.Sprintf("[!] Error load config: %s", err))
	}

	log := logger.SetupLogger(cfg.Env)

	// TODO: инициализировать приложение (app)

	// TODO: запустить gRPC-сервер приложения

	app.Run()
}
