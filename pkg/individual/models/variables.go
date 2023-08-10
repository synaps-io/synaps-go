package models

type SynapsStep string

const (
	Liveness       SynapsStep = "LIVENESS"
	IDDocument     SynapsStep = "ID_DOCUMENT"
	ProofOfAddress SynapsStep = "PROOF_OF_ADDRESS"
	Email          SynapsStep = "EMAIL"
	Phone          SynapsStep = "PHONE"
)

type SynapsStatus string

const (
	Rejected             SynapsStatus = "REJECTED"
	SubmissionRequired   SynapsStatus = "SUBMISSION_REQUIRED"
	ResubmissionRequired SynapsStatus = "RESUBMISSION_REQUIRED"
	Pending              SynapsStatus = "PENDING_VERIFICATION"
	Approved             SynapsStatus = "APPROVED"
)
