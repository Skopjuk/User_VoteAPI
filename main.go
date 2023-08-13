package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"userapi/configs"
	"userapi/container"
	"userapi/repositories"
	"userapi/server"
)

func main() {
	logging := logrus.New()
	logging.SetReportCaller(true)
	logging.Info("create router")

	if err := InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	config := configs.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}

	db, err := server.NewPostgresDB(config)

	if err != nil {
		logrus.Fatalf("cannot connect to db: %s", err.Error())
	}

	containerInstance := container.Container{
		Config:     &config,
		DB:         db,
		Logging:    logging,
		Repository: repositories.NewUsersRepository(db),
	}

	if err := server.Run(viper.GetString("port"), containerInstance); err != nil {
		logrus.Fatalf("error occured while running http server: %s, address: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
