package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "localhost"
		log.Printf("defaulting to port %s", port)
	}
	host := fmt.Sprintf("%s:%s", baseURL, port)

	h := NewWebsocketHandler(host)
	r.Get("/consumer/init", h.handleInitConsumer)
	r.Get("/publisher/init/{session_id}", h.handleInitPublisher)


	log.Printf("listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}