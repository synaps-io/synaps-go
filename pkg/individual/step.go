package synaps

type StepType string

const (
	LivenessStep       StepType = "LIVENESS"
	IDDocumentStep     StepType = "ID_DOCUMENT"
	ProofOfAddressStep StepType = "PROOF_OF_ADDRESS"
	EmailStep          StepType = "EMAIL"
	PhoneStep          StepType = "PHONE"
)

type ReasonCode string

const (
	ForgedRejectionReason              ReasonCode = "FORGED_REJECTION"
	DocumentHiddenReason               ReasonCode = "DOCUMENT_HIDDEN"
	BadEnvironmentReason               ReasonCode = "BAD_ENVIRONMENT"
	BlackWhitePictureReason            ReasonCode = "BLACK_WHITE_PICTURE"
	BadQualityReason                   ReasonCode = "BAD_QUALITY"
	DocumentComplianceReason           ReasonCode = "DOCUMENT_COMPLIANCE"
	IdentityDocumentExpiredReason      ReasonCode = "IDENTITY_DOCUMENT_EXPIRED"
	DocumentInvalidFrontSideReason     ReasonCode = "DOCUMENT_INVALID_FRONT_SIDE"
	DocumentInvalidBackSideReason      ReasonCode = "DOCUMENT_INVALID_BACK_SIDE"
	IdentityDocumentDobDateMinorReason ReasonCode = "IDENTITY_DOCUMENT_DOB_DATE_MINOR"
	RestrictedNationalityTypeReason    ReasonCode = "RESTRICTED_NATIONALITY_TYPE"
)

type StepReason struct {
	Code    ReasonCode
	Message string
}

type IDDocumentType string

const (
	Passport       IDDocumentType = "PASSPORT"
	NationalID     IDDocumentType = "NATIONAL_ID"
	DriverLicense  IDDocumentType = "DRIVER_LICENSE"
	ResidentPermit IDDocumentType = "RESIDENT_PERMIT"
)

type ProofOfAddressDocumentType string

const (
	ElectricityBill ProofOfAddressDocumentType = "ELECTRICITY_BILL"
	InternetBill    ProofOfAddressDocumentType = "INTERNET_BILL"
	LandlineBill    ProofOfAddressDocumentType = "LANDLINE_BILL"
	WaterBill       ProofOfAddressDocumentType = "WATER_BILL"
	GasBill         ProofOfAddressDocumentType = "GAS_BILL"
	BankStatement   ProofOfAddressDocumentType = "BANK_STATEMENT"
)

type PhoneMethod string

const (
	Sms  PhoneMethod = "sms"
	Call PhoneMethod = "call"
)

type File struct {
	URL  string `json:"url"`
	Type string `json:"type"`
	Size int    `json:"size"`
}

type (
	idDocumentData struct {
		Country string            `json:"country"`
		Type    IDDocumentType    `json:"type"`
		Fields  map[string]string `json:"fields"`
		Files   struct {
			Front File `json:"front"`
			Back  File `json:"back"`
			Face  File `json:"face"`
		} `json:"files"`
		OriginalFiles struct {
			Front File `json:"front"`
			Back  File `json:"back"`
		} `json:"original_files"`
	}
	proofOfAddressData struct {
		Country string                     `json:"country"`
		Type    ProofOfAddressDocumentType `json:"type"`
		Fields  map[string]string          `json:"fields"`
		Files   struct {
			Accomodation File `json:"accomodation"`
			Proof        File `json:"proof"`
		} `json:"files"`
		HostIDDocument idDocumentData `json:"host_id_document"`
		OriginalFiles  struct {
			Accomodation File `json:"accomodation"`
			Proof        File `json:"proof"`
		} `json:"original_files"`
	}
	phoneData struct {
		Phone struct {
			CallingCode string      `json:"calling_code"`
			Country     string      `json:"country"`
			Method      PhoneMethod `json:"method"`
			Number      string      `json:"number"`
		} `json:"phone"`
	}
	emailData struct {
		Email struct {
			Value string `json:"value"`
		} `json:"email"`
	}
	livenessData struct {
		Liveness struct {
			File File `json:"file"`
		} `json:"liveness"`
	}
)

type StepMetadata struct {
	IP               string `json:"ip"`
	UserAgent        string `json:"user_agent"`
	Platform         string `json:"platform"`
	BrowserName      string `json:"browser_name"`
	BrowserVersion   string `json:"browser_version"`
	Device           string `json:"device"`
	CustomerLanguage string `json:"customer_language"`
}

type PhoneStepDetailsResponse struct {
	ID       string       `json:"id"`
	Type     StepType     `json:"type"`
	Metadata StepMetadata `json:"metadata"`
	Status   Status       `json:"status"`
	Reason   StepReason   `json:"reason"`

	Verification phoneData `json:"verification"`
}

type EmailStepDetailsResponse struct {
	ID       string       `json:"id"`
	Type     StepType     `json:"type"`
	Metadata StepMetadata `json:"metadata"`
	Status   Status       `json:"status"`
	Reason   StepReason   `json:"reason"`

	Verification emailData `json:"verification"`
}

type ProofOfAddressStepDetailsResponse struct {
	ID       string       `json:"id"`
	Type     StepType     `json:"type"`
	Metadata StepMetadata `json:"metadata"`
	Status   Status       `json:"status"`
	Reason   StepReason   `json:"reason"`

	Document proofOfAddressData `json:"document"`
}

type IDStepDetailsResponse struct {
	ID       string       `json:"id"`
	Type     StepType     `json:"type"`
	Metadata StepMetadata `json:"metadata"`
	Status   Status       `json:"status"`
	Reason   StepReason   `json:"reason"`

	Document idDocumentData `json:"document"`
}

type LivenessStepDetailsResponse struct {
	ID       string       `json:"id"`
	Type     StepType     `json:"type"`
	Metadata StepMetadata `json:"metadata"`
	Status   Status       `json:"status"`
	Reason   StepReason   `json:"reason"`

	Verification livenessData `json:"verification"`
}
