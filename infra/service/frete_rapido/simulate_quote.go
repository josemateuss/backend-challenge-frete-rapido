package frete_rapido

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/josemateuss/backend-challenge-frete-rapido/app/service"
	"github.com/josemateuss/backend-challenge-frete-rapido/conf"
)

const FreteRapidoSimulateURL = "https://sp.freterapido.com/api/v3/quote/simulate"

func NewService() Service {
	return Service{}
}

func (s Service) Simulate(ctx context.Context, input service.SimulateQuotesInput) (
	output *service.SimulateQuotesOutput, err error) {
	requestPayload, err := newRequestPayload(input)
	if err != nil {
		return output, fmt.Errorf("failed to create request payload: %v", err)
	}

	payload, err := json.Marshal(requestPayload)
	if err != nil {
		return output, fmt.Errorf("failed to marshal request payload: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, FreteRapidoSimulateURL, bytes.NewReader(payload))
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
		return output, fmt.Errorf("POST frete_rapido simulate returned status-code: %d", resp.StatusCode)
	}

	responsePayload := ResponsePayload{}
	err = json.NewDecoder(resp.Body).Decode(&responsePayload)
	if err != nil {
		return output, fmt.Errorf("failed to decode response body: %v", err)
	}

	return mapResponseToOutput(responsePayload, output), nil
}

func newRequestPayload(input service.SimulateQuotesInput) (payload RequestPayload, err error) {
	recipientZipcode := input.Recipient.Address.Zipcode
	recipientZipcodeUint, err := strconv.ParseUint(recipientZipcode, 10, 64)
	if err != nil {
		log.Printf("failed to parse recipient zipcode: %s, err: %v", recipientZipcode, err)
		return payload, err
	}

	volumes := make([]Volume, len(input.Volumes))
	for i, v := range input.Volumes {
		volumes[i] = Volume{
			Amount:        v.Amount,
			AmountVolumes: 1,
			Category:      strconv.Itoa(int(v.Category)),
			Height:        v.Height,
			Sku:           v.Sku,
			Width:         v.Width,
			Length:        v.Length,
			UnitaryPrice:  float64(v.Price),
			UnitaryWeight: float64(v.UnitaryWeight),
		}

	}

	payload = RequestPayload{
		Shipper: Shipper{
			RegisteredNumber: conf.FreteRapidoRegisteredNumber,
			Token:            conf.FreteRapidoToken,
			PlatformCode:     conf.FreteRapidoPlatformCode,
		},
		Recipient: Recipient{
			Zipcode: uint(recipientZipcodeUint),
		},
		Dispatchers: []Dispatcher{
			{
				RegisteredNumber: conf.FreteRapidoRegisteredNumber,
				Zipcode:          conf.FreteRapidoZipcode,
				Volumes:          volumes,
			},
		},
		SimulationType: []uint{0},
	}

	return payload, nil
}

func mapResponseToOutput(responsePayload ResponsePayload, output *service.SimulateQuotesOutput) *service.SimulateQuotesOutput {
	carriers := make([]service.Carrier, 0)
	for _, dispatcher := range responsePayload.Dispatchers {
		for _, offer := range dispatcher.Offers {
			carrier := service.Carrier{
				Name:     offer.Carrier.Name,
				Service:  offer.Service,
				Deadline: offer.DeliveryTime.Days,
				Price:    offer.FinalPrice,
			}
			carriers = append(carriers, carrier)
		}
	}

	output = &service.SimulateQuotesOutput{
		Carrier: carriers,
	}

	return output
}
