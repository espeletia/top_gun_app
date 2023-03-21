//go:build tools
// +build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/go-jet/jet/v2/cmd/jet"
	_ "github.com/pressly/goose/v3/cmd/goose"
)