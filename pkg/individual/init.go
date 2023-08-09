package individual

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/synaps.io/synaps-sdk-go/pkg/individual/models"
)

func (c *Client) InitSession(req InitSessionRequest) (sessionID InitSessionResponse, err error) {
	headers := map[string]string{"Api-Key": c.apiKey, "Content-Type": "application/json"}
	body, err := json.Marshal(req)
	if err != nil {
		return InitSessionResponse{}, fmt.Errorf("failed to marshal input: %s", err)
	}

	res, err := makeRequest[InitSessionResponse](c.httpClient, "POST", c.baseURL+"session/init", bytes.NewReader(body), headers)
	if err != nil {
		return InitSessionResponse{}, fmt.Errorf("init session request failed: %s", err)
	}
	return *res, nil
}
