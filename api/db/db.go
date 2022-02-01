package db

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"time"
)

//Connect InitialConnection, create new connection to mongo db
func Connect(dbName string, mongoURL string) (*mongo.Database, error) {

	clientOptions := options.Client().
		ApplyURI("")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to mongo: %v", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to mongo: %v", err)
	}
	return client.Database(dbName), nil
}
