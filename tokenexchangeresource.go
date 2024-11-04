package adaptivecard

type TokenExchangeResource struct {
	ID         string `json:"id"`
	URI        string `json:"uri"`
	ProviderID string `json:"providerId"`
}

func NewTokenExchangeResource(id, uri, providerID string) *TokenExchangeResource {
	return &TokenExchangeResource{
		ID:         id,
		URI:        uri,
		ProviderID: providerID,
	}
}
