package handlers

import (
	"go-chat/pkg/chatroom"
	us "go-chat/pkg/user_session"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Upgrader upgrades HTTP connections to WebSocket connections
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins (or implement your own origin check)
		return true
	},
}

// HandleWebSocket handles WebSocket connections
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	// validate request
	if username == "" {
		http.Error(w, "username header is required", http.StatusBadRequest)
		return
	}

	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close() // Ensure the connection is closed when the function ends

	chatRoomManager := chatroom.GetChatRoomManager()

	// Add user to chatroom
	client := us.NewUserSession(username, uuid.New(), conn)
	chatRoomManager.AddClient(client)
	defer chatRoomManager.RemoveClient(client)

	done := make(chan struct{})
	// Start a loop to read messages from the client
	go func() {
		defer close(done)
		for {
			// Read message from the client
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				break
			}

			if messageType == websocket.CloseMessage {
				log.Println("Client closed the connection")
				break
			}
			chatRoomManager.BroadcastMessage(chatroom.ChatMessage(client, string(message)))
		}
	}()

	// block until done channel is closed
	<-done
	log.Println("Client disconnected:", username)
}
