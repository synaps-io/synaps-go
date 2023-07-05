package models

type Document struct{}

type StepData[T StepType] struct {
	Timeline []T
	Data     T
}

type (
	StepLiveness struct {
		Liveness struct {
			File struct {
				URL  string `json:"url"`
				Type string `json:"type"`
				Size int    `json:"size"`
			} `json:"file"`
		} `json:"liveness"`
	}
	StepIdentity struct {
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
			Front struct {
				URL  string `json:"url"`
				Type string `json:"type"`
				Size int    `json:"size"`
			} `json:"front"`
			Back struct {
				URL  string `json:"url"`
				Type string `json:"type"`
				Size int    `json:"size"`
			} `json:"back"`
			Face struct {
				URL  string `json:"url"`
				Type string `json:"type"`
				Size int    `json:"size"`
			} `json:"face"`
		} `json:"files"`
		OriginalFiles struct {
			Front struct {
				URL  string `json:"url"`
				Type string `json:"type"`
				Size int    `json:"size"`
			} `json:"front"`
			Back struct {
				URL  string `json:"url"`
				Type string `json:"type"`
				Size int    `json:"size"`
			} `json:"back"`
		} `json:"original_files"`
	}
	StepResidency struct{}
	StepPhone     struct{}
	StepAml       struct{}
	StepEmail     struct{}
)

type StepType interface {
	StepIdentity | StepResidency | StepPhone | StepLiveness | StepAml | StepEmail
}

type StepDetailsResponse struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Metadata struct {
		IP               string `json:"ip"`
		UserAgent        string `json:"user_agent"`
		Platform         string `json:"platform"`
		BrowserName      string `json:"browser_name"`
		BrowserVersion   string `json:"browser_version"`
		Device           string `json:"device"`
		CustomerLanguage string `json:"customer_language"`
	} `json:"metadata"`
	Status string `json:"status"`

	// Step type relative fields

	Timeline     []any          `json:"timeline"`
	Verification map[string]any `json:"verification,omitempty"`
	Document     map[string]any `json:"document,omitempty"`
}
