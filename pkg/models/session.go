package synaps

type SessionOverviewResponse struct{}

type SessionDetailsResponse struct {
	Status string
}

type InitSessionResponse struct {
	SessionID string `json:"session_id"`
}
