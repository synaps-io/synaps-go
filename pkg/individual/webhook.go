package synaps

type WebhookPayload struct {
	Reason    string      `json:"reason"`
	Service   StepType  `json:"service"`
	SessionID string      `json:"session_id"`
	Status    WebhookEvent `json:"status"`
	StepID    string      `json:"step_id"`
}

type WebhookEvent string

const (
	EventRejected             WebhookEvent = "REJECTED"
	EventSubmissionRequired   WebhookEvent = "SUBMISSION_REQUIRED"
	EventResubmissionRequired WebhookEvent = "RESUBMISSION_REQUIRED"
	EventPending              WebhookEvent = "PENDING_VERIFICATION"
	EventApproved             WebhookEvent = "APPROVED"
	EventReset                WebhookEvent = "RESET"
)
