package middleware

import (
	"net/http"

	"github.com/lukamandic/logistics/backend/internal/config"
)

func CORSMiddleware(cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", cfg.UIURL)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Max-Age", "3600")

			if r.Method == http.MethodOptions {
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
				w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Range")
				w.WriteHeader(http.StatusOK)
				return
			}

			// For actual requests, set additional headers
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Content-Type", "application/json")

			next.ServeHTTP(w, r)
		})
	}
} 