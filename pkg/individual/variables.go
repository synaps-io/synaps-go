package individual

const (
	Liveness       = "LIVENESS"
	IDDocument     = "ID_DOCUMENT"
	ProofOfAddress = "PROOF_OF_ADDRESS"
	Email          = "EMAIL"
	Phone          = "PHONE"
)

type SynapsStatus string

const (
	Rejected             SynapsStatus = "REJECTED"
	SubmissionRequired   SynapsStatus = "SUBMISSION_REQUIRED"
	ResubmissionRequired SynapsStatus = "RESUBMISSION_REQUIRED"
	Pending              SynapsStatus = "PENDING_VERIFICATION"
	Approved             SynapsStatus = "APPROVED"
)
