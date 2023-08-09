package individual

import (
	"fmt"

	. "github.com/synaps.io/synaps-sdk-go/pkg/individual/models"
)

func (c *Client) GetStepLivenessDetails(sessionID string, stepID string) (LivenessStepDetails, error) {
	res, err := makeRequest[LivenessStepDetails](c.httpClient, "GET", c.baseURL+"individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.apiKey})
	if err != nil {
		return LivenessStepDetails{}, fmt.Errorf("get liveness step details request failed: %s", err)
	}

	return *res, nil
}

func (c *Client) GetStepPhoneDetails(sessionID string, stepID string) (PhoneStepDetails, error) {
	res, err := makeRequest[PhoneStepDetails](c.httpClient, "GET", c.baseURL+"individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.apiKey})
	if err != nil {
		return PhoneStepDetails{}, fmt.Errorf("get phone step details request failed: %s", err)
	}

	return *res, nil
}

func (c *Client) GetStepIDDocumentDetails(sessionID string, stepID string) (IDDocumentStepDetails, error) {
	res, err := makeRequest[IDDocumentStepDetails](c.httpClient, "GET", c.baseURL+"individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.apiKey})
	if err != nil {
		return IDDocumentStepDetails{}, fmt.Errorf("get id document step details request failed: %s", err)
	}

	return *res, nil
}

func (c *Client) GetStepEmailDetails(sessionID string, stepID string) (EmailStepDetails, error) {
	res, err := makeRequest[EmailStepDetails](c.httpClient, "GET", c.baseURL+"individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.apiKey})
	if err != nil {
		return EmailStepDetails{}, fmt.Errorf("get email step details request failed: %s", err)
	}

	return *res, nil
}

func (c *Client) GetStepProofOfAddressDetails(sessionID string, stepID string) (ProofOfAddressStepDetails, error) {
	res, err := makeRequest[ProofOfAddressStepDetails](c.httpClient, "GET", c.baseURL+"individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.apiKey})
	if err != nil {
		return ProofOfAddressStepDetails{}, fmt.Errorf("get proof of address step details request failed: %s", err)
	}

	return *res, nil
}
