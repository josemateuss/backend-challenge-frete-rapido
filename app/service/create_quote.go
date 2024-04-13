package service

import "context"

type SimulateQuotesInput struct {
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

type SimulateQuotesOutput struct {
	Carrier []Carrier `json:"carrier"`
}

type Carrier struct {
	Name     string  `json:"name"`
	Service  string  `json:"service"`
	Deadline uint    `json:"deadline"`
	Price    float64 `json:"price"`
}

type SimulateQuote interface {
	Simulate(ctx context.Context, input SimulateQuotesInput) (output *SimulateQuotesOutput, err error)
}
