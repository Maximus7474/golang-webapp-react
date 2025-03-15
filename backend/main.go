package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/Maximus7474/golang-webapp-react/backend/auth"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/login", auth.LoginHandler).Methods("POST")

	router.Handle("/api/protected", auth.RequireAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.Header.Get("X-User-Email")
		log.Println("/api/protected - Email:", email)
		json.NewEncoder(w).Encode(map[string]string{"message": "Welcome: " + email})
	})))

	router.HandleFunc("/api/logout", auth.LogoutHandler).Methods("POST")

	// Enable CORS
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type"},
	}).Handler(router)

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}
