package api

import (
	"log"
	"net/http"
	"strconv"

	useCasePackage "github.com/josemateuss/backend-challenge-frete-rapido/app/query/read_quote_metrics"
	"github.com/josemateuss/backend-challenge-frete-rapido/cmd/web/api/presenter"
	"github.com/josemateuss/backend-challenge-frete-rapido/cmd/web/internal"
)

func (h handler) ReadQuoteMetrics(w http.ResponseWriter, r *http.Request) {
	response := internal.Response{}
	input := useCasePackage.Input{}

	lastQuotesStr := r.URL.Query().Get("last_quotes")
	if lastQuotesStr != "" {
		lastQuotes, err := strconv.ParseUint(lastQuotesStr, 10, 32)
		if err != nil {
			log.Printf("error parsing query params: %v", err)
			response.WriteError(w, http.StatusBadRequest, "last_quotes must be a valid number")
			return
		}

		input.LastQuotes = uint(lastQuotes)
	}

	output, err := h.application.Queries.ReadQuoteMetrics.Execute(r.Context(), input)
	if err != nil {
		log.Printf("error reading quote metrics: %v", err)
		response.WriteError(w, http.StatusInternalServerError, "failed to read quote metrics")
		return
	}

	response.Write(w, http.StatusOK, presenter.NewReadQuoteMetricsPresenter().Present(
		presenter.ReadQuoteMetricsPresentInput{
			ResultsPerCarrier:      output.ResultsPerCarrier,
			TotalPricePerCarrier:   output.TotalPricePerCarrier,
			AveragePricePerCarrier: output.AveragePricePerCarrier,
			CheapestFreight:        output.CheapestFreight,
			MostExpensiveFreight:   output.MostExpensiveFreight,
		}))
}
