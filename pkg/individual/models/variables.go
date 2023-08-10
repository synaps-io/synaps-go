package models

type SynapsStep string

const (
	LivenessStep       SynapsStep = "LIVENESS"
	IDDocumentStep     SynapsStep = "ID_DOCUMENT"
	ProofOfAddressStep SynapsStep = "PROOF_OF_ADDRESS"
	EmailStep          SynapsStep = "EMAIL"
	PhoneStep          SynapsStep = "PHONE"
)

type SynapsStatus string

const (
	Rejected             SynapsStatus = "REJECTED"
	SubmissionRequired   SynapsStatus = "SUBMISSION_REQUIRED"
	ResubmissionRequired SynapsStatus = "RESUBMISSION_REQUIRED"
	Pending              SynapsStatus = "PENDING_VERIFICATION"
	Approved             SynapsStatus = "APPROVED"
)
