package user_session

import (
	"encoding/json"
	"go-chat/internal/models"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type IUserSession interface {
	PushMessage(message *models.Message)
}

type UserSession struct {
	Username string    `json:"username"`
	Id       uuid.UUID `json:"id"`
	Conn     *websocket.Conn
}

var _ IUserSession = &UserSession{}

func NewUserSession(username string, id uuid.UUID, conn *websocket.Conn) *UserSession {
	return &UserSession{Username: username, Id: id, Conn: conn}
}

func (c *UserSession) PushMessage(message *models.Message) {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		log.Println("Marshal error:", err)
		return
	}
	c.Conn.WriteMessage(websocket.TextMessage, jsonMessage)
}
