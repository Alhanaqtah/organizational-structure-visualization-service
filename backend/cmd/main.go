package main

import (
	"organizational-structure-visualization-service/internal/app"
	"organizational-structure-visualization-service/internal/config"
	"organizational-structure-visualization-service/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := logger.New(cfg.ENV)

	app := app.New(cfg, log)

	go app.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	app.Stop()
}
