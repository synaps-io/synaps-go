package synaps

type Instance struct {
	Individual SessionService
	Corporate  CorporateService
}

func InitSDK(apiKey string) *Instance {
	return &Instance{
		Individual: initIndividual(apiKey),
		Corporate:  initCorporate(apiKey),
	}
}

type individual struct {
	apiKey string
}

func (individual *individual) Init(appID string) (session string, err error) {
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

type corporate struct {
	apiKey string
}

func initCorporate(apiKey string) *corporate {
	return &corporate{
		apiKey: apiKey,
	}
}
