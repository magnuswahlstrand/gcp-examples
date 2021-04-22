package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("hello: service started")

	http.HandleFunc("/hello", helloHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("hello: received request")

	name := os.Getenv("NAME")
	switch name {
	case "Mars":
		log.Printf("Mars is the only NAME not allowed, please reconfigure")
		panic("Mars is the only NAME not allowed")
	case "":
		name = "世界"
		log.Printf("warning: NAME not set, default to %s", name)
	default:
		// OK
	}

	fmt.Fprintf(w, "Hello %s!\n", name)
}