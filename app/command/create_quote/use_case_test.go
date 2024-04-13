package create_quote

import (
	"context"
	"errors"
	"testing"

	"github.com/josemateuss/backend-challenge-frete-rapido/app/repository"
	"github.com/josemateuss/backend-challenge-frete-rapido/app/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

type MockService struct {
	mock.Mock
}

func (m *MockRepository) CreateQuote(ctx context.Context, input repository.CreateQuoteInput) (
	*repository.CreateQuoteOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*repository.CreateQuoteOutput), args.Error(1)
}

func (m *MockService) Simulate(ctx context.Context, input service.SimulateQuotesInput) (
	*service.SimulateQuotesOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*service.SimulateQuotesOutput), args.Error(1)
}

func TestUseCase_Execute(t *testing.T) {
	mockRepo := new(MockRepository)
	mockService := new(MockService)
	uc := UseCase{
		repository: mockRepo,
		service:    mockService,
	}

	ctx := context.Background()
	input := Input{}
	serviceOutput := &service.SimulateQuotesOutput{}
	repoOutput := &repository.CreateQuoteOutput{}

	mockService.On("Simulate", ctx, serviceSimulateQuoteInput(input)).Return(serviceOutput, nil)
	mockRepo.On("CreateQuote", ctx, repository.CreateQuoteInput{}).Return(repoOutput, nil)

	output, err := uc.Execute(ctx, input)

	assert.NoError(t, err)
	assert.Equal(t, Output{Quote: repoOutput.Quote}, output)

	mockService.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestUseCase_Execute_Error(t *testing.T) {
	mockRepo := new(MockRepository)
	mockService := new(MockService)
	uc := UseCase{
		repository: mockRepo,
		service:    mockService,
	}

	ctx := context.Background()
	input := Input{}

	mockService.On("Simulate", ctx, serviceSimulateQuoteInput(input)).
		Return(&service.SimulateQuotesOutput{}, errors.New("error"))

	_, err := uc.Execute(ctx, input)

	assert.Error(t, err)

	mockService.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestUseCase_New(t *testing.T) {
	type args struct {
		repository repository.CreateQuote
		service    service.SimulateQuote
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
				service:    new(MockService),
			},
			want: UseCase{
				repository: new(MockRepository),
				service:    new(MockService),
			},
			wantErr: assert.NoError,
		},
		{
			name: "Test with nil repository",
			args: args{
				repository: nil,
				service:    new(MockService),
			},
			want:    UseCase{},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.repository, tt.args.service)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
