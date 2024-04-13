package via_cep

type Service struct{}

type ResponsePayload struct {
	Error   bool   `json:"erro"`
	Zipcode string `json:"cep"`
}
