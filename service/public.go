package service

import (
    "context"
    "github.com/covid-19-au/covid-19-api-server/database"
    pb "github.com/covid-19-au/covid-19-api-server/protogen"
    "github.com/golang/protobuf/ptypes/empty"
    "github.com/golang/protobuf/ptypes/wrappers"
    "time"
)

type PublicServer struct{}

func (s PublicServer) IsServiceReady(context.Context, *empty.Empty) (*wrappers.BoolValue, error) {
    isMongoReady := database.PingMongo(10 * time.Second)
    isRedisReady := database.PingRedis()

    checkResult := isMongoReady && isRedisReady

    return &wrappers.BoolValue{Value: checkResult}, nil
}

func (s PublicServer) GetCases(context.Context, *pb.GetCasesRequest) (*pb.GetCasesResponse, error) {
    panic("implement me")
}

func (s PublicServer) GetFlights(context.Context, *pb.GetFlightsRequest) (*pb.GetFlightsResponse, error) {
    panic("implement me")
}
