package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	PORT         string `mapstructure:"PORT"`
	AppPort      string `mapstructure:"APP_PORT"`
	HOST         string `mapstructure:"HOST"`
	PASSWORD     string `mapstructure:"PASSWORD"`
	USERNAME     string `mapstructure:"USERNAME"`
	DATABASE     string `mapstructure:"DATABASE"`
	PROTOCOL     string `mapstructure:"PROTOCOL"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&c)
	return
}
