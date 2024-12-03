package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		fmt.Printf("authHeader is: %s", authHeader)
		token := authHeader[len("Bearer "):]

		fmt.Printf("token is: %s", token)

		pass := r.URL.Query().Get("password")
		fmt.Printf("password is: %s", pass)

		if token != pass {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
