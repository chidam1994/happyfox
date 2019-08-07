package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type configKey string

var config *viper.Viper

func init() {
	config = viper.New()
	config.SetConfigName("config")
	config.AddConfigPath("./config/")
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func GetString(key configKey) string {
	return config.GetString(string(key))
}

const (
	PORT        configKey = "port"
	DB_HOST     configKey = "database.host"
	DB_USERNAME configKey = "database.username"
	DB_PASSWORD configKey = "database.password"
	DB_DBNAME   configKey = "database.dbname"
	DB_PORT     configKey = "database.port"
)
