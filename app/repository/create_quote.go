package repository

import (
	"context"
	"time"

	"github.com/josemateuss/backend-challenge-frete-rapido/domain"
)

type CreateQuoteInput struct {
	Carrier   []Carrier `bson:"carrier"`
	CreatedAt time.Time `bson:"created_at"`
}

type Carrier struct {
	Name     string  `bson:"name"`
	Service  string  `bson:"service"`
	Deadline uint    `bson:"deadline"`
	Price    float64 `bson:"price"`
}

type CreateQuoteOutput struct {
	Quote *domain.Quote
}

type CreateQuote interface {
	CreateQuote(ctx context.Context, input CreateQuoteInput) (output *CreateQuoteOutput, err error)
}
