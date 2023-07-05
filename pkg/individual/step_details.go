package individual

import (
	"encoding/json"
	"fmt"
	"reflect"

	. "github.com/synaps.io/synaps-sdk-go/pkg/individual/models"
)

func GetStepSpecificData[T StepType](response StepDetailsResponse) *StepData[T] {
	stepData := &StepData[T]{}

	bytes, _ := json.Marshal(response.Timeline)
	if err := json.Unmarshal(bytes, &stepData.Timeline); err != nil {
		return nil
	}

	if reflect.TypeOf(*stepData) == reflect.TypeOf(StepData[StepLiveness]{}) {
		bytes, _ = json.Marshal(response.Verification)
		if err := json.Unmarshal(bytes, &stepData.Data); err != nil {
			return nil
		}
	}

	if reflect.TypeOf(*stepData) == reflect.TypeOf(StepData[StepIdentity]{}) {
		bytes, _ = json.Marshal(response.Document)
		if err := json.Unmarshal(bytes, &stepData.Data); err != nil {
			return nil
		}
	}

	return stepData
}

func (c *Client) StepDetails(sessionID string, stepID string) (StepDetailsResponse, error) {
	res, err := makeRequest[StepDetailsResponse](c.httpClient, "GET", c.baseURL+"individual/session/"+sessionID+"/step/"+stepID, nil, map[string]string{"Api-Key": c.apiKey})
	if err != nil {
		return StepDetailsResponse{}, fmt.Errorf("failed to make session details request: %w", err)
	}

	return *res, nil
}
