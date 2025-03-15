package main

import (
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
		w.Write([]byte("Welcome: " + email))
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
