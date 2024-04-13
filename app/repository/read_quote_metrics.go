package repository

import (
	"context"

	"github.com/josemateuss/backend-challenge-frete-rapido/domain"
)

type ReadQuotesInput struct {
	LastQuotes uint
}

type ReadQuotesOutput struct {
	Quotes []domain.Quote
}

type ReadQuoteMetrics interface {
	ReadQuotes(ctx context.Context, input ReadQuotesInput) (output *ReadQuotesOutput, err error)
}
