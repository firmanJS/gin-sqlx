package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("../.env")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Connont find config file, %s", err)
	}
}

type Config struct {
	GO_PORT     int
	GO_ENV      string
	DB_URI      string
	GO_DEBUG    bool
	DB_PORT     int
	DB_HOST     string
	DB_NAME     string
	DB_USERNAME string
	DB_PASSWORD string
}

func NewConfig() (defConfig *Config, err error) {
	defConfig = &Config{}

	appPort := viper.GetInt("GO_PORT")
	appEnv := viper.GetString("GO_ENV")
	appDebug := viper.GetBool("GO_DEBUG")
	dbPort := viper.GetInt("DB_PORT")
	dbHost := viper.GetString("DB_HOST")
	dbName := viper.GetString("DB_NAME")
	dbUsername := viper.GetString("DB_USERNAME")
	dbPassword := viper.GetString("DB_PASSWORD")

	defConfig.GO_ENV = appEnv
	defConfig.GO_PORT = appPort
	defConfig.GO_DEBUG = appDebug
	defConfig.DB_HOST = dbHost
	defConfig.DB_NAME = dbName
	defConfig.DB_USERNAME = dbUsername
	defConfig.DB_PASSWORD = dbPassword
	defConfig.DB_PORT = dbPort

	return defConfig, nil
}
