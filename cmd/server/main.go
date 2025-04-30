package server

import (
	"log"
	"web-crawler/internal/api"
	"web-crawler/internal/server"
	"web-crawler/pkg/utils/logger"
)

func StartServer() {
	// Server initialization
	srv := server.NewServer()
	// Logger initialization
	err := logger.Init("pkg/utils/logger/config.yaml")
	if err != nil {
		log.Fatalf("Could not initialize logger: %v", err)
	}
	defer logger.Sync()
	logger := logger.Sugared()
	// API route setting
	api.RegisterRoutes(srv.Router)

	// HTTP-server start
	logger.Info("Starting server on :8080...")
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
