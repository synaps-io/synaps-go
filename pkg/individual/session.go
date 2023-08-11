package synaps

import "errors"

type InitSessionResponse struct {
	SessionID string `json:"session_id"`
	Sandbox   bool   `json:"sandbox"`
}

type InitSessionRequest struct {
	Metadata map[string]string `json:"metadata,omitempty"`
	Alias    string            `json:"alias,omitempty"`
}

type SessionDetailsResponse struct {
	App struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"app"`
	Session struct {
		ID      string       `json:"id"`
		Alias   string       `json:"alias"`
		Status  SynapsStatus `json:"status"`
		Sandbox bool         `json:"sandbox"`
		Steps   []Step       `json:"steps"`
	} `json:"session"`
}

type Step struct {
	ID     string       `json:"id"`
	Status SynapsStatus `json:"status"`
	Type   SynapsStep   `json:"type"`
}

func (details *SessionDetailsResponse) FindSessionStep(stepType SynapsStep) (*Step, error) {
	for _, step := range details.Session.Steps {
		if step.Type == stepType {
			return &step, nil
		}
	}
	return nil, errors.New("step not found")
}
