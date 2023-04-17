package synaps

type Individual interface {
	InitIndividual() (sessionID string, err error)
	Details(sessionID string) (SessionDetails, error)
	Overview(sessionID string) (SessionOverview, error)
}

type SessionDetails struct {
	Status string
}

type SessionOverview struct{}

type individual struct {
	apiKey string
}

func NewIndividualSDK(apiKey string) Individual {
	return &individual{apiKey: apiKey}
}

func (individual *individual) InitIndividual() (session string, err error) {
	panic("not implemented") // TODO: Implement
}

func (individual *individual) Details(sessionID string) (SessionDetails, error) {
	panic("not implemented") // TODO: Implement
}

func (individual *individual) Overview(sessionID string) (SessionOverview, error) {
	panic("not implemented") // TODO: Implement
}

func initIndividual(apiKey string) *individual {
	return &individual{
		apiKey: apiKey,
	}
}
