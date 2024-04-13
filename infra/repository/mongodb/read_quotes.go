package mongodb

import (
	"context"
	"fmt"
	"log"

	"github.com/josemateuss/backend-challenge-frete-rapido/app/repository"
	"github.com/josemateuss/backend-challenge-frete-rapido/conf"
	"github.com/josemateuss/backend-challenge-frete-rapido/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type readQuotes struct {
	Client *mongo.Client
}

func NewReadQuotes(client *mongo.Client) readQuotes {
	return readQuotes{
		Client: client,
	}
}

func (repo readQuotes) ReadQuotes(ctx context.Context, input repository.ReadQuotesInput) (
	output *repository.ReadQuotesOutput, err error) {
	collection := repo.Client.Database(conf.DatabaseName).Collection("quotes")
	var cursor *mongo.Cursor
	if input.LastQuotes == 0 {
		cursor, err = collection.Find(ctx, bson.M{}, options.Find())
	} else {
		lastQuotes := int64(input.LastQuotes)
		cursor, err = collection.Find(ctx, bson.M{}, options.Find().
			SetLimit(lastQuotes).
			SetSort(bson.D{{"created_at", -1}}))
	}
	if err != nil {
		log.Printf("error finding quotes: %v", err)
		return output, err
	}

	var quotes []quoteDB
	err = cursor.All(ctx, &quotes)
	if err != nil {
		log.Printf("error converting cursor to array: %v", err)
		return output, err
	}

	if quotes == nil {
		return output, fmt.Errorf("no quotes found")
	}

	output = &repository.ReadQuotesOutput{
		Quotes: make([]domain.Quote, len(quotes)),
	}

	for i, quote := range quotes {
		output.Quotes[i] = *quote.readQuotesToDomain()
	}

	return output, nil
}

func (q quoteDB) readQuotesToDomain() *domain.Quote {
	carriers := make([]domain.Carrier, len(q.Carriers))
	for i, carrier := range q.Carriers {
		carriers[i] = domain.Carrier{
			Name:     carrier.Name,
			Service:  carrier.Service,
			Deadline: carrier.Deadline,
			Price:    carrier.Price,
		}
	}

	return &domain.Quote{
		Carrier: carriers,
	}
}
