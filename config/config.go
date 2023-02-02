package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Environment          string
	DBSource             string
	DBPort               uint
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

func LoadConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err != viper.ReadInConfig() {
		panic("FAIL TO LOAD CONFIG")
	}

	return config, nil
}
