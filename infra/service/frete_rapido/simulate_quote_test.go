package frete_rapido

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/josemateuss/backend-challenge-frete-rapido/app/service"
	"github.com/stretchr/testify/assert"
)

func TestService_Simulate(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	ctx := context.Background()
	input := service.SimulateQuotesInput{
		Recipient: service.Recipient{
			Address: service.Address{
				Zipcode: "73340030",
			},
		},
		Volumes: []service.Volume{
			{
				Category:      7,
				Amount:        1,
				UnitaryWeight: 5,
				Price:         349,
				Sku:           "test-abc-123",
				Height:        0.2,
				Width:         0.2,
				Length:        0.2,
			},
			{
				Category:      7,
				Amount:        3,
				UnitaryWeight: 4,
				Price:         123,
				Sku:           "test-abc-456",
				Height:        0.4,
				Width:         0.4,
				Length:        0.4,
			},
		},
	}
	expectedOutput := &service.SimulateQuotesOutput{
		Carrier: []service.Carrier{
			{
				Name:     "Correios",
				Service:  "Sedex",
				Deadline: 1,
				Price:    60.0,
			},
			{
				Name:     "Correios",
				Service:  "Pac",
				Deadline: 4,
				Price:    30.0,
			},
		},
	}

	responsePayload := ResponsePayload{
		Dispatchers: []ResponseDispatcher{
			{
				Offers: []Offer{
					{
						Carrier: Carrier{
							Name: "Correios",
						},
						Service: "Sedex",
						DeliveryTime: DeliveryTime{
							Days: 1,
						},
						FinalPrice: 60.0,
					},
					{
						Carrier: Carrier{
							Name: "Correios",
						},
						Service: "Pac",
						DeliveryTime: DeliveryTime{
							Days: 4,
						},
						FinalPrice: 30.0,
					},
				},
			},
		},
	}

	responder, _ := httpmock.NewJsonResponder(http.StatusOK, responsePayload)
	httpmock.RegisterResponder(http.MethodPost, FreteRapidoSimulateURL, responder)

	s := Service{}
	output, err := s.Simulate(ctx, input)

	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, output)
}
