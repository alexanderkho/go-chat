package handlers

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home handler called")
	w.Write([]byte("Welcome to the Home Page!"))
}
