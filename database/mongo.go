package database

import (
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

envFile, _ := godotenv.Read(".env")

var client *mongo.Client

func Init(uri string, database string) error {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoDBURI := os.Getenv("MONGODB_URL")
	clientOptions := options.Client().ApplyURI(mongoDBURI)

	localClient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	client = localClient

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}
}

func Close() error {
	return client.Disconnect(context.Background())
}