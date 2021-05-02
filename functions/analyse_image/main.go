package analyse_image

import (
	"cloud.google.com/go/storage"
	vision "cloud.google.com/go/vision/apiv1"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Global API clients used across function invocations.
var (
	storageClient *storage.Client
	visionClient  *vision.ImageAnnotatorClient
)

func init() {
	// Declare a separate err variable to avoid shadowing the client variables.
	var err error

	storageClient, err = storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("storage.NewClient: %v", err)
	}

	visionClient, err = vision.NewImageAnnotatorClient(context.Background())
	if err != nil {
		log.Fatalf("vision.NewAnnotatorClient: %v", err)
	}
}

type BucketInfo struct {
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

type ImageDetails struct {
	Height int      `json:"height"`
	Width  int      `json:"width"`
	Format string   `json:"format"`
	IsCat  bool     `json:"is_cat"`
	Labels []string `json:"labels"`
}

func AnalyseHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := analyseHandler(r)
	if err != nil {
		log.Printf("error: analyzeImage: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Encode: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func analyseHandler(r *http.Request) (*ImageDetails, error) {
	b := BucketInfo{}

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		return nil, fmt.Errorf("Decode: %w", err)
	}

	if b.Bucket == "" {
		return nil, fmt.Errorf("missing 'bucket' parameter in request")
	}

	if b.Name == "" {
		return nil, fmt.Errorf("missing 'name' parameter in request")
	}

	details, err := analyseImage(r.Context(), b.Bucket, b.Name)
	if err != nil {
		return nil, err
	}

	return details, nil
}

func analyseImage(ctx context.Context, bucket string, name string) (*ImageDetails, error) {
	uri := fmt.Sprintf("gs://%s/%s", bucket, name)
	log.Printf("analysing %q", uri)
	image := vision.NewImageFromURI(uri)

	annotations, err := visionClient.DetectLabels(ctx, image, nil, 10)
	if err != nil {
		return nil, err
	}

	var isCat bool
	var labels []string
	for _, annotation := range annotations {
		labels = append(labels, annotation.Description)

		if annotation.Description == "Cat" {
			isCat = true
		}
	}

	return &ImageDetails{
		Height: 102,
		Width:  101,
		Format: "",
		IsCat:  isCat,
		Labels: labels,
	}, nil
}
