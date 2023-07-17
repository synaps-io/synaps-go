package models

type WebhookPayload struct {
	Reason    string `json:"reason"`
	Service   string `json:"service"`
	SessionID string `json:"session_id"`
	Status    string `json:"status"`
	StepID    string `json:"step_id"`
}
