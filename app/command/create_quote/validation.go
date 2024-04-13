package create_quote

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/josemateuss/backend-challenge-frete-rapido/app/service"
)

type ValidationError struct {
	InvalidArguments []string
}

func (e *ValidationError) Error() string {
	return "invalid arguments"
}

func (uc UseCase) Validate(ctx context.Context, input Input) error {
	invalidArguments := make([]string, 0)

	if input.Recipient.Address.Zipcode == "" {
		invalidArguments = append(invalidArguments, "recipient_address_zipcode is required")
	}

	for i, volume := range input.Volumes {
		if volume.Category == 0 {
			invalidArguments = append(invalidArguments, fmt.Sprintf("volumes[%d].category is required", i))
		}

		if volume.Amount == 0 {
			invalidArguments = append(invalidArguments, fmt.Sprintf("volumes[%d].amount is required", i))
		}

		if volume.UnitaryWeight == 0 {
			invalidArguments = append(invalidArguments, fmt.Sprintf("volumes[%d].unitary_weight is required", i))
		}

		if volume.Price == 0 {
			invalidArguments = append(invalidArguments, fmt.Sprintf("volumes[%d].price is required", i))
		}

		if volume.Height == 0 {
			invalidArguments = append(invalidArguments, fmt.Sprintf("volumes[%d].height is required", i))
		}

		if volume.Width == 0 {
			invalidArguments = append(invalidArguments, fmt.Sprintf("volumes[%d].width is required", i))
		}

		if volume.Length == 0 {
			invalidArguments = append(invalidArguments, fmt.Sprintf("volumes[%d].length is required", i))
		}
	}

	if input.Recipient.Address.Zipcode != "" {
		match, _ := regexp.MatchString(`^\d{8}$`, input.Recipient.Address.Zipcode)
		if !match {
			invalidArguments = append(invalidArguments, "recipient_address_zipcode invalid pattern")
			return &ValidationError{InvalidArguments: invalidArguments}
		}
	}

	if input.Recipient.Address.Zipcode != "" {
		validate, err := uc.validateZipcodeService.Validate(ctx, service.ValidateZipcodeInput{
			Zipcode: input.Recipient.Address.Zipcode,
		})
		if err != nil {
			log.Printf("error validating zipcode: %v", err)
		}

		if validate.Error {
			invalidArguments = append(invalidArguments, "recipient_address_zipcode not found")
		}
	}

	if len(invalidArguments) > 0 {
		return &ValidationError{InvalidArguments: invalidArguments}
	}

	return nil
}
