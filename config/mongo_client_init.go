package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type MongoClient struct {
	client *mongo.Client
}

func NewMongoClient(cfg *AppConfig) *MongoClient {
	// Initialize the database connection using cfg.MongoURI
	// This is a placeholder for actual database initialization logic
	uri := cfg.MongoURI
	println("Database initialized with URI:", cfg.MongoURI)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	println("Server API version set to:", serverAPI)

	var opts *options.ClientOptions
	opts = options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client , err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		// Panic on failure to connect, as this is critical.
		panic(err)
	}

	//check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		// Panic on failure to ping, as this indicates a connection issue.
		panic(err)
	}

	dbClient := &MongoClient{
		client: client,
	}
	return dbClient

}