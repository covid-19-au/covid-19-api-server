package database

import (
	"context"
	"github.com/Kamva/mgm"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var mongoClient *mongo.Client

// InitMongo initialises a new MongoDB client (to be used globally)
func InitMongo(dbName string, connString string) (*mongo.Client, error) {
	if mongoClient != nil {
		log.Println("MongoDB was initialised already. Re-using existing client.")
		return mongoClient, nil
	}

	mongoClientOptions := options.Client().ApplyURI(connString)

	// this is redundant but it's the only way to initialise the default config
	_ = mgm.SetDefaultConfig(nil, dbName, mongoClientOptions)

	var err error
	if mongoClient, err = mgm.NewClient(mongoClientOptions); err != nil { // the actual client in use is this one
		return nil, err
	} else {
		mongoClient.Database(dbName)
		return mongoClient, nil
	}
}

// PingMongo provides health check for MongoDB database
func PingMongo(timeout time.Duration) bool {
	if mongoClient == nil {
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err := mongoClient.Ping(ctx, readpref.Primary())

	return err == nil
}
