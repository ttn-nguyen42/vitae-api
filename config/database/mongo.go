package database

import (
	"Vitae/config"
	"Vitae/tools/logging"
	context2 "context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var Client *mongo.Client = nil
var Context = getContext()

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

func getContext() context2.Context {
	context, _ := context2.WithTimeout(context2.Background(), 5*time.Second)
	return context
}

func Connect() *mongo.Client {
	driver, err, databaseId := getDriver()
	if err != nil {
		logging.Fatal(err.Error())
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(driver))
	if err != nil {
		logging.Fatal(err.Error())
	}
	err = client.Connect(Context)
	if err != nil {
		logging.Fatal(err.Error())
	}
	err = client.Ping(Context, nil)
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
	if Client != nil {
		err := Client.Disconnect(Context)
		if err != nil {
			logging.Fatal(err.Error())
		}
	}
}
