package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-chat/internal/models"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	cfg "go-chat/pkg/config"

	"github.com/gorilla/websocket"
)

type ClientConfig struct {
	Host string `env:"HOST,required"`
}

func main() {
	config, err := cfg.InitializeConfig[ClientConfig]("cmd/client/.env")
	if err != nil {
		log.Fatal("Error parsing env", err)
	}
	log.Printf("Connecting to %s", config.Host)
	serverAddr := fmt.Sprintf("ws://%s/ws", config.Host)

	fmt.Print("[Enter your username]: ")
	reader := bufio.NewReader(os.Stdin)
	username, _ := reader.ReadString('\n')
	username = strings.ReplaceAll(username, "\n", "")

	url := fmt.Sprintf("%s?username=%s", serverAddr, username)

	// Connect to the server
	conn, _, err := websocket.DefaultDialer.Dial(url, http.Header{})
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	done := make(chan struct{})
	// Listen for incoming messages
	go func() {
		defer close(done)
		for {
			messageType, bytes, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				return
			}

			if messageType == websocket.TextMessage {
				var message models.Message
				err = json.Unmarshal(bytes, &message)
				if err != nil {
					log.Println("Unmarshal error:", err)
					return
				}
				printMessage(message)
			}
		}
	}()

	// Send messages to the server
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	inputScanner := bufio.NewScanner(os.Stdin)
	go func() {
		for {
			inputScanner.Scan()
			err := inputScanner.Err()
			if err != nil {
				log.Fatal(err)
			}
			input := inputScanner.Text()
			err = conn.WriteMessage(websocket.TextMessage, []byte(input))
			if err != nil {
				log.Println("Write error:", err)
				return
			}
		}
	}()

	select {
	case <-done:
		log.Println("Server connection closed, exiting program...")
	case <-interrupt:
		log.Println("Interrupt received, closing connection...")
	}

	err = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("Close error:", err)
	}
}

func printMessage(message models.Message) {
	switch message.Data.MessageType {
	case models.ClientConnected:
		log.Printf("User %s connected", message.Sender.Username)
	case models.ClientDisconnected:
		log.Printf("User %s disconnected", message.Sender.Username)
	case models.ChatMessage:
		log.Printf("[%s]: %s", message.Sender.Username, message.Data.Content)
	}
}
