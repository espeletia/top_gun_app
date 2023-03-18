#!/usr/bin/env bash
set -euo pipefail

go mod tidy

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/start_server ./cmd/main.go
echo "FenceLive build completed"
