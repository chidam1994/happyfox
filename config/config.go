package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type configKey string

var config *viper.Viper
var configKeyMap map[string]configKey

func init() {
	config = viper.New()
	config.SetConfigName("config")
	config.AddConfigPath("./config/")
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	configKeyMap = map[string]configKey{
		"port": PORT,
	}
}

func GetString(key string) string {
	return config.GetString(key)
}

const (
	PORT configKey = "port"
)

func StrToConfigKey(key string) configKey {
	configKey, ok := configKeyMap[key]
	if !ok {
		panic(fmt.Errorf("config key %s not found\n", key))
	}
	return configKey
}
