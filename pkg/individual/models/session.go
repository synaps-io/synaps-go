package models

import "errors"

type InitSessionResponse struct {
	SessionID string `json:"session_id"`
	Sandbox   bool   `json:"sandbox"`
}

type InitSessionRequest struct {
	Metadata map[string]any `json:"metadata,omitempty"`
	Alias    string         `json:"alias,omitempty"`
}

type SessionDetailsResponse struct {
	App struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"app"`
	Session struct {
		ID      string `json:"id"`
		Alias   string `json:"alias"`
		Status  string `json:"status"`
		Sandbox bool   `json:"sandbox"`
		Steps   []Step `json:"steps"`
	} `json:"session"`
}

type Step struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Type   string `json:"type"`
}

func (details *SessionDetailsResponse) FindSessionStep(stepType string) (*Step, error) {
	for _, step := range details.Session.Steps {
		if step.Type == stepType {
			return &step, nil
		}
	}
	return nil, errors.New("step not found")
}
