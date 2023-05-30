package jenkins

import (
	"fmt"
	"net/http"
	"net/url"
)

// Client manages communication with Jenkins API.
type Client struct {
	HTTPClient *http.Client
	BaseURL    *url.URL
}

func RunCLI() {
	fmt.Println("Running jctl")
}
