package chatroom_client

import (
	"encoding/json"
	"go-chat/internal/models"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type IChatroomClient interface {
	PushMessage(message *models.Message)
}

type ChatroomClient struct {
	Username string    `json:"username"`
	Id       uuid.UUID `json:"id"`
	Conn     *websocket.Conn
}

// verify that ChatroomClient implements IChatroomClient
var _ IChatroomClient = &ChatroomClient{}

func NewClient(username string, id uuid.UUID, conn *websocket.Conn) *ChatroomClient {
	return &ChatroomClient{Username: username, Id: id, Conn: conn}
}

func (c *ChatroomClient) PushMessage(message *models.Message) {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		log.Println("Marshal error:", err)
		return
	}
	c.Conn.WriteMessage(websocket.TextMessage, jsonMessage)
}
