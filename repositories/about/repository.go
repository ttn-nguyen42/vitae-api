package about

import (
	"Vitae/config/database"
	"Vitae/models"
	"Vitae/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	col     *mongo.Collection
}

func New(client *mongo.Client) *Repository {
	return &Repository{
		col:     client.Database(repositories.CVDatabaseName).Collection(repositories.CollectionAbout),
	}
}

func (repo *Repository) GetAll(queryAmount int) ([]models.About, error) {
	context, cancel := database.GetContext()
	defer cancel()
	var results []models.About
	opts := options.Find().SetSort(bson.D{{Key: "_id", Value: 1}})
	cursor, err := repo.col.Find(context, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	if (queryAmount == repositories.QueryAll) {
		err := cursor.All(context, &results)
		if err != nil {
			return nil, repositories.NewInternalError(err.Error())
		}
		return results, nil
	}
	amount := queryAmount
	for cursor.Next(context) && amount > 0 {
		var result models.About
		err := cursor.Decode(&result)
		if err != nil {
			return nil, repositories.NewInternalError(err.Error())
		}
		results = append(results, result)
	}
	err = cursor.Err()
	if err != nil {
		return nil, repositories.NewInternalError(err.Error())
	}
	return results, nil
}

func (repo *Repository) GetOne(result *models.About, id string) error {
	context, cancel := database.GetContext()
	defer cancel()
	filter := bson.D{
		{Key: "_id", Value: id},
	}
	err := repo.col.FindOne(context, filter).Decode(result)
	if err == mongo.ErrNoDocuments {
		return repositories.NewNotFoundError(err.Error())
	}
	if err != nil {
		return repositories.NewInternalError(err.Error())
	}
	return nil
}

func (repo *Repository) AddOne(document models.About) (string, error) {
	context, cancel := database.GetContext()
	defer cancel()
	document.Id = primitive.NilObjectID
	result, err := repo.col.InsertOne(context, document)
	if err != nil {
		return "", repositories.NewInternalError(err.Error())
	}
	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", repositories.NewInternalError("id cannot be parsed")
	}
	return id.Hex(), nil
}

func (repo *Repository) UpdateOne() {

}
