package individual

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	. "github.com/synaps.io/synaps-sdk-go/pkg/individual/models"
)

type Client struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

type IndividualClient interface {
	InitSession(alias *string, metadata map[string]string) (InitSessionResponse, error)
	GetSessionDetails(sessionID string) (SessionDetailsResponse, error)
	GetStepLivenessDetails(sessionID string, stepID string) (LivenessStepDetails, error)
	GetStepPhoneDetails(sessionID string, stepID string) (PhoneStepDetails, error)
	GetStepIDDocumentDetails(sessionID string, stepID string) (IDDocumentStepDetails, error)
	GetStepEmailDetails(sessionID string, stepID string) (EmailStepDetails, error)
	GetStepProofOfAddressDetails(sessionID string, stepID string) (ProofOfAddressStepDetails, error)
}

func NewClient(baseURL string, apiKey string) IndividualClient {
	return &Client{
		httpClient: http.DefaultClient,
		apiKey:     apiKey,
		baseURL:    baseURL,
	}
}

func NewClientFromEnv() IndividualClient {
	godotenv.Load()

	apiKey, ok := os.LookupEnv("SYNAPS_API_KEY")
	if !ok {
		log.Fatalf("Missing required SYNAPS_API_KEY env variable")
	}

	baseURL, ok := os.LookupEnv("SYNAPS_BASE_URL")
	if !ok {
		log.Fatalf("Missing required SYNAPS_BASE_URL env variable")
	}

	return NewClient(baseURL, apiKey)
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

	if !(res.StatusCode >= 200 && res.StatusCode < 300) {
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
