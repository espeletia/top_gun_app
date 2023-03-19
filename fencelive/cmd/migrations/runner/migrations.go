package runner

import (
	"FenceLive/internal/config"
	"FenceLive/internal/setup"
	"context"

	"github.com/wexder/goose/v3"
)

func RunMigrations() error {
	configuration := config.LoadConfig()
	return RunMigrationsWithConfiguration(*configuration)
}

func RunMigrationsWithConfiguration(configuration config.Config) error {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "config", configuration)

	dBconn, err := setup.SetupDb(&configuration)
	if err != nil {
		return err
	}

	if err := goose.UpCtx(ctx, dBconn, configuration.MigrationsConfig.MigrationPath); err != nil {
		return err
	}

	return nil
}
