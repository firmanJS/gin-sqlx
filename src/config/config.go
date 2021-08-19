package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort int
	AppEnv  string
	Debug   bool
	DB      *DB
}

type DB struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	DBSSL    string
}

func AppConfig() (defConfig *Config, err error) {
	viper.SetConfigFile("../.env")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Connont find config file, %s", err)
	}
	defConfig = &Config{}

	appEnv := viper.GetString("GO_ENV")
	appPort := viper.GetInt("GO_PORT")
	debug := viper.GetString("GO_DEUG")

	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetInt("DB_PORT")
	dbName := viper.GetString("DB_NAME")
	dbUser := viper.GetString("DB_USERNAME")
	dbPassword := viper.GetString("DB_PASSWORD")

	defConfig.AppEnv = appEnv
	defConfig.AppPort = appPort
	defConfig.Debug = debug

	dbConfig := &DB{
		Host:     dbHost,
		Port:     dbPort,
		Username: dbUser,
		Password: dbPassword,
		Name:     dbName,
	}

	defConfig.DB = dbConfig

	return defConfig, nil
}
