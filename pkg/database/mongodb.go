package database

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var mongoCollection *mongo.Collection

func MongoDBInit() error {
	mongoHost := viper.GetString("DATABASE_HOST")
	mongoPort := viper.GetString("DATABASE_PORT")
	mongoDbname := viper.GetString("DATABASE_DBNAME")

	mongoDns := fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)
	clientOptions := options.Client().ApplyURI(mongoDns)
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	slog.Info("Connected to MongoDB!")

	mongoCollection = client.Database(mongoDbname).Collection("Job")
	return nil
}

func DisconnectMongoDB() error {
	err := client.Disconnect(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func GetMondoCollection() *mongo.Collection {
	return mongoCollection
}
