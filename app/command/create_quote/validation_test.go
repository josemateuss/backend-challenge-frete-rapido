package create_quote

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUseCase_Validate(t *testing.T) {
	uc := UseCase{}

	t.Run("Test with valid input", func(t *testing.T) {
		input := Input{
			Recipient: Recipient{
				Address: Address{
					Zipcode: "12345",
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

		err := uc.Validate(input)
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

		err := uc.Validate(input)
		assert.NotNil(t, err)
	})
}
