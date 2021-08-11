package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	MongoUri string `mapstructure:"mongo_uri"`
	Client   *mongo.Client
}

func (m *MongoDB) BindingClient() error {

	// Set client options
	clientOptions := options.Client().ApplyURI(m.MongoUri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return err
	}

	fmt.Println("Connected to MongoDB!")
	m.Client = client
	return nil
}

func (m *MongoDB) GetDatabase() *mongo.Database {
	return m.Client.Database("LineWallet")
}
