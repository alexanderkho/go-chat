package main

import (
	"go-chat/internal/handlers"
	cfg "go-chat/pkg/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ServerConfig struct {
	Port      string `env:"PORT,required"`
	PublicDir string `env:"PUBLIC_DIR,required"`
}

func main() {
	config, err := cfg.InitializeConfig[ServerConfig]("cmd/server/.env")
	if err != nil {
		log.Fatal("Error parsing env", err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/ws", handlers.HandleWebSocket)

	fs := http.FileServer(http.Dir(config.PublicDir))
	r.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	// TODO: don't hardcode `localhost`
	log.Println("Server running on localhost:" + config.Port)
	log.Fatal(http.ListenAndServe("localhost:"+config.Port, r))
}
