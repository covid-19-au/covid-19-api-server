package service

import (
    "context"
    pb "github.com/covid-19-au/covid-19-api-server/protogen/go"
    "github.com/golang/protobuf/ptypes/wrappers"
)

func (s PublicServer) GetFlights(context.Context, *pb.GetFlightsRequest) (*pb.GetFlightsResponse, error) {
	panic("implement me")
}

func (s AdminServer) AddFlights(context.Context, *pb.AddFlightsRequest) (*wrappers.BoolValue, error) {
	panic("implement me")
}

func (s AdminServer) DelFlights(context.Context, *pb.DelFlightsRequest) (*wrappers.BoolValue, error) {
	panic("implement me")
}
