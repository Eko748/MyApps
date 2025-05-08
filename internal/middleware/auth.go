package middleware

import (
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Extract and validate token (you can implement more logic here)
		// tokenStr := strings.TrimPrefix(token, "Bearer ")
		// if !validateToken(tokenStr) { ... }

		next.ServeHTTP(w, r)
	})
}
