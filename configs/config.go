package configs

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

type Config struct {
	Port       string
	Host       string
	DBPort     string
	Username   string
	Password   string
	DBName     string
	SSLMode    string
	SigningKey string
	RedisPort  string
	RedisHost  string
	ExpTime    time.Duration
}

func NewConfig() (Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()

	expTime, err := strconv.Atoi(viper.GetString("redis_db.exp_time_seconds"))
	if err != nil {
		logrus.Errorf("error while parsing exp time from config:%s", err)
	}

	return Config{
		Port:       viper.GetString("port"),
		Host:       viper.GetString("db.host"),
		DBPort:     viper.GetString("db.port"),
		Username:   viper.GetString("db.username"),
		Password:   viper.GetString("db.password"),
		DBName:     viper.GetString("db.dbname"),
		SSLMode:    viper.GetString("db.sslmode"),
		SigningKey: viper.GetString("signingKey"),
		RedisPort:  viper.GetString("redis_db.port"),
		RedisHost:  viper.GetString("redis_db.host"),
		ExpTime:    time.Duration(expTime) * time.Second,
	}, err
}
