package individual

import (
	"fmt"

	. "github.com/synaps.io/synaps-sdk-go/pkg/individual/models"
)

func (c *Client) InitSession() (sessionID string, err error) {
	res, err := makeRequest[InitSessionResponse](c.httpClient, "POST", c.baseURL+"session/init", nil, map[string]string{"Api-Key": c.apiKey})
	if err != nil {
		return "", fmt.Errorf("failed to make init session request: %w", err)
	}
	return res.SessionID, nil
}
