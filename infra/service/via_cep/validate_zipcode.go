package via_cep

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/josemateuss/backend-challenge-frete-rapido/app/service"
)

const ViaCepURL = "https://viacep.com.br/ws/%s/json/"

func NewService() Service {
	return Service{}
}

func (s Service) Validate(ctx context.Context, input service.ValidateZipcodeInput) (
	output *service.ValidateZipcodeOutput, err error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(ViaCepURL, input.Zipcode), nil)
	if err != nil {
		return output, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req = req.WithContext(ctx)

	var httpClient = &http.Client{Transport: http.DefaultTransport}
	resp, err := httpClient.Do(req)
	if err != nil {
		return output, fmt.Errorf("failed to send request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return output, fmt.Errorf("GET via_cep returned status-code: %d", resp.StatusCode)
	}

	responsePayload := ResponsePayload{}
	err = json.NewDecoder(resp.Body).Decode(&responsePayload)
	if err != nil {
		return output, fmt.Errorf("failed to decode response body: %v", err)
	}

	output = &service.ValidateZipcodeOutput{
		Error:   responsePayload.Error,
		Zipcode: responsePayload.Zipcode,
	}

	return output, nil
}
