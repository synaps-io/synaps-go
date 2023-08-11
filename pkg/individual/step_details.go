package individual

import (
	"fmt"

	. "github.com/synaps.io/synaps-sdk-go/pkg/individual/models"
)

func (c *Client) GetStepLivenessDetails(sessionID string, stepID string) (LivenessStepDetailsResponse, error) {
	res, err := makeRequest[LivenessStepDetailsResponse](c.httpClient, "GET", c.baseURL+"individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.apiKey})
	if err != nil {
		return LivenessStepDetailsResponse{}, fmt.Errorf("get liveness step details request failed: %s", err)
	}

	return *res, nil
}

func (c *Client) GetStepPhoneDetails(sessionID string, stepID string) (PhoneStepDetailsResponse, error) {
	res, err := makeRequest[PhoneStepDetailsResponse](c.httpClient, "GET", c.baseURL+"individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.apiKey})
	if err != nil {
		return PhoneStepDetailsResponse{}, fmt.Errorf("get phone step details request failed: %s", err)
	}

	return *res, nil
}

func (c *Client) GetStepIDDocumentDetails(sessionID string, stepID string) (IDDocumentStepDetailsResponse, error) {
	res, err := makeRequest[IDDocumentStepDetailsResponse](c.httpClient, "GET", c.baseURL+"individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.apiKey})
	if err != nil {
		return IDDocumentStepDetailsResponse{}, fmt.Errorf("get id document step details request failed: %s", err)
	}

	return *res, nil
}

func (c *Client) GetStepEmailDetails(sessionID string, stepID string) (EmailStepDetailsResponse, error) {
	res, err := makeRequest[EmailStepDetailsResponse](c.httpClient, "GET", c.baseURL+"individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.apiKey})
	if err != nil {
		return EmailStepDetailsResponse{}, fmt.Errorf("get email step details request failed: %s", err)
	}

	return *res, nil
}

func (c *Client) GetStepProofOfAddressDetails(sessionID string, stepID string) (ProofOfAddressStepDetailsResponse, error) {
	res, err := makeRequest[ProofOfAddressStepDetailsResponse](c.httpClient, "GET", c.baseURL+"individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.apiKey})
	if err != nil {
		return ProofOfAddressStepDetailsResponse{}, fmt.Errorf("get proof of address step details request failed: %s", err)
	}

	return *res, nil
}
