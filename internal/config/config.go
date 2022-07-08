package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBConfig     DBConfig
	ServerConfig ServerConfig
}

func LoadConfig() *Config {
	config := &Config{
		DBConfig:     loadDbConfig(),
		ServerConfig: loadServerConfig(),
	}

	return config
}

func configViper(configName string) *viper.Viper {
	v := viper.New()
	v.AutomaticEnv()
	v.SetConfigName(configName)
	v.SetConfigType("yaml")
	v.AddConfigPath("./configurations/")
	return v
}
