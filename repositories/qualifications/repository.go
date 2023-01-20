package qualifications

import (
	"Vitae/repositories"

	 "Vitae/tools/utils"
     "go.mongodb.org/mongo-driver/mongo"

	context2 "context"
)

type Repository struct {
	col     *mongo.Collection
	context context2.Context
}

func New(client *mongo.Client) *Repository {
	return &Repository{
		col:     client.Database(utils.GetDatabaseName(repositories.CVDatabaseName)).Collection(repositories.CollectionQualifications),
	}
}

func (repo *Repository) GetAll() {

}

func (repo *Repository) AddOne() {

}

func (repo *Repository) UpdateOne() {

}
