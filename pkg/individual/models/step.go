package models

type SynapsStep string

const (
	LivenessStep       SynapsStep = "LIVENESS"
	IDDocumentStep     SynapsStep = "ID_DOCUMENT"
	ProofOfAddressStep SynapsStep = "PROOF_OF_ADDRESS"
	EmailStep          SynapsStep = "EMAIL"
	PhoneStep          SynapsStep = "PHONE"
)

type ReasonCode string

const (
	ForgedRejection              ReasonCode = "FORGED_REJECTION"
	DocumentHidden               ReasonCode = "DOCUMENT_HIDDEN"
	BadEnvironment               ReasonCode = "BAD_ENVIRONMENT"
	BlackWhitePicture            ReasonCode = "BLACK_WHITE_PICTURE"
	BadQuality                   ReasonCode = "BAD_QUALITY"
	DocumentCompliance           ReasonCode = "DOCUMENT_COMPLIANCE"
	IdentityDocumentExpired      ReasonCode = "IDENTITY_DOCUMENT_EXPIRED"
	DocumentInvalidFrontSide     ReasonCode = "DOCUMENT_INVALID_FRONT_SIDE"
	DocumentInvalidBackSide      ReasonCode = "DOCUMENT_INVALID_BACK_SIDE"
	IdentityDocumentDobDateMinor ReasonCode = "IDENTITY_DOCUMENT_DOB_DATE_MINOR"
	RestrictedNationalityType    ReasonCode = "RESTRICTED_NATIONALITY_TYPE"
)

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

type StepReason struct {
	Code    ReasonCode
	Message string
}

type File struct {
	URL  string `json:"url"`
	Type string `json:"type"`
	Size int    `json:"size"`
}

type (
	IdDocumentData struct {
		Country string         `json:"country"`
		Type    IDDocumentType `json:"type"`
		Fields  struct {
			Firstname          string `json:"firstname"`
			Lastname           string `json:"lastname"`
			BirthDate          string `json:"birth_date"`
			DocumentExpiration string `json:"document_expiration"`
			DocumentNumber     string `json:"document_number"`
			Nationality        string `json:"nationality"`
		} `json:"fields"`
		Files struct {
			Front File `json:"front"`
			Back  File `json:"back"`
			Face  File `json:"face"`
		} `json:"files"`
		OriginalFiles struct {
			Front File `json:"front"`
			Back  File `json:"back"`
		} `json:"original_files"`
	}
	ProofOfAddressData struct {
		Country string `json:"country"`
		Type    string `json:"type"` // Either IDDocumentType or ProofOfAddressDocumentType
		Fields  struct {
			Address     string `json:"address"`
			City        string `json:"city"`
			IssuingDate string `json:"issuing_date"`
			Zipcode     string `json:"zipcode"`
		} `json:"fields"`
		Files struct {
			Accomodation File `json:"accomodation"`
			Proof        File `json:"proof"`
		} `json:"files"`
		HostIDDocument File `json:"host_id_document"`
		OriginalFiles  struct {
			Accomodation File `json:"accomodation"`
			Proof        File `json:"proof"`
		} `json:"original_files"`
	}
	PhoneData struct {
		Phone struct {
			CallingCode string      `json:"calling_code"`
			Country     string      `json:"country"`
			Method      PhoneMethod `json:"method"`
			Number      string      `json:"number"`
		} `json:"phone"`
	}
	EmailData struct {
		Email struct {
			Value string `json:"value"`
		} `json:"email"`
	}
	LivenessData struct {
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

type PhoneStepDetails struct {
	ID       string       `json:"id"`
	Type     SynapsStep   `json:"type"`
	Metadata StepMetadata `json:"metadata"`
	Status   SynapsStatus `json:"status"`
	Reason   StepReason   `json:"reason"`

	Verification PhoneData `json:"verification"`
}

type EmailStepDetails struct {
	ID       string       `json:"id"`
	Type     SynapsStep   `json:"type"`
	Metadata StepMetadata `json:"metadata"`
	Status   SynapsStatus `json:"status"`
	Reason   StepReason   `json:"reason"`

	Verification EmailData `json:"verification"`
}

type ProofOfAddressStepDetails struct {
	ID       string       `json:"id"`
	Type     SynapsStep   `json:"type"`
	Metadata StepMetadata `json:"metadata"`
	Status   SynapsStatus `json:"status"`
	Reason   StepReason   `json:"reason"`

	Document ProofOfAddressData `json:"document"`
}

type IDDocumentStepDetails struct {
	ID       string       `json:"id"`
	Type     SynapsStep   `json:"type"`
	Metadata StepMetadata `json:"metadata"`
	Status   SynapsStatus `json:"status"`
	Reason   StepReason   `json:"reason"`

	Document IdDocumentData `json:"document"`
}

type LivenessStepDetails struct {
	ID       string       `json:"id"`
	Type     SynapsStep   `json:"type"`
	Metadata StepMetadata `json:"metadata"`
	Status   SynapsStatus `json:"status"`
	Reason   StepReason   `json:"reason"`

	Timeline     []any
	Verification LivenessData `json:"verification"`
}
