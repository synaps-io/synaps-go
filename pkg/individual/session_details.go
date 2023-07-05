package individual

import (
	"fmt"

	. "github.com/synaps.io/synaps-sdk-go/pkg/individual/models"
)

func (c *Client) GetSessionDetails(sessionID string) (SessionDetailsResponse, error) {
	res, err := makeRequest[SessionDetailsResponse](c.httpClient, "GET", c.baseURL+"individual/session/"+sessionID, nil, map[string]string{"Api-Key": c.apiKey})
	if err != nil {
		return SessionDetailsResponse{}, fmt.Errorf("failed to make session details request: %w", err)
	}
	return *res, nil
}
