package main

import (
	"cloud.google.com/go/storage"
	"context"
	"embed"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

//go:embed index.html
var fs embed.FS

//go:embed assets/*
var assertsFS embed.FS

// templateData provides template parameters.
type templateData struct {
	Service  string
	Revision string
	Secret   string
}

// Variables used to generate the HTML page.
var (
	data  templateData
	data2 map[string]string
	tmpl  *template.Template

	p12Secret  = os.Getenv("SERVICE_ACCOUNT")
	jsonSecret = os.Getenv("SERVICE_ACCOUNT_JSON")

	bucketName  = os.Getenv("BUCKET_NAME")
)

func main() {
	// Initialize template parameters.
	service := os.Getenv("K_SERVICE")
	if service == "" {
		service = "???"
	}

	revision := os.Getenv("K_REVISION")
	if revision == "" {
		revision = "???"
	}

	// Create the client.
	secret := getSecret(p12Secret)
	log.Printf("Plaintext: %s\n", secret)

	storageClient, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("storage.NewClient: %v", err)
	}

	object := "loan.jpg"
	bucket := storageClient.Bucket(bucketName).Object(object)

	r, err := bucket.NewReader(context.Background())
	if err != nil {
		log.Fatalf("bucket.NewReader: %v", err)
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatalf("ioutil.ReadAll: %v", err)
	}

	accessID := "gcp-cats-demo@appspot.gserviceaccount.com"
	url, err := storage.SignedURL(bucketName, object, &storage.SignedURLOptions{
		GoogleAccessID: accessID,
		PrivateKey:     []byte(secret),
		Method:         "GET",
		Expires:        time.Now().Add(5 * time.Minute),
	})
	if err != nil {
		log.Fatalf("storage.SignedURL: %v", err)
	}

	urlPUT, err := storage.SignedURL(bucketName, uuid.NewString()+"/dog2.jpeg", &storage.SignedURLOptions{
		GoogleAccessID: accessID,
		PrivateKey:     []byte(secret),
		Method:         "PUT",
		Expires:        time.Now().Add(50 * time.Minute),
	})

	if err != nil {
		log.Fatalf("storage.SignedURL: %v", err)
	}

	fmt.Println(url)

	// Prepare template for execution.
	tmpl = template.Must(template.ParseFS(fs, "index.html"))

	data2 = map[string]string{
		"Service": service,
		//"Secret":     secret,
		"URL":        url,
		"PutURL":     urlPUT,
		"BucketName": bucket.ObjectName(),
		"Len":        strconv.Itoa(len(b)),
	}

	// Define HTTP server.
	http.HandleFunc("/", helloRunHandler)
	http.HandleFunc("/form", helloFormHandler)

	fileServer := http.FileServer(http.FS(assertsFS))
	http.Handle("/assets/", fileServer)

	// PORT environment variable is provided by Cloud Run.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Print("Hello from Cloud Run! The container started successfully and is listening for HTTP requests on $PORT")
	log.Printf("Listening on port %s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}

func getSecret(name string) string {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to create secretmanager client: %v", err)
	}

	// Build the request.
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Fatalf("failed to access secret version: %v", err)
	}

	// WARNING: Do not print the secret in a production environment - this snippet
	// is showing how to access the secret material.
	secret := string(result.Payload.Data)
	return secret
}

// helloRunHandler responds to requests by rendering an HTML page.
func helloRunHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.Execute(w, data2); err != nil {
		msg := http.StatusText(http.StatusInternalServerError)
		log.Printf("template.Execute: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
	}
}
func helloFormHandler(w http.ResponseWriter, r *http.Request) {
	secret := getSecret(jsonSecret)
	conf, err := google.JWTConfigFromJSON([]byte(secret))
	if err != nil {
		log.Fatalf("google.JWTConfigFromJSON: %v", err)
	}

	policy, err := generateSignedPostPolicyV4(w, bucketName, "magnus"+uuid.NewString(), conf)
	if err != nil {
		log.Printf("generateSignedPostPolicyV4: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Println(policy.URL)
}


// form is a template for an HTML form that will use the data from the signed
var form = `
<html>
  <body>
	<form action="{{ .URL }}" method="POST" enctype="multipart/form-data">
			{{- range $name, $value := .Fields }}
			<input name="{{ $name }}" value="{{ $value }}" type="hidden"/>
			{{- end }}
			<input type="file" name="file"/><br />
			<input type="submit" value="Upload File" /><br />
	</form>
  </body>
</html>
`

// post policy.

var tmpl2 = template.Must(template.New("policyV4").Parse(form))

// generateSignedPostPolicyV4 generates a signed post policy.
func generateSignedPostPolicyV4(w io.Writer, bucket, object string, conf *jwt.Config) (*storage.PostPolicyV4, error) {
	metadata := map[string]string{
		"x-goog-meta-test":        "data",
	}

	opts := &storage.PostPolicyV4Options{
		GoogleAccessID: conf.Email,
		PrivateKey:     conf.PrivateKey,
		Expires:        time.Now().Add(30 * time.Minute),
		Fields: &storage.PolicyV4Fields{
			Metadata:               metadata,
			//StatusCodeOnSuccess:    201,
			RedirectToURLOnSuccess: "http://localhost:8080",
		},
	}

	policy, err := storage.GenerateSignedPostPolicyV4(bucket, "${filename}", opts)
	if err != nil {
		return nil, fmt.Errorf("storage.GenerateSignedPostPolicyV4: %v", err)
	}

	// Generate the form, using the data from the policy.
	if err = tmpl2.Execute(w, policy); err != nil {
		return policy, fmt.Errorf("executing template: %v", err)
	}

	return policy, nil
}
