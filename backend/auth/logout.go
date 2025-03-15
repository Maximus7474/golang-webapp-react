package auth

import (
	"net/http"
	"time"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Set the cookie to expire in the past to delete it
	http.SetCookie(w, &http.Cookie{
		Name:    "auth_token",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0), // Set to Unix epoch (1970 Jan fist)
		MaxAge:  -1,              // Negative max age makes the cookie expire immediately
	})

	w.Write([]byte("Logged out successfully"))
}
