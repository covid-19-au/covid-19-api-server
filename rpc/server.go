package rpc

import (
    "fmt"
    pb "github.com/covid-19-au/covid-19-api-server/protogen"
    "github.com/covid-19-au/covid-19-api-server/service"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "net"
)

func InitServer(listenAddress string) error {
    if grpcServer, grpcListener, err := registerServer(listenAddress); err != nil {
        return err
    } else {
        registerServices(grpcServer)

        if err := grpcServer.Serve(*grpcListener); err != nil {
            return fmt.Errorf("failed to start gRPC service: %v\n", err)
        }
    }

    return nil
}

func registerServer(listenAddress string) (*grpc.Server, *net.Listener, error) {
    if listener, err := net.Listen("tcp", listenAddress); err != nil || listener == nil {
        return nil, nil, fmt.Errorf("failed to register admin service: %v\n", err)
    } else {
        grpcServer := grpc.NewServer()
        reflection.Register(grpcServer)

        return grpcServer, &listener, nil
    }
}

func registerServices(grpcServer *grpc.Server) {
    pb.RegisterPublicQueryServer(grpcServer, &service.PublicServer{})
    pb.RegisterDataAdminServer(grpcServer, &service.AdminServer{})
}
