package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

	db, err := server.NewPostgresDB(config)

	if err != nil {
		logrus.Fatalf("cannot connect to db: %s", err.Error())
	}

	containerInstance := container.NewContainer(&config, db, logging)
	if err := server.Run(viper.GetString("port"), *containerInstance); err != nil {
		logrus.Fatalf("error occured while running http server: %s, address: %s", err.Error())
	}
}
