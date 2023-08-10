package models

type SynapsStatus string

const (
	Rejected             SynapsStatus = "REJECTED"
	SubmissionRequired   SynapsStatus = "SUBMISSION_REQUIRED"
	ResubmissionRequired SynapsStatus = "RESUBMISSION_REQUIRED"
	Pending              SynapsStatus = "PENDING_VERIFICATION"
	Approved             SynapsStatus = "APPROVED"
)
