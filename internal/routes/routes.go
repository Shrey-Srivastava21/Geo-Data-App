// internal/routes/routes.go
package routes

import (
	"github.com/gorilla/mux"
	"geo-data-app/internal/handlers"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/upload", handlers.UploadGeoJSON).Methods("POST")
	r.HandleFunc("/geodata", handlers.RetrieveGeoJSON).Methods("GET")
}
