package config

import (
	"log"

	"github.com/fhlarif/goblog/app/env"
	"github.com/spf13/viper"
)

var configuration env.Config

func GetConfig() env.Config {
	return configuration
}

func SetConfig() {
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("app/env") // path to look for the config file in

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading the config")
	}

	err := viper.Unmarshal(&configuration)

	if err != nil {
		log.Fatal("unable to decode into struct")
	}
}
