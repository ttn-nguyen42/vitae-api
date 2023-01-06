package database

import (
	"Vitae/config"
	"Vitae/tools/logging"
	context2 "context"
	"errors"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = nil

func getDriver() (string, error, string) {
	clusterId := os.Getenv(config.EnvMongoClusterID)
	if clusterId == "" {
		return "", errors.New("missing MongoDB cluster ID"), ""
	}
	username := os.Getenv(config.EnvMongoUsername)
	if username == "" {
		return "", errors.New("missing MongoDB username"), clusterId
	}
	password := os.Getenv(config.EnvMongoPassword)
	if password == "" {
		return "", errors.New("missing MongoDB password"), clusterId
	}
	settings := "?retryWrites=true&w=majority"
	driver := fmt.Sprintf("mongodb+srv://%v:%v@%v.mongodb.net/%v", username, password, clusterId, settings)
	return driver, nil, clusterId
}

func GetContext() (context2.Context, context2.CancelFunc) {
	context, cancel := context2.WithTimeout(context2.Background(), 5*time.Second)
	return context, cancel
}

func Connect() *mongo.Client {
	driver, err, databaseId := getDriver()
	context, cancel := GetContext()
	if err != nil {
		logging.Fatal(err.Error())
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(driver))
	if err != nil {
		logging.Fatal(err.Error())
	}
	err = client.Connect(context)
	defer cancel()
	if err != nil {
		logging.Fatal(err.Error())
	}
	err = client.Ping(context, nil)
	if err != nil {
		logging.Fatal(err.Error())
	}
	logging.Info(fmt.Sprintf("Database %v connected successfully", databaseId))
	if Client == nil {
		Client = client
	}
	return client
}

func Close() {
	context, cancel := GetContext()
	defer cancel()
	if Client != nil {
		err := Client.Disconnect(context)
		if err != nil {
			logging.Fatal(err.Error())
		}
	}
}
