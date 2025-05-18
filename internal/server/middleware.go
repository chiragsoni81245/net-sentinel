package server

import (
	"context"
	"net/http"

	"github.com/chiragsoni81245/net-sentinel/internal/types"
	"github.com/chiragsoni81245/net-sentinel/internal/utils"
)

type Middleware struct {
    Server *types.Server
}

func (m *Middleware) ProtectedRoute(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // 1. Check auth (e.g., headers, token, cookie)
        tokenCookie, err := r.Cookie("token")
        if err != nil {
            http.Redirect(w, r, "/login", 302)
            return
        }

        userId, err := utils.ValidateJWT(tokenCookie.Value, m.Server.Config)
        if err != nil {
            http.Redirect(w, r, "/login", 302)
            return
        }

        r = r.WithContext(context.WithValue(r.Context(), "userId", userId))

        next(w, r)
    }
}

func (m *Middleware) PublicRoute(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // 1. Check auth (e.g., headers, token, cookie)
        tokenCookie, err := r.Cookie("token")
        if err != nil {
            next(w, r)
            return
        }

        userId, err := utils.ValidateJWT(tokenCookie.Value, m.Server.Config)
        if err != nil {
            next(w, r)
            return
        }

        r = r.WithContext(context.WithValue(r.Context(), "userId", userId))

        next(w, r)
    }
}

func (m *Middleware) MethodSegregator(method string, next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == method {
            next(w, r)
        }
    }
}
