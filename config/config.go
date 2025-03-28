package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server Server `yaml:"server"`
}

type Server struct {
	Port int `yaml:"port"`
}

func GetConfig() (*Config, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	config := &Config{}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return config, fmt.Errorf("config not found")
		} else {
			return config, fmt.Errorf("failed to read config")
		}
	}

	if err := viper.Unmarshal(config); err != nil {
		return config, fmt.Errorf("failed to unmarshal config")
	}

	return config, nil
}
