package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"userapi/handlers"
)

func main() {
	logging := logrus.New()
	logging.Info("create router")

	if err := InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	srv := new(Server)
	handler := handlers.NewHandler(logging)

	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
