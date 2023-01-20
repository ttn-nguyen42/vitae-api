package experiences

import (
	"Vitae/repositories"

	 "Vitae/tools/utils"
     "go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	col     *mongo.Collection
}

func New(client *mongo.Client) *Repository {
	return &Repository{
		col:     client.Database(utils.GetDatabaseName(repositories.CVDatabaseName)).Collection(repositories.CollectionExperiences),
	}
}

func (repo *Repository) GetAll() {

}

func (repo *Repository) AddOne() {

}

func (repo *Repository) UpdateOne() {

}
