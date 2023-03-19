package main

import (
	"FenceLive/cmd/migrations/runner"
	"log"
	"os"
)

func main() {
	if err := runner.RunMigrations(); err != nil {
		log.Println("Migration Error: ", err)
		os.Exit(1)
	}
}