package main

import (
	"log"
	"net/http"
)

func main() {
	mux := Routes()

	log.Println("Starting we server on port 8080")

	_ = http.ListenAndServe(":8080", mux)
}
