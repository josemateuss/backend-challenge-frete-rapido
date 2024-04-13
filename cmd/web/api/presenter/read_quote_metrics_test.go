package presenter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadQuoteMetrics_Present(t *testing.T) {
	presenter := NewReadQuoteMetricsPresenter()

	t.Run("Test Read Quote Metrics Presenter", func(t *testing.T) {
		input := ReadQuoteMetricsPresentInput{
			ResultsPerCarrier:      map[string]int{"Correios": 6, "Loggi": 2},
			TotalPricePerCarrier:   map[string]float64{"Correios": 469.57, "Loggi": 173.6},
			AveragePricePerCarrier: map[string]float64{"Correios": 78.27, "Loggi": 88.3},
			CheapestFreight:        70.0,
			MostExpensiveFreight:   110.5,
		}

		expected := ReadQuoteMetricsPresentOutput{
			ResultsPerCarrier:      map[string]int{"Correios": 6, "Loggi": 2},
			TotalPricePerCarrier:   map[string]float64{"Correios": 469.57, "Loggi": 173.6},
			AveragePricePerCarrier: map[string]float64{"Correios": 78.27, "Loggi": 88.3},
			CheapestFreight:        70.0,
			MostExpensiveFreight:   110.5,
		}

		output := presenter.Present(input)
		assert.Equal(t, expected, output)
	})
}
