package auth

import (
	"net/http"

	"backend/utils"
)

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			http.Error(w, "Unauthorized - missing cookie", http.StatusUnauthorized)
			return
		}

		claims, err := utils.ValidateJWT(cookie.Value)
		if err != nil {
			http.Error(w, "Unauthorized - invalid token", http.StatusUnauthorized)
			return
		}

		r.Header.Set("X-User-Email", claims.Email) // (optional) attach user info

		next.ServeHTTP(w, r)
	})
}
