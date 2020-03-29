package service

import (
    "fmt"
    "github.com/Kamva/mgm"
    m "github.com/covid-19-au/covid-19-api-server/model"
    pb "github.com/covid-19-au/covid-19-api-server/protogen/go"
    u "github.com/covid-19-au/covid-19-api-server/util"
    "github.com/golang/protobuf/ptypes/wrappers"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/net/context"
    "google.golang.org/grpc/grpclog"
)

func (s PublicServer) GetCases(context.Context, *pb.GetCasesRequest) (*pb.GetCasesResponse, error) {
	panic("implement me")
}

func (s AdminServer) AddCases(ctx context.Context, req *pb.AddCasesRequest) (*wrappers.BoolValue, error) {
	if req == nil || len(req.Cases) < 1 {
		return u.WrapBool(false), fmt.Errorf("empty or invalid request")
	}

	var caseModels []*m.Case
	for _, c := range req.Cases { // loop each case and build model
		if newModel, err := m.NewCase(c); err != nil {
			grpclog.Warningf("error occurred when adding case: %v\n", err)
			return u.WrapBool(false), fmt.Errorf("empty or invalid case details")
		} else {
			caseModels = append(caseModels, newModel) // append to model list
		}
	}

	// run transaction to insert all generated models
	err := mgm.Transaction(func(session mongo.Session, sc mongo.SessionContext) error {
		for _, c := range caseModels {
			if err := mgm.Coll(c).CreateWithCtx(sc, c); err != nil { // perform insertion
				return err
			}
		}

		return session.CommitTransaction(sc) // commit transaction if successful
	})

	return u.WrapBool(err == nil), err
}

func (s AdminServer) PutCases(context.Context, *pb.PutCasesRequest) (*wrappers.BoolValue, error) {
	panic("implement me")
}

func (s AdminServer) DelCases(context.Context, *pb.DelCasesRequest) (*wrappers.BoolValue, error) {
	panic("implement me")
}
