package config

import (
	"fmt"
	"time"
)

type MinioConfig struct {
	Endpoint        string
	AccessKeyId     string
	SecretAccessKey string
	UseSSL          bool
	UrlExpiration   time.Duration
	Bucket          string
	Location        string
}

func loadMinioConfig() MinioConfig {
	minioConfig := &MinioConfig{}
	v := configViper("minio")
	err := v.BindEnv("URL", "S3_URL")
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	err = v.BindEnv("Credentials", "S3_CREDENTIALS")
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	err = v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	err = v.Unmarshal(minioConfig)
	if err != nil{
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return *minioConfig
}
