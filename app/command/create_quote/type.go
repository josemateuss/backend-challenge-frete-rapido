package create_quote

import (
	"github.com/josemateuss/backend-challenge-frete-rapido/app/repository"
	"github.com/josemateuss/backend-challenge-frete-rapido/app/service"
	"github.com/josemateuss/backend-challenge-frete-rapido/domain"
)

type UseCase struct {
	repository repository.CreateQuote
	service    service.SimulateQuote
}

type Input struct {
	Recipient Recipient `json:"recipient"`
	Volumes   []Volume  `json:"volumes"`
}

type Recipient struct {
	Address Address `json:"address"`
}

type Address struct {
	Zipcode string `json:"zipcode"`
}

type Volume struct {
	Category      uint    `json:"category"`
	Amount        uint    `json:"amount"`
	UnitaryWeight uint    `json:"unitary_weight"`
	Price         uint    `json:"price"`
	Sku           string  `json:"sku"`
	Height        float64 `json:"height"`
	Width         float64 `json:"width"`
	Length        float64 `json:"length"`
}

type Output struct {
	Quote *domain.Quote
}
