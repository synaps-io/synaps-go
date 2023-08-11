package synaps

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) InitSession(alias *string) (sessionID InitSessionResponse, err error) {
	req := InitSessionRequest{}
	if alias != nil {
		req.Alias = *alias
	}

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
