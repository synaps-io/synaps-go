package synaps

import (
	"fmt"
)

func (c *Client) GetLivenessStepDetails(sessionID string, stepID string) (LivenessStepDetailsResponse, error) {
	res, err := makeRequest[LivenessStepDetailsResponse](c.httpClient, "GET", c.baseURL+"/individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.apiKey})
	if err != nil {
		return LivenessStepDetailsResponse{}, fmt.Errorf("get liveness step details request failed: %s", err)
	}

	return *res, nil
}

func (c *Client) GetPhoneStepDetails(sessionID string, stepID string) (PhoneStepDetailsResponse, error) {
	res, err := makeRequest[PhoneStepDetailsResponse](c.httpClient, "GET", c.baseURL+"/individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.apiKey})
	if err != nil {
		return PhoneStepDetailsResponse{}, fmt.Errorf("get phone step details request failed: %s", err)
	}

	return *res, nil
}

func (c *Client) GetIDStepDetails(sessionID string, stepID string) (IDStepDetailsResponse, error) {
	res, err := makeRequest[IDStepDetailsResponse](c.httpClient, "GET", c.baseURL+"/individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.apiKey})
	if err != nil {
		return IDStepDetailsResponse{}, fmt.Errorf("get id document step details request failed: %s", err)
	}

	return *res, nil
}

func (c *Client) GetIDStepsDetails(sessionID string) ([]IDStepDetailsResponse, error) {
	sessionDetails, err := c.GetSessionDetails(sessionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get details for session[%s]: %s", sessionID, err)
	}

	var amlSteps []IDStepDetailsResponse

	for _, step := range sessionDetails.Session.Steps {
		if step.Type == IDDocumentStep {
			res, err := makeRequest[IDStepDetailsResponse](c.httpClient, "GET", c.baseURL+"/individual/session/"+sessionID+"/step/"+step.ID, nil, map[string]string{"Api-Key": c.apiKey})
			if err != nil {
				return nil, fmt.Errorf("get AML step details request failed: %s", err)
			}
			amlSteps = append(amlSteps, *res)
		}
	}

	return amlSteps, nil
}

func (c *Client) GetEmailStepDetails(sessionID string, stepID string) (EmailStepDetailsResponse, error) {
	res, err := makeRequest[EmailStepDetailsResponse](c.httpClient, "GET", c.baseURL+"/individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.apiKey})
	if err != nil {
		return EmailStepDetailsResponse{}, fmt.Errorf("get email step details request failed: %s", err)
	}

	return *res, nil
}

func (c *Client) GetProofOfAddressStepDetails(sessionID string, stepID string) (ProofOfAddressStepDetailsResponse, error) {
	res, err := makeRequest[ProofOfAddressStepDetailsResponse](c.httpClient, "GET", c.baseURL+"/individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.apiKey})
	if err != nil {
		return ProofOfAddressStepDetailsResponse{}, fmt.Errorf("get proof of address step details request failed: %s", err)
	}

	return *res, nil
}

func (c *Client) GetAMLStepDetails(sessionID string, stepID string) (AMLStepDetailsResponse, error) {
	res, err := makeRequest[AMLStepDetailsResponse](c.httpClient, "GET", c.baseURL+"/individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.apiKey})
	if err != nil {
		return AMLStepDetailsResponse{}, fmt.Errorf("get AML step details request failed: %s", err)
	}

	return *res, nil
}

func (c *Client) GetAMLStepsDetails(sessionID string) ([]AMLStepDetailsResponse, error) {
	sessionDetails, err := c.GetSessionDetails(sessionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get details for session[%s]: %s", sessionID, err)
	}

	var amlSteps []AMLStepDetailsResponse

	for _, step := range sessionDetails.Session.Steps {
		if step.Type == AMLStep {
			res, err := makeRequest[AMLStepDetailsResponse](c.httpClient, "GET", c.baseURL+"/individual/session/"+sessionID+"/step/"+step.ID, nil, map[string]string{"Api-Key": c.apiKey})
			if err != nil {
				return nil, fmt.Errorf("get AML step details request failed: %s", err)
			}
			amlSteps = append(amlSteps, *res)
		}
	}

	return amlSteps, nil
}
