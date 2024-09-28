package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_SSL_MODE string
	DB_DRIVER   string
	PORT        string
}

func New() *Config {
	viper.SetConfigFile("config.env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("cant load cfg")
	}

	return &Config{
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_SSL_MODE"),
		viper.GetString("DB_DRIVER"),
		viper.GetString("PORT"),
	}
}
