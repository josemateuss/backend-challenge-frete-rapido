package via_cep

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/josemateuss/backend-challenge-frete-rapido/app/service"
	"github.com/stretchr/testify/assert"
)

func TestService_Validate(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	ctx := context.Background()
	input := service.ValidateZipcodeInput{
		Zipcode: "73340030",
	}
	expectedOutput := &service.ValidateZipcodeOutput{
		Error:   false,
		Zipcode: "73340030",
	}
	responsePayload := ResponsePayload{
		Zipcode: "73340030",
	}

	responder, _ := httpmock.NewJsonResponder(http.StatusOK, responsePayload)
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(ViaCepURL, input.Zipcode), responder)

	s := Service{}
	output, err := s.Validate(ctx, input)

	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, output)
}
