package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// LoggerMiddleware logs the details of the HTTP request
func LoggerMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)

		method := r.Method
		path := r.URL.Path
		logMessage := fmt.Sprintf("%s %s took %v\n", method, path, duration)

		fmt.Print(logMessage)
		w.WriteHeader(http.StatusAccepted)

		w.Write([]byte(logMessage))
	})
}
