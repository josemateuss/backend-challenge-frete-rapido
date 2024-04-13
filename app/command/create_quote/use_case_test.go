package create_quote

import (
	"context"
	"errors"
	"testing"

	"github.com/josemateuss/backend-challenge-frete-rapido/app/repository"
	"github.com/josemateuss/backend-challenge-frete-rapido/app/service"
	"github.com/josemateuss/backend-challenge-frete-rapido/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

type ValidateZipcodeMockService struct {
	mock.Mock
}

type SimulateQuoteMockService struct {
	mock.Mock
}

func (m *MockRepository) CreateQuote(ctx context.Context, input repository.CreateQuoteInput) (
	*repository.CreateQuoteOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*repository.CreateQuoteOutput), args.Error(1)
}

func (m *ValidateZipcodeMockService) Validate(ctx context.Context, input service.ValidateZipcodeInput) (
	*service.ValidateZipcodeOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*service.ValidateZipcodeOutput), args.Error(1)
}

func (m *SimulateQuoteMockService) Simulate(ctx context.Context, input service.SimulateQuotesInput) (
	*service.SimulateQuotesOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*service.SimulateQuotesOutput), args.Error(1)
}

func TestUseCase_Execute(t *testing.T) {
	mockRepo := new(MockRepository)
	validateZipcodeMockService := new(ValidateZipcodeMockService)
	simulateQuoteMockService := new(SimulateQuoteMockService)
	uc := UseCase{
		createQuoteRepository:  mockRepo,
		validateZipcodeService: validateZipcodeMockService,
		simulateQuoteService:   simulateQuoteMockService,
	}

	ctx := context.Background()
	input := Input{}
	simulateServiceOutput := &service.SimulateQuotesOutput{
		Carrier: []service.Carrier{
			{
				Name:     "Correios",
				Service:  "Sedex",
				Deadline: 1,
				Price:    80.0,
			},
		},
	}
	repoOutput := &repository.CreateQuoteOutput{
		Quote: &domain.Quote{
			Carrier: []domain.Carrier{
				{
					Name:     "Correios",
					Service:  "Sedex",
					Deadline: 1,
					Price:    80.0,
				},
			},
		},
	}

	simulateQuoteMockService.On("Simulate", ctx, serviceSimulateQuoteInput(input)).
		Return(simulateServiceOutput, nil)
	mockRepo.On("CreateQuote", ctx, repository.CreateQuoteInput{
		Carrier: []repository.Carrier{
			{
				Name:     "Correios",
				Service:  "Sedex",
				Deadline: 1,
				Price:    80.0,
			},
		},
	}).Return(repoOutput, nil)

	output, err := uc.Execute(ctx, input)

	assert.NoError(t, err)
	assert.Equal(t, Output{Quote: repoOutput.Quote}, output)

	simulateQuoteMockService.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestUseCase_Execute_Error(t *testing.T) {
	mockRepo := new(MockRepository)
	validateZipcodeService := new(ValidateZipcodeMockService)
	simulateQuoteService := new(SimulateQuoteMockService)
	uc := UseCase{
		createQuoteRepository:  mockRepo,
		validateZipcodeService: validateZipcodeService,
		simulateQuoteService:   simulateQuoteService,
	}

	ctx := context.Background()
	input := Input{}

	simulateQuoteService.On("Simulate", ctx, serviceSimulateQuoteInput(input)).
		Return(&service.SimulateQuotesOutput{}, errors.New("error"))

	_, err := uc.Execute(ctx, input)

	assert.Error(t, err)

	simulateQuoteService.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestUseCase_New(t *testing.T) {
	type args struct {
		createQuoteRepository  repository.CreateQuote
		validateZipcodeService service.ValidateZipCode
		simulateQuoteService   service.SimulateQuote
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
				createQuoteRepository:  new(MockRepository),
				validateZipcodeService: new(ValidateZipcodeMockService),
				simulateQuoteService:   new(SimulateQuoteMockService),
			},
			want: UseCase{
				createQuoteRepository:  new(MockRepository),
				validateZipcodeService: new(ValidateZipcodeMockService),
				simulateQuoteService:   new(SimulateQuoteMockService),
			},
			wantErr: assert.NoError,
		},
		{
			name: "Test with nil repository",
			args: args{
				createQuoteRepository:  nil,
				validateZipcodeService: new(ValidateZipcodeMockService),
				simulateQuoteService:   new(SimulateQuoteMockService),
			},
			want:    UseCase{},
			wantErr: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.createQuoteRepository, tt.args.validateZipcodeService, tt.args.simulateQuoteService)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
