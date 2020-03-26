package database

import (
    "github.com/go-redis/redis/v7"
    "log"
)

var redisClient *redis.UniversalClient

// InitRedis initialises a new Redis client (to be used globally)
func InitRedis(hosts []string, password string, db int) *redis.UniversalClient {
    if redisClient != nil {
        log.Println("Redis was initialised already. Re-using existing client.")
    } else {
        newClient := redis.NewUniversalClient(&redis.UniversalOptions{
            Addrs: hosts,
            Password: password,
            DB: db,
        })
        redisClient = &newClient
    }

    return redisClient
}

// PingRedis provides health check for Redis database
//
// Note: Expected response from Redis is a constant, "PONG"
func PingRedis() bool {
    if redisClient == nil {
        return false
    }

    pong, err := (*redisClient).Ping().Result()
    return pong == "PONG" && err == nil
}
