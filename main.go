package main

import (
    "context"
    "github.com/covid-19-au/covid-19-api-server/database"
    "github.com/covid-19-au/covid-19-api-server/rpc"
    "log"
    "os"
)

func main() {
    log.Println("Hello World! Starting COVID-19 API Server...")

    // init MongoDB
    if mongoClient, err := database.InitMongo(os.Getenv("MONGO_DB_NAME"), os.Getenv("MONGO_CONN_STR")); err != nil {
        log.Fatalf("failed to initialise MongoDB connection: %v\n", err)
    } else {
        log.Println("Connected to MongoDB server.")
        defer func() { _ = mongoClient.Disconnect(context.TODO()) }()
    }

    // init Redis
    if redisClient := database.InitRedis([]string{":6379"}, "", 0); redisClient == nil {
        log.Fatalln("failed to initialise Redis connection")
    } else {
        log.Println("Connected to Redis server.")
        defer func() { _ = (*redisClient).Close() }()
    }

    // init gRPC server
    if err := rpc.InitServer(os.Getenv("RPC_LISTEN_ADDR")); err != nil {
        log.Fatalf("failed to initialise RPC server: %v\n", err)
    }
}
