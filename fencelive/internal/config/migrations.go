package config

import (
	"fmt"
)

type MigrationsConfig struct {
	MigrationPath string
}

func loadMigrationsConfig() MigrationsConfig {
	migrationsConfig := &MigrationsConfig{}
	u := configViper("migrations")
	err := u.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	err = u.Unmarshal(migrationsConfig)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	return *migrationsConfig
}
