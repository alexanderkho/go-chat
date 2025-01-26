package main

import (
	"fmt"
	"go-chat/internal/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var Port = ":8080"

func main() {
	fmt.Println("Server is running on port ", Port)
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/ws", handlers.HandleWebSocket)
	log.Fatal(http.ListenAndServe(Port, r))
}
