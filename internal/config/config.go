package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerConfig  ServerConfig
	LoggingConfig LoggingConfig
	ServiceConfig ServiceConfig
	DBConfig      DBConfig
	MinioConfig   MinioConfig
}

func LoadConfig() *Config {
	config := &Config{
		ServerConfig:  loadServerConfig(),
		LoggingConfig: loadLoggingConfig(),
		ServiceConfig: loadServiceConfig(),
		DBConfig:      loadDbConfig(),
		MinioConfig:   loadMinioConfig(),
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
