package synaps

type Status string

const (
	StatusRejected             Status = "REJECTED"
	StatusSubmissionRequired   Status = "SUBMISSION_REQUIRED"
	StatusResubmissionRequired Status = "RESUBMISSION_REQUIRED"
	StatusPending              Status = "PENDING_VERIFICATION"
	StatusApproved             Status = "APPROVED"
)
