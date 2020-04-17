package service

import (
    "fmt"
    "github.com/Kamva/mgm"
    "github.com/Kamva/mgm/operator"
    m "github.com/covid-19-au/covid-19-api-server/model"
    pb "github.com/covid-19-au/covid-19-api-server/protogen/go"
    u "github.com/covid-19-au/covid-19-api-server/util"
    "github.com/golang/protobuf/ptypes/wrappers"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/net/context"
    "google.golang.org/grpc/grpclog"
)

func (s PublicServer) GetCases(ctx context.Context, req *pb.GetCasesRequest) (*pb.GetCasesResponse, error) {
    if req == nil {
        return nil, fmt.Errorf("empty or invalid request")
    }

    var results []m.Case
    caseStates := req.GetStates()
    caseStartTime := req.GetStartTime()
    caseEndTime := req.GetEndTime()
    query := buildQueryFromRequest(req.GetLocation(), &caseStates, &caseStartTime, &caseEndTime) // build Mongo query

    if err := mgm.Coll(&m.Case{}).SimpleFind(&results, query); err != nil { // find matched cases
        grpclog.Errorf("error occurred when querying case: %v\n", err)
        return nil, fmt.Errorf("internal service error")
    } else {
        var cases []*pb.ExistingCaseDetail
        for _, c := range results { // for each case, convert it to protobuf struct
            cases = append(cases, &pb.ExistingCaseDetail{
                CaseId:       c.ID.Hex(),
                PatientId:    c.PatientID,
                ReportedTime: c.ReportedTime,
                State:        c.GetCaseState(),
                InfectSrc:    c.GetInfectionSource(),
                Location:     c.GetLocation(),
            })
        }

        return &pb.GetCasesResponse{
            Cases: cases,
            // TODO (Robin): implement pagination
            Pagination: &pb.Pagination{
                TotalResults: uint32(len(results)),
            },
        }, nil
    }
}

func (s PublicServer) GetCaseStats(ctx context.Context, req *pb.GetCaseStatsRequest) (*pb.GetCaseStatsResponse, error) {
    return nil, fmt.Errorf("not implemented")
}

func (s AdminServer) AddCases(ctx context.Context, req *pb.AddCasesRequest) (*wrappers.BoolValue, error) {
    if req == nil || len(req.Cases) < 1 {
        return u.WrapBool(false), fmt.Errorf("empty or invalid request")
    }

    var caseModels []*m.Case
    for _, c := range req.Cases { // loop each case and build model
        if newModel, err := m.NewCase(c); err != nil {
            grpclog.Warningf("error occurred when adding case: %v\n", err)
            return u.WrapBool(false), fmt.Errorf("invalid case details: %v", err)
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
    return nil, fmt.Errorf("not implemented")
}

func (s AdminServer) DelCases(context.Context, *pb.DelCasesRequest) (*wrappers.BoolValue, error) {
    return nil, fmt.Errorf("not implemented")
}

func buildQueryFromRequest(loc *pb.Location, states *[]pb.CaseState, startTime, endTime *int64) bson.M {
    query := bson.M{}

    // TODO (Robin): add case locations
    //if loc != nil {
    //    query["location"] = ""
    //}

    // case states to be included
    if states != nil {
        query["case_state"] = bson.D{{operator.In, states}}
    }

    // report starting timestamp
    if startTime != nil && *startTime > 0 {
       query["report_ts"] = bson.M{operator.Gte: *startTime}
    }

    // report ending timestamp: must be greater than starting time iff starting time is specified
    if endTime != nil && (startTime == nil || *endTime >= *startTime) {
        query["report_ts"] = bson.M{operator.Lte: *endTime}
    }

    return query
}
