package presenter

import (
	"testing"

	"github.com/josemateuss/backend-challenge-frete-rapido/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreateQuote_Present(t *testing.T) {
	presenter := NewCreateQuotePresenter()

	t.Run("Test Create Quote Presenter", func(t *testing.T) {
		input := CreateQuotePresentInput{
			Quote: &domain.Quote{
				Carrier: []domain.Carrier{
					{
						Name:     "Correios",
						Service:  "PAC",
						Deadline: 7,
						Price:    60.0,
					},
					{
						Name:     "Correios",
						Service:  "Sedex",
						Deadline: 2,
						Price:    100.0,
					},
				},
			},
		}

		expected := CreateQuotePresentOutput{
			Carrier: []Carrier{
				{
					Name:     "Correios",
					Service:  "PAC",
					Deadline: 7,
					Price:    60.0,
				},
				{
					Name:     "Correios",
					Service:  "Sedex",
					Deadline: 2,
					Price:    100.0,
				},
			},
		}

		output := presenter.Present(input)
		assert.Equal(t, expected, output)
	})
}
