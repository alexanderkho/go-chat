package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	serverAddr := "ws://localhost:8080/ws"
	log.Printf("Connecting to %s", serverAddr)

	log.Println("Enter your username:")
	reader := bufio.NewReader(os.Stdin)
	username, _ := reader.ReadString('\n')

	header := http.Header{
		"username": []string{username},
	}

	// Connect to the server
	conn, _, err := websocket.DefaultDialer.Dial(serverAddr, header)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	done := make(chan struct{})
	// Listen for incoming messages
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				return
			}
			log.Printf("Received: %s", message)
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
