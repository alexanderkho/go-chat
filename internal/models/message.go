package models

import (
	"github.com/google/uuid"
)

type Message struct {
	Sender *Sender      `json:"sender"`
	Data   *MessageData `json:"data"`
	Id     uuid.UUID    `json:"id"`
}

type MessageData struct {
	Content     string      `json:"content"`
	MessageType MessageType `json:"type"`
}

type Sender struct {
	Username string    `json:"username"`
	Id       uuid.UUID `json:"id"`
}

type MessageType string

const (
	ClientConnected    MessageType = "client_connected"
	ClientDisconnected MessageType = "client_disconnected"
	ChatMessage        MessageType = "chat_message"
)
