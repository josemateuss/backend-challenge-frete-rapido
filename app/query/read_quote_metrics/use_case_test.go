package read_quote_metrics

import (
	"context"
	"testing"

	"github.com/josemateuss/backend-challenge-frete-rapido/app/repository"
	"github.com/josemateuss/backend-challenge-frete-rapido/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) ReadQuotes(ctx context.Context, input repository.ReadQuotesInput) (
	output *repository.ReadQuotesOutput, err error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*repository.ReadQuotesOutput), args.Error(1)
}

func TestUseCase_Execute(t *testing.T) {
	mockRepo := new(MockRepository)
	uc := UseCase{
		repository: mockRepo,
	}

	ctx := context.Background()
	lastQuotes := uint(5)
	input := Input{LastQuotes: lastQuotes}
	readQuotesOutput := &repository.ReadQuotesOutput{
		Quotes: []domain.Quote{
			{
				Carrier: []domain.Carrier{
					{
						Name:     "Correios",
						Service:  "Sedex",
						Deadline: 1,
						Price:    55.9,
					},
					{
						Name:     "Correios",
						Service:  "Pac",
						Deadline: 4,
						Price:    40.0,
					},
				},
			},
		},
	}

	mockRepo.On("ReadQuotes", ctx, repository.ReadQuotesInput{LastQuotes: input.LastQuotes}).Return(readQuotesOutput, nil)

	output, err := uc.Execute(ctx, input)

	assert.NoError(t, err)
	assert.Equal(t, Output{
		ResultsPerCarrier:      map[string]int{"Correios": 2},
		TotalPricePerCarrier:   map[string]float64{"Correios": 95.9},
		AveragePricePerCarrier: map[string]float64{"Correios": 47.95},
		CheapestFreight:        40.0,
		MostExpensiveFreight:   55.9,
	}, output)

	mockRepo.AssertExpectations(t)
}

func TestUseCase_New(t *testing.T) {
	type args struct {
		repository repository.ReadQuoteMetrics
	}
	tests := []struct {
		name    string
		args    args
		want    UseCase
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Test with valid repository and service",
			args: args{
				repository: new(MockRepository),
			},
			want: UseCase{
				repository: new(MockRepository),
			},
			wantErr: assert.NoError,
		},
		{
			name: "Test with nil repository",
			args: args{
				repository: nil,
			},
			want:    UseCase{},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.repository)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
