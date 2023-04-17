package synaps

type SessionService interface {
	Init(appID string) (sessionID string, err error)
	Details(sessionID string) (SessionDetails, error)
	Overview(sessionID string) (SessionOverview, error)
}

type SessionDetails struct{}

type SessionOverview struct{}
