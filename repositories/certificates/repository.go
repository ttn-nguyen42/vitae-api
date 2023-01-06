package certificates

import (
	"Vitae/repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	col     *mongo.Collection
}

func New(client *mongo.Client) *Repository {
	return &Repository{
		col:     client.Database(repositories.CVDatabaseName).Collection(repositories.CollectionCertificates),
	}
}

func (repo *Repository) GetAll() {

}

func (repo *Repository) AddOne() {

}

func (repo *Repository) UpdateOne() {

}
