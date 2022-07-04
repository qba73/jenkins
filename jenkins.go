package jenkins

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "jenkins-go/" + libraryVersion
	mediaType      = "application/json"
)

// Client manages communication with Jenkins API.
type Client struct {
	// HTTP client used to communicate with the Jenkins API.
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// User agent for the client.
	UserAgent string

	// Optional HTTP headers to set on every request to the API.
	headers map[string]string
}

func RunCLI() {
	fmt.Println("Running jctl")
}
