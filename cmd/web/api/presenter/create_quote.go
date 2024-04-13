package presenter

import "github.com/josemateuss/backend-challenge-frete-rapido/domain"

type CreateQuote struct{}

type CreateQuotePresentInput struct {
	Quote *domain.Quote
}

type CreateQuotePresentOutput struct {
	Carrier []Carrier `json:"carrier"`
}

type Carrier struct {
	Name     string  `json:"name"`
	Service  string  `json:"service"`
	Deadline uint    `json:"deadline"`
	Price    float64 `json:"price"`
}

func NewCreateQuotePresenter() CreateQuote {
	return CreateQuote{}
}

func (CreateQuote) Present(input CreateQuotePresentInput) (output CreateQuotePresentOutput) {
	carriers := make([]Carrier, len(input.Quote.Carrier))
	for i, carrier := range input.Quote.Carrier {
		carriers[i] = Carrier{
			Name:     carrier.Name,
			Service:  carrier.Service,
			Deadline: carrier.Deadline,
			Price:    carrier.Price,
		}
	}

	output = CreateQuotePresentOutput{
		Carrier: carriers,
	}

	return output
}
