package server

import (
	"context"
	"log"
	"net/http"

	"github.com/chiragsoni81245/net-sentinel/internal/types"
	"github.com/chiragsoni81245/net-sentinel/internal/utils"
)

type Middleware struct {
    Server *types.Server
}

func (m *Middleware) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // 1. Check auth (e.g., headers, token, cookie)
        tokenCookie, err := r.Cookie("token")
        if err != nil {
            log.Println(err)
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        userId, err := utils.ValidateJWT(tokenCookie.Value, m.Server.Config)

        r = r.WithContext(context.WithValue(r.Context(), "userId", userId))

        // 2. Call the next handler if authorized
        next(w, r)
    }
}

