package main

import (
	"net/http"
	"ws-chat/internal/handlers"

	"github.com/bmizerany/pat"
)

func Routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))

	return mux
}