package synaps

import (
	"fmt"
)

func (c *Client) GetSessionDetails(sessionID string) (SessionDetailsResponse, error) {
	res, err := makeRequest[SessionDetailsResponse](c.httpClient, "GET", c.baseURL+"individual/session/"+sessionID, nil, map[string]string{"Api-Key": c.apiKey})
	if err != nil {
		return SessionDetailsResponse{}, fmt.Errorf("session details request failed: %s", err)
	}
	return *res, nil
}
