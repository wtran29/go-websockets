package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/wtran29/go-websockets/internal/handlers"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))
	return mux
}
