package model

import (
    "fmt"
    "github.com/Kamva/mgm"
    pb "github.com/covid-19-au/covid-19-api-server/protogen/go"
)

type Case struct {
    mgm.DefaultModel `bson:",inline"` // defines _id, created_at and updated_at

    PatientID       string       `json:"patient_id" bson:"patient_id"`
    Location        CaseLocation `json:"location" bson:"location"`
    CaseState       string       `json:"case_state" bson:"case_state"`
    InfectionSource string       `json:"inf_src" bson:"inf_src"`
    ReportedTime    int64        `json:"report_ts" bson:"report_ts"`
}

type CaseLocation struct {
    Country    string `json:"country" bson:"country"`
    Region     string `json:"region" bson:"region"`
    City       string `json:"city" bson:"city"`
    PostalCode string `json:"post_code" bson:"post_code"`
    Latitude     float64 `json:"lat" bson:"lat"`
    Longitude    float64 `json:"lng" bson:"lng"`
    PrecisionRad uint64  `json:"prec_rad" bson:"prec_rad"`
}

func NewCase(req *pb.NewCaseDetail) (*Case, error) {
    if req == nil {
        return nil, fmt.Errorf("empty or invalid case details")
    }

    caseDetail := *req

    // run proto validator
    if err := caseDetail.Validate(); err != nil {
        return nil, fmt.Errorf("validation error: %v", err)
    }
    // additional checks for enums
    if val, ok := pb.CaseState_value[caseDetail.GetState().String()]; !ok || val < 1 {
        return nil, fmt.Errorf("validation error: invalid case state")
    }
    if val, ok := pb.InfectionSource_value[caseDetail.GetInfectSrc().String()]; !ok || val < 1 {
        return nil, fmt.Errorf("validation error: invalid case infection source")
    }

    caseLocation := caseDetail.GetLocation()

    return &Case{
        PatientID: caseDetail.GetPatientId(),
        Location: CaseLocation{
            Country:      caseLocation.GetCountry(),
            Region:       caseLocation.GetRegion(),
            City:         caseLocation.GetCity(),
            PostalCode:   caseLocation.GetPostalCode(),
            Latitude:     caseLocation.GetGeoLat(),
            Longitude:    caseLocation.GetGeoLng(),
            PrecisionRad: caseLocation.GetGeoPrecisionRad(),
        },
        CaseState:       caseDetail.GetState().String(),
        InfectionSource: caseDetail.GetInfectSrc().String(),
        ReportedTime:    caseDetail.GetReportedTime(),
    }, nil
}

func (m *Case) GetCaseState() pb.CaseState {
    if val, ok := pb.CaseState_value[m.CaseState]; !ok {
        return 0
    } else {
        return pb.CaseState(val)
    }
}

func (m *Case) GetInfectionSource() pb.InfectionSource {
    if val, ok := pb.InfectionSource_value[m.InfectionSource]; !ok {
        return 0
    } else {
        return pb.InfectionSource(val)
    }
}

func (m *Case) GetLocation() *pb.Location {
    loc := m.Location

    return &pb.Location{
        Country:         loc.Country,
        Region:          loc.Region,
        City:            loc.City,
        PostalCode:      loc.PostalCode,
        GeoLat:          loc.Latitude,
        GeoLng:          loc.Longitude,
        GeoPrecisionRad: loc.PrecisionRad,
    }
}
