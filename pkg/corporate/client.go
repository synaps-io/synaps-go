package corporate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	. "github.com/synaps.io/synaps-sdk-go/pkg/models"
)

type Client struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

type CorporateService interface {
	Init() (sessionID string, err error)
	Details(sessionID string) (SessionDetailsResponse, error)
	Overview(sessionID string) (SessionOverviewResponse, error)
}

func NewClient(apiKey string) *Client {
	return &Client{
		httpClient: http.DefaultClient,
		apiKey:     apiKey,
		baseURL:    "https://api.dev.synaps.run/v4/",
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

func (c *Client) Init() (sessionID string, err error) {
	res, err := makeRequest[InitSessionResponse](c.httpClient, "POST", c.baseURL+"session/init", nil, map[string]string{"Api-Key": c.apiKey})
	if err != nil {
		return "", fmt.Errorf("failed to make init session request: %w", err)
	}
	return res.SessionID, nil
}

func (c *Client) Details(sessionID string) (details SessionDetailsResponse, err error) {
	res, err := makeRequest[SessionDetailsResponse](c.httpClient, "GET", c.baseURL+"onboarding/details", nil, map[string]string{"Api-Key": c.apiKey, "session-id": sessionID})
	if err != nil {
		return SessionDetailsResponse{}, fmt.Errorf("failed to make session details request: %w", err)
	}
	return *res, nil
}

func (c *Client) Overview(sessionID string) (details SessionOverviewResponse, err error) {
	res, err := makeRequest[SessionOverviewResponse](c.httpClient, "GET", c.baseURL+"onboarding/overview", nil, map[string]string{"Api-Key": c.apiKey, "session-id": sessionID})
	if err != nil {
		return SessionOverviewResponse{}, fmt.Errorf("failed to make session details request: %w", err)
	}
	return *res, nil
}
