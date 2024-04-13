package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	useCasePackage "github.com/josemateuss/backend-challenge-frete-rapido/app/command/create_quote"
	"github.com/josemateuss/backend-challenge-frete-rapido/cmd/web/api/presenter"
	"github.com/josemateuss/backend-challenge-frete-rapido/cmd/web/internal"
)

func (h handler) CreateQuote(w http.ResponseWriter, r *http.Request) {
	response := internal.Response{}
	input := useCasePackage.Input{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Printf("error decoding input: %v", err)
		response.WriteError(w, http.StatusBadRequest, "failed to decode input")
		return
	}

	err = h.application.Commands.CreateQuote.Validate(input)
	if err != nil {
		var validationError *useCasePackage.ValidationError
		if errors.As(err, &validationError) {
			log.Printf("error validating input: %v", validationError.InvalidArguments)
			response.Write(w, http.StatusUnprocessableEntity, map[string][]string{
				"invalid_arguments": validationError.InvalidArguments,
			})
		}
		return
	}

	output, err := h.application.Commands.CreateQuote.Execute(r.Context(), input)
	if err != nil {
		log.Printf("error creating quote: %v", err)
		response.WriteError(w, http.StatusInternalServerError, "failed to create quote")
		return
	}

	response.Write(w, http.StatusOK, presenter.NewCreateQuotePresenter().Present(
		presenter.CreateQuotePresentInput{
			Quote: output.Quote,
		}))
}
