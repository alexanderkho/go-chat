package chatroom

import (
	"go-chat/internal/models"
	us "go-chat/pkg/user_session"
	"log"
	"sync"

	"github.com/google/uuid"
)

var (
	once                    sync.Once
	chatRoomManagerInstance ChatRoomManager
)

type ChatRoomManager interface {
	AddClient(client *us.UserSession)
	RemoveClient(client *us.UserSession)
	BroadcastMessage(message *models.Message)
}

type chatRoomManager struct {
	clients map[uuid.UUID]*us.UserSession
}

func GetChatRoomManager() ChatRoomManager {
	once.Do(func() {
		chatRoomManagerInstance = &chatRoomManager{clients: make(map[uuid.UUID]*us.UserSession)}
	})
	return chatRoomManagerInstance
}

func (c *chatRoomManager) AddClient(client *us.UserSession) {
	c.clients[client.Id] = client
	c.BroadcastMessage(connectedMessage(client))
	log.Println("Client added to chat room", client.Username)
}

func (c *chatRoomManager) RemoveClient(client *us.UserSession) {
	delete(c.clients, client.Id)
	c.BroadcastMessage(disconnectedMessage(client))
	log.Println("Client removed from chat room", client.Username)
}

func (c *chatRoomManager) BroadcastMessage(message *models.Message) {
	for clientId := range c.clients {
		if clientId != message.Sender.Id {
			client := c.clients[clientId]
			client.PushMessage(message)
		}
	}
}

func ChatMessage(client *us.UserSession, message string) *models.Message {
	return &models.Message{
		Id:     uuid.New(),
		Sender: &models.Sender{Id: client.Id, Username: client.Username},
		Data: &models.MessageData{
			Content:     message,
			MessageType: models.ChatMessage,
		},
	}
}

func connectedMessage(client *us.UserSession) *models.Message {
	return &models.Message{
		Id:     uuid.New(),
		Sender: &models.Sender{Id: client.Id, Username: client.Username},
		Data: &models.MessageData{
			MessageType: models.ClientConnected,
		},
	}
}

func disconnectedMessage(client *us.UserSession) *models.Message {
	return &models.Message{
		Id:     uuid.New(),
		Sender: &models.Sender{Id: client.Id, Username: client.Username},
		Data: &models.MessageData{
			MessageType: models.ClientDisconnected,
		},
	}
}
