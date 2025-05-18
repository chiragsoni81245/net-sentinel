package server

import (
	"log"
	"net/http"

	"github.com/chiragsoni81245/net-sentinel/internal/types"
	"github.com/gorilla/websocket"
)


type WebSocketControllers struct {
    Server *types.Server
}

// WebSocket upgrader
// Upgrade HTTP connection to WebSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins for simplicity — restrict in production
		return true
	},
}

func (wsc *WebSocketControllers) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer conn.Close()

	log.Println("Client connected to WebSocket")

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		log.Printf("Received: %s\n", message)

		// Echo message back
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}
