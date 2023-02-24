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
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	v.Unmarshal(minioConfig)
	return *minioConfig
}
