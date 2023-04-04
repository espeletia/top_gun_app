package config

import "fmt"

type DBConfig struct {
	ConnectionURI string `json:"-"`
	DriverName    string
	RunMigration  bool
}

func loadDbConfig() DBConfig {
	dbConfig := &DBConfig{}
	v := configViper("db")
	v.BindEnv("ConnectionURI", "DATABASE_URL")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w ", err))
	}
	v.Unmarshal(dbConfig)
	return *dbConfig
}
