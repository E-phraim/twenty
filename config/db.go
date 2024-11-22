package config

// func LoadMongo() string {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("error loading .env file")
// 	}
// 	return os.Getenv("MONGO_URI")
// }

// func ConnectMongoDB() *mongo.Client {
// 	client, err := mongo.NewClient(options.Client().ApplyURI(LoadMongo()))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = client.Ping(ctx, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("connected to MongoDB")
// 	return client
// }

// func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
// 	collection := client.Database("twenty-api").Collection(collectionName)
// 	return collection
// }
