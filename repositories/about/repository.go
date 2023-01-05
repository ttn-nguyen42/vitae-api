package about

import (
	context2 "context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	client  *mongo.Client
	context context2.Context
}

func New(client *mongo.Client, ctx context2.Context) *Repository {
	return &Repository{
		client:  client,
		context: ctx,
	}
}
