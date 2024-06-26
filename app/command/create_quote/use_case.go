package create_quote

import (
	"context"
	"fmt"
	"log"

	"github.com/josemateuss/backend-challenge-frete-rapido/app/repository"
	"github.com/josemateuss/backend-challenge-frete-rapido/app/service"
)

func New(
	createQuoteRepository repository.CreateQuote,
	validateZipcodeService service.ValidateZipCode,
	simulateQuoteService service.SimulateQuote) (UseCase, error) {
	if createQuoteRepository == nil {
		return UseCase{}, fmt.Errorf("repository is required")
	}

	return UseCase{
		createQuoteRepository:  createQuoteRepository,
		simulateQuoteService:   simulateQuoteService,
		validateZipcodeService: validateZipcodeService,
	}, nil
}

func (uc UseCase) Execute(ctx context.Context, input Input) (output Output, err error) {
	simulateOutput, err := uc.simulateQuoteService.Simulate(ctx, serviceSimulateQuoteInput(input))
	if err != nil {
		log.Printf("error simulating quote: %v", err)
		return output, fmt.Errorf("error simulating quote: %v", err)
	}

	if len(simulateOutput.Carrier) == 0 {
		return output, nil
	}

	createInput := repository.CreateQuoteInput{}
	for _, carrier := range simulateOutput.Carrier {
		createInput.Carrier = append(createInput.Carrier, repository.Carrier{
			Name:     carrier.Name,
			Service:  carrier.Service,
			Deadline: carrier.Deadline,
			Price:    carrier.Price,
		})
	}

	repositoryOutput, err := uc.createQuoteRepository.CreateQuote(ctx, createInput)
	if err != nil {
		log.Printf("error creating quote: %v", err)
		return output, fmt.Errorf("error saving quote on database: %v", err)
	}

	return Output{
		Quote: repositoryOutput.Quote,
	}, err
}

func serviceSimulateQuoteInput(input Input) service.SimulateQuotesInput {
	volumes := make([]service.Volume, len(input.Volumes))
	for i, volume := range input.Volumes {
		volumes[i] = service.Volume{
			Category:      volume.Category,
			Amount:        volume.Amount,
			UnitaryWeight: volume.UnitaryWeight,
			Price:         volume.Price,
			Sku:           volume.Sku,
			Height:        volume.Height,
			Width:         volume.Width,
			Length:        volume.Length,
		}
	}

	simulateQuoteInput := service.SimulateQuotesInput{
		Recipient: service.Recipient{
			Address: service.Address{
				Zipcode: input.Recipient.Address.Zipcode,
			},
		},
		Volumes: volumes,
	}

	return simulateQuoteInput
}
