package synaps

type InitSessionResponse struct {
	SessionID string `json:"session_id"`
	Sandbox   bool   `json:"sandbox"`
}

type InitSessionParams struct {
	Metadata map[string]string `json:"metadata,omitempty"`
	Alias    string            `json:"alias,omitempty"`
}

type SessionDetailsResponse struct {
	App struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"app"`
	Session struct {
		ID      string `json:"id"`
		Alias   string `json:"alias"`
		Status  Status `json:"status"`
		Sandbox bool   `json:"sandbox"`
		Steps   []Step `json:"steps"`
	} `json:"session"`
}

type Step struct {
	ID     string   `json:"id"`
	Status Status   `json:"status"`
	Type   StepType `json:"type"`
}
