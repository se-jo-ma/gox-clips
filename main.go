package main

import (
	"fmt"
	"log"
	"time"

	internal_api "gox-clips/internal/api" // Importing API for route setup
	"gox-clips/internal/config"
	internal_ui "gox-clips/internal/ui"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	r := gin.Default()
	r.Static("/static", "./assets")
	r.LoadHTMLGlob("internal/ui/templates/*")

	// Set up middleware if needed
	setupMiddleware(r, cfg)

	// Register API routes
	registerRoutes(r)

	// Start the server
	log.Printf("Starting server on :%d\n", cfg.ServerPort)
	if err := r.Run(fmt.Sprintf(":%d", cfg.ServerPort)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupMiddleware(router *gin.Engine, cfg *config.Config) {
	// Example of setting up a middleware
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s\" %d %s%s\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.ErrorMessage,
		)
	}))
}

func registerRoutes(router *gin.Engine) {
	internal_api.RegisterAPIRoutes(router)
	internal_ui.RegisterUIRoutes(router)
}
