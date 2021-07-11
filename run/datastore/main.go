package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	serviceName = "customer-profile"
)

var datastoreClient *datastore.Client

func main() {
	ctx := context.Background()

	log.Printf("%s: service started", serviceName)

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	var err error
	datastoreClient, err = datastore.NewClient(ctx, "mbw-explorer-prod")
	if err != nil {
		log.Fatal(err)
	}

	err = recordVisit(ctx, time.Now(), "127.0.0.2")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

type visit struct {
	Timestamp time.Time
	UserIP    string
}

func recordVisit(ctx context.Context, now time.Time, userIP string) error {
	v := &visit{
		Timestamp: now,
		UserIP:    userIP,
	}

	k := datastore.IncompleteKey("Visit", nil)

	_, err := datastoreClient.Put(ctx, k, v)
	return err
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s says hello", serviceName)
}
