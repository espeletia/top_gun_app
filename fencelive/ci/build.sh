#!/usr/bin/env bash
set -euo pipefail

go mod tidy
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/run_migrations ./cmd/migrations/main.go
echo "FenceLive migrations build completed"

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/start_server ./cmd/main.go
echo "FenceLive build completed"
