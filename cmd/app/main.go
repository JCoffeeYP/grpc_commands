package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"wsedge_grpc/config"
	"wsedge_grpc/internal/app"
	l "wsedge_grpc/logger"
)

func main() {
	var cfg config.Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatal("Error reading config: ", err)
	}
	logger := l.New(&cfg)

	app.Run(cfg, logger)
}
