package individual

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	. "github.com/synaps.io/synaps-sdk-go/pkg/individual/models"
)

type Client struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

type IndividualService interface {
	InitSession() (sessionID string, err error)
	SessionDetails(sessionID string) (SessionDetailsResponse, error)
	StepDetails(sessionID string, stepID string) (StepDetailsResponse, error)
}

func NewClient(baseURL string, apiKey string) *Client {
	return &Client{
		httpClient: http.DefaultClient,
		apiKey:     apiKey,
		baseURL:    baseURL,
	}
}

func NewClientFromEnv() *Client {
	return &Client{
		httpClient: http.DefaultClient,
		apiKey:     os.Getenv("SYNAPS_API_KEY"),
		baseURL:    os.Getenv("SYNAPS_BASE_URL"),
	}
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

	var output T
	if err := json.NewDecoder(res.Body).Decode(&output); err != nil {
		return nil, fmt.Errorf("failed to unmarshal output: %w", err)
	}
	defer res.Body.Close()

	return &output, nil
}
