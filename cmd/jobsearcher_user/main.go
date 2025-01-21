package main

import (
	"jobsearcher_auth/config"
	"jobsearcher_auth/internal/app"
	"log"
)

const configPath = "config/config.yaml"

func main() {
	// Configuration
	cfg, err := config.New(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
