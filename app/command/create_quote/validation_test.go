package create_quote

import (
	"context"
	"testing"

	"github.com/josemateuss/backend-challenge-frete-rapido/app/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUseCase_Validate(t *testing.T) {
	mockRepo := new(MockRepository)
	validateZipcodeService := new(ValidateZipcodeMockService)
	simulateQuoteService := new(SimulateQuoteMockService)
	uc := UseCase{
		createQuoteRepository:  mockRepo,
		validateZipcodeService: validateZipcodeService,
		simulateQuoteService:   simulateQuoteService,
	}

	t.Run("Test with valid input", func(t *testing.T) {
		input := Input{
			Recipient: Recipient{
				Address: Address{
					Zipcode: "73340030",
				},
			},
			Volumes: []Volume{
				{
					Category:      1,
					Amount:        1,
					UnitaryWeight: 1,
					Price:         1,
					Height:        1,
					Width:         1,
					Length:        1,
				},
			},
		}

		validateZipcodeService.
			On("Validate", mock.Anything, service.ValidateZipcodeInput{Zipcode: "73340030"}).
			Return(&service.ValidateZipcodeOutput{}, nil)

		err := uc.Validate(context.Background(), input)
		assert.Nil(t, err)
	})

	t.Run("Test with invalid input", func(t *testing.T) {
		input := Input{
			Recipient: Recipient{
				Address: Address{
					Zipcode: "",
				},
			},
			Volumes: []Volume{
				{
					Category:      0,
					Amount:        0,
					UnitaryWeight: 0,
					Price:         0,
					Height:        0,
					Width:         0,
					Length:        0,
				},
			},
		}

		err := uc.Validate(context.Background(), input)
		assert.NotNil(t, err)
	})
}
