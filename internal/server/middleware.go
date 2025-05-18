package server

import "net/http"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // 1. Check auth (e.g., headers, token, cookie)
        token := r.Header.Get("Authorization")
        if token != "Bearer secret-token" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // 2. Call the next handler if authorized
        next(w, r)
    }
}

