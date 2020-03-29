package model

import (
    "errors"
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
}

type CaseLocation struct {
	NamedLocation *NamedCaseLocation `json:"named" bson:"named"`
	GPSLocation   *GPSCaseLocation   `json:"gps" bson:"gps"`
}

type NamedCaseLocation struct {
	Country    string `json:"country" bson:"country"`
	Region     string `json:"region" bson:"region"`
	City       string `json:"city" bson:"city"`
	PostalCode string `json:"post_code" bson:"post_code"`
}

type GPSCaseLocation struct {
	Latitude     float64 `json:"lat" bson:"lat"`
	Longitude    float64 `json:"lng" bson:"lng"`
	PrecisionRad uint64  `json:"prec_rad" bson:"prec_rad"`
}

func NewCase(req *pb.NewCaseDetail) (*Case, error) {
	if req == nil {
		return nil, fmt.Errorf("empty or invalid case details")
	}

	caseDetail := *req
	var namedCaseLocation *NamedCaseLocation
	var gpsCaseLocation *GPSCaseLocation

	if loc := caseDetail.GetNamedLocation(); loc != nil {
		namedCaseLocation = &NamedCaseLocation{
			Country:    loc.GetCountry(),
			Region:     loc.GetRegion(),
			City:       loc.GetCity(),
			PostalCode: loc.GetPostalCode(),
		}
	}

	if loc := caseDetail.GetGpsLocation(); loc != nil {
		gpsCaseLocation = &GPSCaseLocation{
			Latitude:     loc.GetGeoLat(),
			Longitude:    loc.GetGeoLng(),
			PrecisionRad: loc.GetPrecisionRad(),
		}
	}

	if namedCaseLocation == nil && gpsCaseLocation == nil {
		return nil, fmt.Errorf("specify either named or GPS case location")
	}

	return &Case{
		PatientID: caseDetail.GetPatientId(),
		Location: CaseLocation{
			NamedLocation: namedCaseLocation,
			GPSLocation:   gpsCaseLocation,
		},
		CaseState:       caseDetail.GetState().String(),
		InfectionSource: caseDetail.GetInfectSrc().String(),
	}, nil
}

func (m *Case) Saving() error {
	if err := m.DefaultModel.Saving(); err != nil {
	   return err
	}

	// validate case state
	if m.CaseState == "_UNASSIGNED_CASE_STATE" {
	   return errors.New("invalid case state")
	}

	// validate infection source
    if m.InfectionSource == "_UNASSIGNED_INF_SRC" {
        return errors.New("invalid case infection source")
    }

	return nil
}
