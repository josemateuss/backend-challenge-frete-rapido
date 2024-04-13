package app

import (
	createQuoteCommand "github.com/josemateuss/backend-challenge-frete-rapido/app/command/create_quote"
	readQuoteMetricsQuery "github.com/josemateuss/backend-challenge-frete-rapido/app/query/read_quote_metrics"
	"github.com/josemateuss/backend-challenge-frete-rapido/app/repository"
	"github.com/josemateuss/backend-challenge-frete-rapido/app/service"
)

func NewApplication(repositories Repositories, services Services) (Application, error) {
	createQuote, err := createQuoteCommand.New(repositories.CreateQuoteRepository, services.SimulateQuoteService)
	if err != nil {
		return Application{}, err
	}

	readQuoteMetrics, err := readQuoteMetricsQuery.New(repositories.ReadQuoteRepository)
	if err != nil {
		return Application{}, err
	}

	return Application{
		Commands: Commands{
			CreateQuote: createQuote,
		},
		Queries: Queries{
			ReadQuoteMetrics: readQuoteMetrics,
		},
	}, nil
}

type Repositories struct {
	CreateQuoteRepository repository.CreateQuote
	ReadQuoteRepository   repository.ReadQuoteMetrics
}

type Services struct {
	SimulateQuoteService service.SimulateQuote
}

type Commands struct {
	CreateQuote createQuoteCommand.UseCase
}

type Queries struct {
	ReadQuoteMetrics readQuoteMetricsQuery.UseCase
}

type Application struct {
	Commands Commands
	Queries  Queries
}
