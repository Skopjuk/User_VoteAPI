package main

import (
	"github.com/sirupsen/logrus"
	"userapi/configs"
	"userapi/container"
	"userapi/server"
)

func main() {
	logging := logrus.New()
	logging.SetReportCaller(true)
	logging.Info("create router")

	config, err := configs.NewConfig()

	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	containerInstance := *container.NewContainer(config, logging)
	if err := server.Run(config.Port, containerInstance); err != nil {
		logrus.Fatalf("error occured while running http server: %s, address: %s", err.Error(), config.Port)
	}
}
