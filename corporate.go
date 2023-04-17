package synaps

type CorporateService interface {
	// TODO
}

type corporate struct {
	apiKey string
}

func InitCorporate(apiKey string) *corporate {
	return &corporate{
		apiKey: apiKey,
	}
}
