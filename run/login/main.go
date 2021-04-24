package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/compute/metadata"
)

const (
	serviceName = "login"
)

func main() {
	log.Printf("%s: service started", serviceName)

	http.HandleFunc("/", handler)
	http.HandleFunc("/check", checkHandler)

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

func makeGetRequest(serviceURL string) (*http.Response, error) {
	// query the id_token with ?audience as the serviceURL
	tokenURL := fmt.Sprintf("/instance/service-accounts/default/identity?audience=%s", serviceURL)
	idToken, err := metadata.Get(tokenURL)
	if err != nil {
		return nil, fmt.Errorf("metadata.Get: failed to query id_token: %+v", err)
	}
	req, err := http.NewRequest("GET", serviceURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", idToken))
	return http.DefaultClient.Do(req)
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	serviceURL := os.Getenv("CUSTOMER_PROFILE_SERVICE_URL")

	resp, err := makeGetRequest(serviceURL)
	if err != nil {
		log.Printf("http.Get: %v", err)
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("unexpected status code: %v", resp.StatusCode)
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	fmt.Fprintf(w, "%s says logged in hello", serviceName)
}
