package main

import (
    "github.com/covid-19-au/covid-19-api-server/database"
    "github.com/covid-19-au/covid-19-api-server/rpc"
    "log"
    "os"
)

func main() {
    log.Printf("Starting API server...")

    if _, err := database.InitMongo(os.Getenv("MONGO_DB_NAME"), os.Getenv("MONGO_CONN_STR")); err != nil {
        log.Fatalf("failed to initialise MongoDB connection: %v\n", err)
    }

    if err := rpc.InitServer(os.Getenv("SERVER_LISTEN_ADDR")); err != nil {
        log.Fatalf("failed to initialise RPC server: %v\n", err)
    }
}
