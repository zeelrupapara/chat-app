package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var AllConfig AppConfig

type RedisConfig struct {
	Addr     string `envconfig:"REDIS_ADDR"`
	Password string `envconfig:"REDIS_PASSWORD"`
}

type AppConfig struct {
	Port  string `envconfig:"APP_PORT"`
	Host  string  `envconfig:"APP_HOST"`
	Redis RedisConfig
}

func GetConfig() AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("warning .env file not found, scanning from OS ENV")
	}

	AllConfig = AppConfig{}

	err = envconfig.Process("APP_PORT", &AllConfig)
	if err != nil {
		log.Fatal(err)
	}

	return AllConfig
}
