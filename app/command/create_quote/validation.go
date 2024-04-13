package create_quote

import "fmt"

type ValidationError struct {
	InvalidArguments []string
}

func (e *ValidationError) Error() string {
	return "invalid arguments"
}

func (uc UseCase) Validate(input Input) error {
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

	if len(invalidArguments) > 0 {
		return &ValidationError{InvalidArguments: invalidArguments}
	}

	return nil
}
