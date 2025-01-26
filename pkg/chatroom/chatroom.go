package chatroom

import (
	"go-chat/internal/models"
	cc "go-chat/pkg/chatroom_client"
	"log"
	"sync"
)

var (
	once                    sync.Once
	chatRoomManagerInstance ChatRoomManager
)

type ChatRoomManager interface {
	AddClient(client *cc.ChatroomClient)
	RemoveClient(client *cc.ChatroomClient)
	BroadcastMessage(message *models.Message)
}

type chatRoomManager struct {
	clients map[*cc.ChatroomClient]bool
}

func GetChatRoomManager() ChatRoomManager {
	once.Do(func() {
		chatRoomManagerInstance = &chatRoomManager{clients: make(map[*cc.ChatroomClient]bool)}
	})
	return chatRoomManagerInstance
}

func (c *chatRoomManager) AddClient(client *cc.ChatroomClient) {
	c.clients[client] = true
	log.Println("Client added to chat room", client.Username)
}

func (c *chatRoomManager) RemoveClient(client *cc.ChatroomClient) {
	delete(c.clients, client)
	log.Println("Client removed from chat room", client.Username)
}

func (c *chatRoomManager) BroadcastMessage(message *models.Message) {
	for client := range c.clients {
		if client.Id != message.Client.Id {
			client.PushMessage(message.Client.Username + ": " + message.Message)
		}
	}
}
