package read_quote_metrics

import "github.com/josemateuss/backend-challenge-frete-rapido/app/repository"

type UseCase struct {
	repository repository.ReadQuoteMetrics
}

type Input struct {
	LastQuotes uint `json:"last_quotes"`
}

type Output struct {
	ResultsPerCarrier      map[string]int     `json:"results_per_carrier"`
	TotalPricePerCarrier   map[string]float64 `json:"total_price_per_carrier"`
	AveragePricePerCarrier map[string]float64 `json:"average_price_per_carrier"`
	CheapestFreight        float64            `json:"cheapest_freight"`
	MostExpensiveFreight   float64            `json:"most_expensive_freight"`
}
