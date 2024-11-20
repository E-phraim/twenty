package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = ""

var MongoClient *mongo.Client

func init() {
	if err := ConnectMongoDB(); err != nil {
		log.Fatal("error while connecting to MongoDB")
	}
}

func ConnectMongoDB() error {
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverApi)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	MongoClient = client
	return err
}
