package mongodb

import (
	"context"
	"testing"

	"github.com/josemateuss/backend-challenge-frete-rapido/app/repository"
	"github.com/josemateuss/backend-challenge-frete-rapido/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockReadQuotesRepository struct {
	mock.Mock
}

func (m *MockReadQuotesRepository) ReadQuotes(ctx context.Context, input repository.ReadQuotesInput) (
	*repository.ReadQuotesOutput, error) {
	args := m.Called(ctx)
	return args.Get(0).(*repository.ReadQuotesOutput), args.Error(1)
}

func TestMongodb_ReadQuotes(t *testing.T) {
	mockRepo := new(MockReadQuotesRepository)
	ctx := context.Background()

	output := &repository.ReadQuotesOutput{
		Quotes: []domain.Quote{
			{
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
		},
	}

	mockRepo.On("ReadQuotes", ctx).Return(output, nil)

	_, err := mockRepo.ReadQuotes(ctx, repository.ReadQuotesInput{})

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
