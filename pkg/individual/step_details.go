package individual

import (
	"fmt"
)

func (c *Client) GetStepDetails(sessionID string, stepID string) (any, error) {
	res, err := makeRequest[any](c.httpClient, "GET", c.baseURL+"individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.apiKey})
	if err != nil {
		return nil, fmt.Errorf("failed to make session details request: %w", err)
	}

	return *res, nil
}
