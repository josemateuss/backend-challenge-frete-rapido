package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/josemateuss/backend-challenge-frete-rapido/app"
	handler "github.com/josemateuss/backend-challenge-frete-rapido/cmd/web/api"
	mongodbRepository "github.com/josemateuss/backend-challenge-frete-rapido/infra/repository/mongodb"
	freteRapidoService "github.com/josemateuss/backend-challenge-frete-rapido/infra/service/frete_rapido"
	viaCepService "github.com/josemateuss/backend-challenge-frete-rapido/infra/service/via_cep"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("could not connect to mongo: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("ping not working: %v", err)
	}

	log.Println("connected to mongo database")

	repositories := app.Repositories{
		CreateQuoteRepository: mongodbRepository.NewCreateQuote(client),
		ReadQuoteRepository:   mongodbRepository.NewReadQuotes(client),
	}

	services := app.Services{
		ValidateZipcode:      viaCepService.NewService(),
		SimulateQuoteService: freteRapidoService.NewService(),
	}

	application, err := app.NewApplication(repositories, services)
	if err != nil {
		log.Fatalf("error creating application: %s", err)
	}

	router := mux.NewRouter()
	h := handler.New(application)

	api := router.PathPrefix("/frete-rapido/api/v1").Subrouter()
	api.HandleFunc("/quote", h.CreateQuote).Methods(http.MethodPost)
	api.HandleFunc("/metrics", h.ReadQuoteMetrics).Methods(http.MethodGet)

	address := ":3000"
	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 60 * time.Second,
		Addr:         address,
		Handler:      router,
	}

	log.Fatal(server.ListenAndServe())
}
