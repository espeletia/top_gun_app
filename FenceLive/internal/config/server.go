package config

import "time"

type ServerConfig struct {
	Port            string
	TLSEnable       bool
	TLSCertPath     string
	TLSKeyPath      string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
	DebugPort       string
	DebugEnable     bool
}

func loadServerConfig() ServerConfig {
	serverConfig := &ServerConfig{}
	v := configViper("server")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v.Unmarshal(serverConfig)
	return *serverConfig
}
