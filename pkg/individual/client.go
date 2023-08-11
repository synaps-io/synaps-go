package synaps

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const DefaultBaseURL = "https://api.synaps.io/v4/"

type Client struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

func NewCustomClient(baseURL string, apiKey string) *Client {
	return &Client{
		httpClient: http.DefaultClient,
		apiKey:     apiKey,
		baseURL:    baseURL,
	}
}

func NewClient(apiKey string) *Client {
	return &Client{
		httpClient: http.DefaultClient,
		apiKey:     apiKey,
		baseURL:    DefaultBaseURL,
	}
}

func NewClientFromEnv() *Client {
	godotenv.Load()

	apiKey, ok := os.LookupEnv("SYNAPS_API_KEY")
	if !ok {
		log.Fatalf("Missing required SYNAPS_API_KEY env variable")
	}

	baseURL, ok := os.LookupEnv("SYNAPS_BASE_URL")
	if !ok {
		baseURL = DefaultBaseURL
	}

	return NewCustomClient(baseURL, apiKey)
}

func makeRequest[T any](httpClient *http.Client, method string, path string, body io.Reader, headers map[string]string) (*T, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create init session request: %w", err)
	}

	for key, header := range headers {
		req.Header.Add(key, header)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make init session request: %w", err)
	}

	if (res.StatusCode < 200 || res.StatusCode >= 300) {
		var error Error
		if err := json.NewDecoder(res.Body).Decode(&error); err != nil {
			return nil, fmt.Errorf("failed to unmarshal error output: %w", err)
		}
		defer res.Body.Close()

		return nil, fmt.Errorf("request failed with status code %d: %s", res.StatusCode, error.Message)
	}

	var output T
	if err := json.NewDecoder(res.Body).Decode(&output); err != nil {
		return nil, fmt.Errorf("failed to unmarshal output: %w", err)
	}
	defer res.Body.Close()

	return &output, nil
}
