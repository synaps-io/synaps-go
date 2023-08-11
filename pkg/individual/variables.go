package synaps

type SynapsStatus string

const (
	StatusRejected             SynapsStatus = "REJECTED"
	StatusSubmissionRequired   SynapsStatus = "SUBMISSION_REQUIRED"
	StatusResubmissionRequired SynapsStatus = "RESUBMISSION_REQUIRED"
	StatusPending              SynapsStatus = "PENDING_VERIFICATION"
	StatusApproved             SynapsStatus = "APPROVED"
)
