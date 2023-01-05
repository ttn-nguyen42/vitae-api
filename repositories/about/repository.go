package about

import (
	"Vitae/config"
	"Vitae/repositories"
	context2 "context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	col     *mongo.Collection
	context context2.Context
}

func New(client *mongo.Client, ctx context2.Context) *Repository {
	return &Repository{
		col:     client.Database(config.CVDatabaseName).Collection(repositories.CollectionAbout),
		context: ctx,
	}
}

func (repo *Repository) GetAll() {

}

func (repo *Repository) AddOne() {

}

func (repo *Repository) UpdateOne() {

}
