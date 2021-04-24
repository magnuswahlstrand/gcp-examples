package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	serviceName = "customer-profile"
)

func main() {
	log.Printf("%s: service started", serviceName)

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s says hello", serviceName)
}
