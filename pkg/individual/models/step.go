package models

type File struct {
	URL  string `json:"url"`
	Type string `json:"type"`
	Size int    `json:"size"`
}

type (
	idDocumentData struct {
		Country string `json:"country"`
		Type    string `json:"type"`
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
	proofOfAddressData struct {
		Country string `json:"country"`
		Type    string `json:"type"`
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
	phoneData struct {
		Phone struct {
			CallingCode string `json:"calling_code"`
			Country     string `json:"country"`
			Method      string `json:"method"`
			Number      string `json:"number"`
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

type Metadata struct {
	IP               string `json:"ip"`
	UserAgent        string `json:"user_agent"`
	Platform         string `json:"platform"`
	BrowserName      string `json:"browser_name"`
	BrowserVersion   string `json:"browser_version"`
	Device           string `json:"device"`
	CustomerLanguage string `json:"customer_language"`
}

type PhoneStepDetails struct {
	ID       string   `json:"id"`
	Type     string   `json:"type"`
	Metadata Metadata `json:"metadata"`
	Status   string   `json:"status"`
	Reason   string   `json:"reason"`

	Timeline     []any
	Verification phoneData `json:"verification"`
}

type EmailStepDetails struct {
	ID       string   `json:"id"`
	Type     string   `json:"type"`
	Metadata Metadata `json:"metadata"`
	Status   string   `json:"status"`
	Reason   string   `json:"reason"`

	Timeline     []any
	Verification emailData `json:"verification"`
}

type ProofOfAddressStepDetails struct {
	ID       string   `json:"id"`
	Type     string   `json:"type"`
	Metadata Metadata `json:"metadata"`
	Status   string   `json:"status"`
	Reason   string   `json:"reason"`

	Timeline []any
	Document proofOfAddressData `json:"document"`
}

type IDDocumentStepDetails struct {
	ID       string   `json:"id"`
	Type     string   `json:"type"`
	Metadata Metadata `json:"metadata"`
	Status   string   `json:"status"`
	Reason   string   `json:"reason"`

	Timeline []any
	Document idDocumentData `json:"document"`
}

type LivenessStepDetails struct {
	ID       string   `json:"id"`
	Type     string   `json:"type"`
	Metadata Metadata `json:"metadata"`
	Status   string   `json:"status"`
	Reason   string   `json:"reason"`

	Timeline     []any
	Verification livenessData `json:"verification"`
}
