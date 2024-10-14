// cmd/server/main.go
package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"geo-data-app/internal/config"
	"geo-data-app/internal/database"
	"geo-data-app/internal/routes"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to the database
	database.ConnectDatabase(cfg.DbUrl)

	// Set up router
	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	// Start server
	log.Printf("Starting server on port %s...", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
