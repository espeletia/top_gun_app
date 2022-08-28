package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/go-jet/jet/v2/cmd/jet"
	_ "github.com/pressly/goose/v3/cmd/goose"
)

//go run github.com/pressly/goose/v3/cmd/goose postgres postgres://postgres:postgres@localhost:5432/FenceLive?sslmode=disable status
//go run github.com/go-jet/jet/v2/cmd/jet -dsn=postgres://postgres:postgres@localhost:5432/FenceLive?sslmode=disable -path=./internal/ports/database/gen
