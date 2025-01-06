package models

type Address struct {
	Cep          string `json:"cep"`
	State        string `json:"state,omitempty"`        // Para BrasilAPI
	City         string `json:"city,omitempty"`         // Para BrasilAPI
	Neighborhood string `json:"neighborhood,omitempty"` // Para BrasilAPI e ViaCEP
	Street       string `json:"street,omitempty"`       // Para BrasilAPI
	Logradouro   string `json:"logradouro,omitempty"`   // Para ViaCEP
	Complemento  string `json:"complemento,omitempty"`  // Para ViaCEP
	Localidade   string `json:"localidade,omitempty"`   // Para ViaCEP
	Uf           string `json:"uf,omitempty"`           // Para ViaCEP
	Service      string `json:"service,omitempty"`      // Indica a origem na BrasilAPI
	ApiSource    string `json:"api_source"`             // Indica a API utilizada (BrasilAPI ou ViaCEP)
}
