package main

import (
	"gsk/app/config"
	logger "gsk/app/lib"
	router "gsk/app/routes"
)

func main() {
	// Create logger
	log := logger.CreateLogger("main")

	// Init Config
	if err := config.Init(); err != nil {
		log.Errorf("%v", err)
		return
	}

	// Logging environment
	log.Infof("ENVIRONMENT: %v", config.App.Environment)

	// Init Routes
	router.Init()
}
