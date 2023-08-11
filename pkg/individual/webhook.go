package synaps

type WebhookPayload struct {
	Reason    string      `json:"reason"`
	Service   SynapsStep  `json:"service"`
	SessionID string      `json:"session_id"`
	Status    SynapsEvent `json:"status"`
	StepID    string      `json:"step_id"`
}

type SynapsEvent string

const (
	EventRejected             SynapsEvent = "REJECTED"
	EventSubmissionRequired   SynapsEvent = "SUBMISSION_REQUIRED"
	EventResubmissionRequired SynapsEvent = "RESUBMISSION_REQUIRED"
	EventPending              SynapsEvent = "PENDING_VERIFICATION"
	EventApproved             SynapsEvent = "APPROVED"
	EventReset                SynapsEvent = "RESET"
)
