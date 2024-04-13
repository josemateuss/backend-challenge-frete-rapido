package read_quote_metrics

import (
	"context"
	"fmt"
	"math"

	"github.com/josemateuss/backend-challenge-frete-rapido/app/repository"
)

func New(repository repository.ReadQuoteMetrics) (UseCase, error) {
	if repository == nil {
		return UseCase{}, fmt.Errorf("repository is required")
	}

	return UseCase{
		repository: repository,
	}, nil
}

func (uc UseCase) Execute(ctx context.Context, input Input) (output Output, err error) {
	readQuotesOutput, err := uc.repository.ReadQuotes(ctx, repository.ReadQuotesInput{
		LastQuotes: input.LastQuotes,
	})
	if err != nil {
		return output, err
	}

	output = Output{
		ResultsPerCarrier:      make(map[string]int),
		TotalPricePerCarrier:   make(map[string]float64),
		AveragePricePerCarrier: make(map[string]float64),
		CheapestFreight:        math.MaxFloat64,
		MostExpensiveFreight:   -math.MaxFloat64,
	}

	for _, quote := range readQuotesOutput.Quotes {
		for _, carrier := range quote.Carrier {
			output.ResultsPerCarrier[carrier.Name]++
			output.TotalPricePerCarrier[carrier.Name] =
				math.Round((output.TotalPricePerCarrier[carrier.Name]+carrier.Price)*100) / 100

			if carrier.Price < output.CheapestFreight {
				output.CheapestFreight = carrier.Price
			}
			if carrier.Price > output.MostExpensiveFreight {
				output.MostExpensiveFreight = carrier.Price
			}
		}
	}

	for carrier, totalPrice := range output.TotalPricePerCarrier {
		output.AveragePricePerCarrier[carrier] =
			math.Round((totalPrice/float64(output.ResultsPerCarrier[carrier]))*100) / 100
	}

	return output, nil
}
