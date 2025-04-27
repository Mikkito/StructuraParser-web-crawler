package server

import (
	"log"
	"web-crawler/internal/api"
	"web-crawler/internal/server"
)

func StartServer() {
	// Server initialization
	srv := server.NewServer()

	// API route setting
	api.RegisterRoutes(srv.Router)

	// HTTP-server start
	log.Println("Starting server on :8080...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

/*err := logger.Init("./internal/utils/logger/config.yaml")
if err != nil {
	log.Fatalf("Could not initialize logger: %v", err)
}
defer logger.Sync()*/
