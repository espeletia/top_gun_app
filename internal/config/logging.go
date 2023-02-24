package config

import "fmt"

type LoggingConfig struct {
	EnableDebugLogger bool
	EnableFileLogger  bool
	FileLogLevel      string
	FileLogOutput     string
	LoggerConfig      LoggerConfig
}

type LoggerConfig struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Sampling          SamplingConfig
	SamplingEnable    bool
	InitialFields     map[string]interface{}
}

type SamplingConfig struct {
	Initial    int
	Thereafter int
}

func loadLoggingConfig() LoggingConfig {
	loggingConfig := &LoggingConfig{}
	v := configViper("logging")
	v.BindEnv("EnableFileLogger", "ENABLE_FILE_LOGGER")
	v.BindEnv("EnableDebugLogger", "ENABLE_DEBUG_LOGGER")
	v.BindEnv("FileLogOutput", "FILE_LOG_OUTPUT")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w ", err))
	}
	v.Unmarshal(loggingConfig)
	return *loggingConfig
}
