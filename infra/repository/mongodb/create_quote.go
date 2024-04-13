package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/josemateuss/backend-challenge-frete-rapido/app/repository"
	"github.com/josemateuss/backend-challenge-frete-rapido/conf"
	"github.com/josemateuss/backend-challenge-frete-rapido/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type createQuote struct {
	Client *mongo.Client
}

func NewCreateQuote(client *mongo.Client) createQuote {
	return createQuote{
		Client: client,
	}
}

func (repo createQuote) CreateQuote(ctx context.Context, input repository.CreateQuoteInput) (
	output *repository.CreateQuoteOutput, err error) {
	collection := repo.Client.Database(conf.DatabaseName).Collection("quotes")
	input.CreatedAt = time.Now()
	quoteResult, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Printf("error creating quote on database: %v", err)
		return output, err
	}

	objectID, ok := quoteResult.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Printf("error converting ID to ObjectID")
		return output, fmt.Errorf("error converting ID to ObjectID")
	}

	var quote quoteDB
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&quote)
	if err != nil {
		log.Printf("error finding document: %v", err)
		return output, err
	}

	output = quote.createQuoteToDomain()

	return output, nil
}

func (q quoteDB) createQuoteToDomain() *repository.CreateQuoteOutput {
	carriers := make([]domain.Carrier, len(q.Carriers))
	for i, carrier := range q.Carriers {
		carriers[i] = domain.Carrier{
			Name:     carrier.Name,
			Service:  carrier.Service,
			Deadline: carrier.Deadline,
			Price:    carrier.Price,
		}
	}

	return &repository.CreateQuoteOutput{
		Quote: &domain.Quote{
			Carrier: carriers,
		},
	}
}
