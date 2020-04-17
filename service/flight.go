package service

import (
    "context"
    "fmt"
    pb "github.com/covid-19-au/covid-19-api-server/protogen/go"
    "github.com/golang/protobuf/ptypes/wrappers"
)

func (s PublicServer) GetFlights(context.Context, *pb.GetFlightsRequest) (*pb.GetFlightsResponse, error) {
    return nil, fmt.Errorf("not implemented")
}

func (s AdminServer) AddFlights(context.Context, *pb.AddFlightsRequest) (*wrappers.BoolValue, error) {
    return nil, fmt.Errorf("not implemented")
}

func (s AdminServer) DelFlights(context.Context, *pb.DelFlightsRequest) (*wrappers.BoolValue, error) {
    return nil, fmt.Errorf("not implemented")
}
