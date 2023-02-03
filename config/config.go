package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Environment string
	Database    Database
	Token       Token
}

type Database struct {
	Driver   string
	Username string
	Password string
	Source   string
	Port     uint64
}

type Token struct {
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

func LoadConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return config, err
}
