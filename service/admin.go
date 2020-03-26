package service

import (
    "context"
    "github.com/covid-19-au/covid-19-api-server/database"
    pb "github.com/covid-19-au/covid-19-api-server/protogen"
    "github.com/golang/protobuf/ptypes/empty"
    "github.com/golang/protobuf/ptypes/wrappers"
    "time"
)

type AdminServer struct{}

func (s AdminServer) IsServiceReady(context.Context, *empty.Empty) (*wrappers.BoolValue, error) {
    isMongoReady := database.PingMongo(5 * time.Second)
    isRedisReady := database.PingRedis()

    checkResult := isMongoReady && isRedisReady

    return &wrappers.BoolValue{Value: checkResult}, nil
}

func (s AdminServer) AddCases(context.Context, *pb.AddCasesRequest) (*wrappers.BoolValue, error) {
    panic("implement me")
}

func (s AdminServer) PutCases(context.Context, *pb.PutCasesRequest) (*wrappers.BoolValue, error) {
    panic("implement me")
}

func (s AdminServer) DelCases(context.Context, *pb.DelCasesRequest) (*wrappers.BoolValue, error) {
    panic("implement me")
}

func (s AdminServer) AddFlights(context.Context, *pb.AddFlightsRequest) (*wrappers.BoolValue, error) {
    panic("implement me")
}

func (s AdminServer) DelFlights(context.Context, *pb.DelFlightsRequest) (*wrappers.BoolValue, error) {
    panic("implement me")
}
