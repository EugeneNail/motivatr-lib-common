package http

import (
	"net/http"
	"regexp"
)

func DisableLocalCors(next http.Handler) http.Handler {
	var localOrigins = []*regexp.Regexp{
		regexp.MustCompile(`^https?://localhost(:\d+)?$`),
		regexp.MustCompile(`^https?://127\.0\.0\.1(:\d+)?$`),
		regexp.MustCompile(`^https?://192\.168\.\d+\.\d+(:\d+)?$`),
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		origin := r.Header.Get("Origin")
		for _, re := range localOrigins {
			if re.MatchString(origin) {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
				break
			}
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
