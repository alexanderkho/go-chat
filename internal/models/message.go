package models

import (
	"go-chat/pkg/chatroom_client"

	"github.com/google/uuid"
)

type Message struct {
	Client  *chatroom_client.ChatroomClient `json:"username"`
	Message string                          `json:"message"`
	Id      uuid.UUID                       `json:"id"`
}
