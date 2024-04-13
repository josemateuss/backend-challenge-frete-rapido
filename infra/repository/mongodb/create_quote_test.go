package mongodb

import (
	"context"
	"testing"

	"github.com/josemateuss/backend-challenge-frete-rapido/app/repository"
	"github.com/josemateuss/backend-challenge-frete-rapido/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCreateQuoteRepository struct {
	mock.Mock
}

func (m *MockCreateQuoteRepository) CreateQuote(ctx context.Context, input repository.CreateQuoteInput) (
	*repository.CreateQuoteOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*repository.CreateQuoteOutput), args.Error(1)
}

func TestMongodb_CreateQuote(t *testing.T) {
	mockRepo := new(MockCreateQuoteRepository)

	ctx := context.Background()
	input := repository.CreateQuoteInput{
		Carrier: []repository.Carrier{
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

	output := &repository.CreateQuoteOutput{
		Quote: &domain.Quote{
			Carrier: []domain.Carrier{
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
		},
	}

	mockRepo.On("CreateQuote", ctx, input).Return(output, nil)

	_, err := mockRepo.CreateQuote(ctx, input)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
