package main

import (
	"log"
	"net/http"
	"ws-chat/internal/handlers"
)

func main() {
	mux := Routes()
	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()

	log.Println("Starting we server on port 8080")

	_ = http.ListenAndServe(":8080", mux)
}
