package chatroom_client

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type IChatroomClient interface {
	PushMessage(message string)
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

func (c *ChatroomClient) PushMessage(message string) {
	c.Conn.WriteMessage(websocket.TextMessage, []byte(message))
}
