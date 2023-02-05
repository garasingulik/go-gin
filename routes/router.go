package routes

import (
	"fmt"
	config "gsk/app/config"
	logger "gsk/app/lib"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Router is our main router instance
var Router *gin.Engine

// Log is the logger
var Log = logger.CreateLogger("HTTP")

func attach(r *gin.Engine) {
	// Health Check
	health(r)
}

// Init - initialize the router
func Init() {
	Router = gin.Default()

	// CORS
	if config.Server.EnableCORS {
		Router.Use(cors.Default())
	}

	// Logger
	var SkipPaths = []string{"/", "/status", "/ping", "/healtz"}
	Router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: SkipPaths,
	}))

	// Attach routes
	attach(Router)

	// Start Listening
	Router.Run(fmt.Sprintf(":%d", config.Server.Port))
}
