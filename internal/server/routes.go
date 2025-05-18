package server

import (
	"net/http"

	"github.com/chiragsoni81245/net-sentinel/internal/types"
	"github.com/chiragsoni81245/net-sentinel/internal/utils"
)

func NewRouter(server *types.Server) *http.ServeMux {
    mux := http.NewServeMux()
    fileServer := http.FileServer(http.Dir("internal/server/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

    middleware := Middleware{Server: server}

	ui := UIControllers{Server: server}
    api := APIControllers{Server: server}
    
    mux.HandleFunc("/", middleware.ProtectedRoute(ui.Dashboard))

    loginMethodHandler := utils.NewMethodHandler()
    loginMethodHandler.Get(ui.Login)
    loginMethodHandler.Post(api.Login)
    mux.HandleFunc("/login", middleware.PublicRoute(loginMethodHandler.Handler))
    mux.HandleFunc("/logout", middleware.ProtectedRoute(ui.Logout))


	// Web Socket Route
	ws := WebSocketControllers{Server: server}
    mux.HandleFunc("/ws", ws.HandleWebSocket)

    return mux
}
