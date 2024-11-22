package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func init() {
	if err := ConnectMongoDB(); err != nil {
		fmt.Println(err.Error())
		log.Fatal("error while connecting to MongoDB")
	}

	log.Println("Connected to MongoDB!!🆙")
}

func ConnectMongoDB() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		panic("Error Loading .env File, Check If It Exists.")
	}

	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverApi)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	MongoClient = client
	return err
}
