package main

import (
	"jobsearcher_user/config"
	"jobsearcher_user/internal/app"
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
