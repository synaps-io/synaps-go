package individual

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

func NewClient(apiKey string) *Client {
	return &Client{
		httpClient: http.DefaultClient,
		apiKey:     apiKey,
		baseURL:    "https://api.synaps.io/v4/",
	}
}

func makeRequest[T any](httpClient *http.Client, method string, path string, body io.Reader) (*T, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create init session request: %w", err)
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

func (c *Client) InitSession() (sessionID string, err error) {
	type InitSessionResponse struct {
		SessionID string `json:"session_id"`
	}

	res, err := makeRequest[InitSessionResponse](c.httpClient, "GET", c.baseURL+"session/init", nil)
	if err != nil {
		return "", fmt.Errorf("failed to make init session request: %w", err)
	}
	return res.SessionID, nil
}
