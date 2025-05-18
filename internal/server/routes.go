package server

import (
	"net/http"

	"github.com/chiragsoni81245/net-sentinel/internal/types"
)

func NewRouter(server *types.Server) *http.ServeMux {
    mux := http.NewServeMux()
    fileServer := http.FileServer(http.Dir("internal/server/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

    middleware := Middleware{Server: server}

	// UI Routes
	ui := UIControllers{Server: server}
    {
        mux.HandleFunc("/", middleware.AuthMiddleware(ui.Dashboard))
        mux.HandleFunc("/login", ui.Login)
    }


	// Web Socket Route
	ws := WebSocketControllers{Server: server}
    mux.HandleFunc("/ws", ws.HandleWebSocket)

    return mux
}
