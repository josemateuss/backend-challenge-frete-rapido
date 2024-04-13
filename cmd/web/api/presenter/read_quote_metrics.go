package presenter

type ReadQuoteMetrics struct{}

type ReadQuoteMetricsPresentInput struct {
	ResultsPerCarrier      map[string]int     `json:"results_per_carrier"`
	TotalPricePerCarrier   map[string]float64 `json:"total_price_per_carrier"`
	AveragePricePerCarrier map[string]float64 `json:"average_price_per_carrier"`
	CheapestFreight        float64            `json:"cheapest_freight"`
	MostExpensiveFreight   float64            `json:"most_expensive_freight"`
}

type ReadQuoteMetricsPresentOutput struct {
	ResultsPerCarrier      map[string]int     `json:"results_per_carrier"`
	TotalPricePerCarrier   map[string]float64 `json:"total_price_per_carrier"`
	AveragePricePerCarrier map[string]float64 `json:"average_price_per_carrier"`
	CheapestFreight        float64            `json:"cheapest_freight"`
	MostExpensiveFreight   float64            `json:"most_expensive_freight"`
}

func NewReadQuoteMetricsPresenter() ReadQuoteMetrics {
	return ReadQuoteMetrics{}
}

func (ReadQuoteMetrics) Present(input ReadQuoteMetricsPresentInput) (output ReadQuoteMetricsPresentOutput) {
	return ReadQuoteMetricsPresentOutput{
		ResultsPerCarrier:      input.ResultsPerCarrier,
		TotalPricePerCarrier:   input.TotalPricePerCarrier,
		AveragePricePerCarrier: input.AveragePricePerCarrier,
		CheapestFreight:        input.CheapestFreight,
		MostExpensiveFreight:   input.MostExpensiveFreight,
	}
}
