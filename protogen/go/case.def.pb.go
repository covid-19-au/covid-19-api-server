// Code generated by protoc-gen-go. DO NOT EDIT.
// source: case.def.proto

package covid19api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// CaseState describes state of the case
type CaseState int32

const (
	CaseState__UNASSIGNED_CASE_STATE CaseState = 0
	CaseState_UNKNOWN                CaseState = 11
	CaseState_UNDISCLOSED            CaseState = 12
	CaseState_CONFIRMED              CaseState = 21
	CaseState_RECOVERED              CaseState = 22
	CaseState_DEATH                  CaseState = 23
)

var CaseState_name = map[int32]string{
	0:  "_UNASSIGNED_CASE_STATE",
	11: "UNKNOWN",
	12: "UNDISCLOSED",
	21: "CONFIRMED",
	22: "RECOVERED",
	23: "DEATH",
}

var CaseState_value = map[string]int32{
	"_UNASSIGNED_CASE_STATE": 0,
	"UNKNOWN":                11,
	"UNDISCLOSED":            12,
	"CONFIRMED":              21,
	"RECOVERED":              22,
	"DEATH":                  23,
}

func (x CaseState) String() string {
	return proto.EnumName(CaseState_name, int32(x))
}

func (CaseState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c687d06fd30aa5ff, []int{0}
}

// InfectionSource describes source of the infection
type InfectionSource int32

const (
	InfectionSource__UNASSIGNED_INF_SRC InfectionSource = 0
	InfectionSource_DOMESTIC            InfectionSource = 1
	InfectionSource_INTL_IMPORTED       InfectionSource = 2
	InfectionSource_SHIP_IMPORTED       InfectionSource = 3
)

var InfectionSource_name = map[int32]string{
	0: "_UNASSIGNED_INF_SRC",
	1: "DOMESTIC",
	2: "INTL_IMPORTED",
	3: "SHIP_IMPORTED",
}

var InfectionSource_value = map[string]int32{
	"_UNASSIGNED_INF_SRC": 0,
	"DOMESTIC":            1,
	"INTL_IMPORTED":       2,
	"SHIP_IMPORTED":       3,
}

func (x InfectionSource) String() string {
	return proto.EnumName(InfectionSource_name, int32(x))
}

func (InfectionSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c687d06fd30aa5ff, []int{1}
}

// ExistingCaseDetail describes an existing (exists in DB) case
type ExistingCaseDetail struct {
	CaseUuid     string               `protobuf:"bytes,1,opt,name=case_uuid,json=caseUuid,proto3" json:"case_uuid,omitempty"`
	PatientId    string               `protobuf:"bytes,2,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	ReportedTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=reported_time,json=reportedTime,proto3" json:"reported_time,omitempty"`
	State        CaseState            `protobuf:"varint,11,opt,name=state,proto3,enum=covid19api.CaseState" json:"state,omitempty"`
	InfectSrc    InfectionSource      `protobuf:"varint,12,opt,name=infect_src,json=infectSrc,proto3,enum=covid19api.InfectionSource" json:"infect_src,omitempty"`
	// Types that are valid to be assigned to Location:
	//	*ExistingCaseDetail_NamedLocation
	//	*ExistingCaseDetail_GpsLocation
	Location             isExistingCaseDetail_Location `protobuf_oneof:"location"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ExistingCaseDetail) Reset()         { *m = ExistingCaseDetail{} }
func (m *ExistingCaseDetail) String() string { return proto.CompactTextString(m) }
func (*ExistingCaseDetail) ProtoMessage()    {}
func (*ExistingCaseDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_c687d06fd30aa5ff, []int{0}
}

func (m *ExistingCaseDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExistingCaseDetail.Unmarshal(m, b)
}
func (m *ExistingCaseDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExistingCaseDetail.Marshal(b, m, deterministic)
}
func (m *ExistingCaseDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExistingCaseDetail.Merge(m, src)
}
func (m *ExistingCaseDetail) XXX_Size() int {
	return xxx_messageInfo_ExistingCaseDetail.Size(m)
}
func (m *ExistingCaseDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_ExistingCaseDetail.DiscardUnknown(m)
}

var xxx_messageInfo_ExistingCaseDetail proto.InternalMessageInfo

func (m *ExistingCaseDetail) GetCaseUuid() string {
	if m != nil {
		return m.CaseUuid
	}
	return ""
}

func (m *ExistingCaseDetail) GetPatientId() string {
	if m != nil {
		return m.PatientId
	}
	return ""
}

func (m *ExistingCaseDetail) GetReportedTime() *timestamp.Timestamp {
	if m != nil {
		return m.ReportedTime
	}
	return nil
}

func (m *ExistingCaseDetail) GetState() CaseState {
	if m != nil {
		return m.State
	}
	return CaseState__UNASSIGNED_CASE_STATE
}

func (m *ExistingCaseDetail) GetInfectSrc() InfectionSource {
	if m != nil {
		return m.InfectSrc
	}
	return InfectionSource__UNASSIGNED_INF_SRC
}

type isExistingCaseDetail_Location interface {
	isExistingCaseDetail_Location()
}

type ExistingCaseDetail_NamedLocation struct {
	NamedLocation *NamedLocation `protobuf:"bytes,21,opt,name=named_location,json=namedLocation,proto3,oneof"`
}

type ExistingCaseDetail_GpsLocation struct {
	GpsLocation *GPSLocation `protobuf:"bytes,22,opt,name=gps_location,json=gpsLocation,proto3,oneof"`
}

func (*ExistingCaseDetail_NamedLocation) isExistingCaseDetail_Location() {}

func (*ExistingCaseDetail_GpsLocation) isExistingCaseDetail_Location() {}

func (m *ExistingCaseDetail) GetLocation() isExistingCaseDetail_Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *ExistingCaseDetail) GetNamedLocation() *NamedLocation {
	if x, ok := m.GetLocation().(*ExistingCaseDetail_NamedLocation); ok {
		return x.NamedLocation
	}
	return nil
}

func (m *ExistingCaseDetail) GetGpsLocation() *GPSLocation {
	if x, ok := m.GetLocation().(*ExistingCaseDetail_GpsLocation); ok {
		return x.GpsLocation
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExistingCaseDetail) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExistingCaseDetail_NamedLocation)(nil),
		(*ExistingCaseDetail_GpsLocation)(nil),
	}
}

// NewCaseDetail describes a new (not exists in DB) case
type NewCaseDetail struct {
	PatientId    string               `protobuf:"bytes,2,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	ReportedTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=reported_time,json=reportedTime,proto3" json:"reported_time,omitempty"`
	State        CaseState            `protobuf:"varint,11,opt,name=state,proto3,enum=covid19api.CaseState" json:"state,omitempty"`
	InfectSrc    InfectionSource      `protobuf:"varint,12,opt,name=infect_src,json=infectSrc,proto3,enum=covid19api.InfectionSource" json:"infect_src,omitempty"`
	// Types that are valid to be assigned to Location:
	//	*NewCaseDetail_NamedLocation
	//	*NewCaseDetail_GpsLocation
	Location             isNewCaseDetail_Location `protobuf_oneof:"location"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *NewCaseDetail) Reset()         { *m = NewCaseDetail{} }
func (m *NewCaseDetail) String() string { return proto.CompactTextString(m) }
func (*NewCaseDetail) ProtoMessage()    {}
func (*NewCaseDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_c687d06fd30aa5ff, []int{1}
}

func (m *NewCaseDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewCaseDetail.Unmarshal(m, b)
}
func (m *NewCaseDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewCaseDetail.Marshal(b, m, deterministic)
}
func (m *NewCaseDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewCaseDetail.Merge(m, src)
}
func (m *NewCaseDetail) XXX_Size() int {
	return xxx_messageInfo_NewCaseDetail.Size(m)
}
func (m *NewCaseDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_NewCaseDetail.DiscardUnknown(m)
}

var xxx_messageInfo_NewCaseDetail proto.InternalMessageInfo

func (m *NewCaseDetail) GetPatientId() string {
	if m != nil {
		return m.PatientId
	}
	return ""
}

func (m *NewCaseDetail) GetReportedTime() *timestamp.Timestamp {
	if m != nil {
		return m.ReportedTime
	}
	return nil
}

func (m *NewCaseDetail) GetState() CaseState {
	if m != nil {
		return m.State
	}
	return CaseState__UNASSIGNED_CASE_STATE
}

func (m *NewCaseDetail) GetInfectSrc() InfectionSource {
	if m != nil {
		return m.InfectSrc
	}
	return InfectionSource__UNASSIGNED_INF_SRC
}

type isNewCaseDetail_Location interface {
	isNewCaseDetail_Location()
}

type NewCaseDetail_NamedLocation struct {
	NamedLocation *NamedLocation `protobuf:"bytes,21,opt,name=named_location,json=namedLocation,proto3,oneof"`
}

type NewCaseDetail_GpsLocation struct {
	GpsLocation *GPSLocation `protobuf:"bytes,22,opt,name=gps_location,json=gpsLocation,proto3,oneof"`
}

func (*NewCaseDetail_NamedLocation) isNewCaseDetail_Location() {}

func (*NewCaseDetail_GpsLocation) isNewCaseDetail_Location() {}

func (m *NewCaseDetail) GetLocation() isNewCaseDetail_Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *NewCaseDetail) GetNamedLocation() *NamedLocation {
	if x, ok := m.GetLocation().(*NewCaseDetail_NamedLocation); ok {
		return x.NamedLocation
	}
	return nil
}

func (m *NewCaseDetail) GetGpsLocation() *GPSLocation {
	if x, ok := m.GetLocation().(*NewCaseDetail_GpsLocation); ok {
		return x.GpsLocation
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*NewCaseDetail) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*NewCaseDetail_NamedLocation)(nil),
		(*NewCaseDetail_GpsLocation)(nil),
	}
}

// GetCasesRequest describes request payload for querying cases match specific criteria
type GetCasesRequest struct {
	// Types that are valid to be assigned to Location:
	//	*GetCasesRequest_NamedLocation
	//	*GetCasesRequest_GpsLocation
	Location             isGetCasesRequest_Location `protobuf_oneof:"location"`
	States               []CaseState                `protobuf:"varint,11,rep,packed,name=states,proto3,enum=covid19api.CaseState" json:"states,omitempty"`
	StartTime            *timestamp.Timestamp       `protobuf:"bytes,21,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp       `protobuf:"bytes,22,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *GetCasesRequest) Reset()         { *m = GetCasesRequest{} }
func (m *GetCasesRequest) String() string { return proto.CompactTextString(m) }
func (*GetCasesRequest) ProtoMessage()    {}
func (*GetCasesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c687d06fd30aa5ff, []int{2}
}

func (m *GetCasesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCasesRequest.Unmarshal(m, b)
}
func (m *GetCasesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCasesRequest.Marshal(b, m, deterministic)
}
func (m *GetCasesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCasesRequest.Merge(m, src)
}
func (m *GetCasesRequest) XXX_Size() int {
	return xxx_messageInfo_GetCasesRequest.Size(m)
}
func (m *GetCasesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCasesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCasesRequest proto.InternalMessageInfo

type isGetCasesRequest_Location interface {
	isGetCasesRequest_Location()
}

type GetCasesRequest_NamedLocation struct {
	NamedLocation *NamedLocation `protobuf:"bytes,1,opt,name=named_location,json=namedLocation,proto3,oneof"`
}

type GetCasesRequest_GpsLocation struct {
	GpsLocation *GPSLocation `protobuf:"bytes,2,opt,name=gps_location,json=gpsLocation,proto3,oneof"`
}

func (*GetCasesRequest_NamedLocation) isGetCasesRequest_Location() {}

func (*GetCasesRequest_GpsLocation) isGetCasesRequest_Location() {}

func (m *GetCasesRequest) GetLocation() isGetCasesRequest_Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *GetCasesRequest) GetNamedLocation() *NamedLocation {
	if x, ok := m.GetLocation().(*GetCasesRequest_NamedLocation); ok {
		return x.NamedLocation
	}
	return nil
}

func (m *GetCasesRequest) GetGpsLocation() *GPSLocation {
	if x, ok := m.GetLocation().(*GetCasesRequest_GpsLocation); ok {
		return x.GpsLocation
	}
	return nil
}

func (m *GetCasesRequest) GetStates() []CaseState {
	if m != nil {
		return m.States
	}
	return nil
}

func (m *GetCasesRequest) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *GetCasesRequest) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GetCasesRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GetCasesRequest_NamedLocation)(nil),
		(*GetCasesRequest_GpsLocation)(nil),
	}
}

// GetCasesResponse describes response payload for matched cases
type GetCasesResponse struct {
	Pagination           *Pagination           `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Cases                []*ExistingCaseDetail `protobuf:"bytes,11,rep,name=cases,proto3" json:"cases,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetCasesResponse) Reset()         { *m = GetCasesResponse{} }
func (m *GetCasesResponse) String() string { return proto.CompactTextString(m) }
func (*GetCasesResponse) ProtoMessage()    {}
func (*GetCasesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c687d06fd30aa5ff, []int{3}
}

func (m *GetCasesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCasesResponse.Unmarshal(m, b)
}
func (m *GetCasesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCasesResponse.Marshal(b, m, deterministic)
}
func (m *GetCasesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCasesResponse.Merge(m, src)
}
func (m *GetCasesResponse) XXX_Size() int {
	return xxx_messageInfo_GetCasesResponse.Size(m)
}
func (m *GetCasesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCasesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCasesResponse proto.InternalMessageInfo

func (m *GetCasesResponse) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *GetCasesResponse) GetCases() []*ExistingCaseDetail {
	if m != nil {
		return m.Cases
	}
	return nil
}

// AddCasesRequest describes request payload for adding cases
type AddCasesRequest struct {
	Cases                []*NewCaseDetail `protobuf:"bytes,1,rep,name=cases,proto3" json:"cases,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AddCasesRequest) Reset()         { *m = AddCasesRequest{} }
func (m *AddCasesRequest) String() string { return proto.CompactTextString(m) }
func (*AddCasesRequest) ProtoMessage()    {}
func (*AddCasesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c687d06fd30aa5ff, []int{4}
}

func (m *AddCasesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddCasesRequest.Unmarshal(m, b)
}
func (m *AddCasesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddCasesRequest.Marshal(b, m, deterministic)
}
func (m *AddCasesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddCasesRequest.Merge(m, src)
}
func (m *AddCasesRequest) XXX_Size() int {
	return xxx_messageInfo_AddCasesRequest.Size(m)
}
func (m *AddCasesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddCasesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddCasesRequest proto.InternalMessageInfo

func (m *AddCasesRequest) GetCases() []*NewCaseDetail {
	if m != nil {
		return m.Cases
	}
	return nil
}

// PutCasesRequest describes request payload for updating existing cases
type PutCasesRequest struct {
	Cases                []*ExistingCaseDetail `protobuf:"bytes,1,rep,name=cases,proto3" json:"cases,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PutCasesRequest) Reset()         { *m = PutCasesRequest{} }
func (m *PutCasesRequest) String() string { return proto.CompactTextString(m) }
func (*PutCasesRequest) ProtoMessage()    {}
func (*PutCasesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c687d06fd30aa5ff, []int{5}
}

func (m *PutCasesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutCasesRequest.Unmarshal(m, b)
}
func (m *PutCasesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutCasesRequest.Marshal(b, m, deterministic)
}
func (m *PutCasesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutCasesRequest.Merge(m, src)
}
func (m *PutCasesRequest) XXX_Size() int {
	return xxx_messageInfo_PutCasesRequest.Size(m)
}
func (m *PutCasesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutCasesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutCasesRequest proto.InternalMessageInfo

func (m *PutCasesRequest) GetCases() []*ExistingCaseDetail {
	if m != nil {
		return m.Cases
	}
	return nil
}

// DelCasesRequest describes request payload for deleting existing cases
type DelCasesRequest struct {
	CaseUuids            []string `protobuf:"bytes,1,rep,name=case_uuids,json=caseUuids,proto3" json:"case_uuids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelCasesRequest) Reset()         { *m = DelCasesRequest{} }
func (m *DelCasesRequest) String() string { return proto.CompactTextString(m) }
func (*DelCasesRequest) ProtoMessage()    {}
func (*DelCasesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c687d06fd30aa5ff, []int{6}
}

func (m *DelCasesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelCasesRequest.Unmarshal(m, b)
}
func (m *DelCasesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelCasesRequest.Marshal(b, m, deterministic)
}
func (m *DelCasesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelCasesRequest.Merge(m, src)
}
func (m *DelCasesRequest) XXX_Size() int {
	return xxx_messageInfo_DelCasesRequest.Size(m)
}
func (m *DelCasesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DelCasesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DelCasesRequest proto.InternalMessageInfo

func (m *DelCasesRequest) GetCaseUuids() []string {
	if m != nil {
		return m.CaseUuids
	}
	return nil
}

func init() {
	proto.RegisterEnum("covid19api.CaseState", CaseState_name, CaseState_value)
	proto.RegisterEnum("covid19api.InfectionSource", InfectionSource_name, InfectionSource_value)
	proto.RegisterType((*ExistingCaseDetail)(nil), "covid19api.ExistingCaseDetail")
	proto.RegisterType((*NewCaseDetail)(nil), "covid19api.NewCaseDetail")
	proto.RegisterType((*GetCasesRequest)(nil), "covid19api.GetCasesRequest")
	proto.RegisterType((*GetCasesResponse)(nil), "covid19api.GetCasesResponse")
	proto.RegisterType((*AddCasesRequest)(nil), "covid19api.AddCasesRequest")
	proto.RegisterType((*PutCasesRequest)(nil), "covid19api.PutCasesRequest")
	proto.RegisterType((*DelCasesRequest)(nil), "covid19api.DelCasesRequest")
}

func init() {
	proto.RegisterFile("case.def.proto", fileDescriptor_c687d06fd30aa5ff)
}

var fileDescriptor_c687d06fd30aa5ff = []byte{
	// 692 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x54, 0x5d, 0x6f, 0xd3, 0x4a,
	0x10, 0xad, 0x13, 0xb5, 0x8d, 0xc7, 0x49, 0xe3, 0xbb, 0x57, 0x49, 0x73, 0x53, 0xf5, 0x12, 0xf5,
	0x85, 0xaa, 0x08, 0x17, 0x4a, 0x41, 0x2a, 0x42, 0x42, 0x49, 0xec, 0xa6, 0x16, 0xad, 0x13, 0xd9,
	0x09, 0xbc, 0x61, 0xb9, 0xf6, 0xc6, 0x5a, 0x29, 0xf1, 0x1a, 0xef, 0x06, 0x78, 0xe4, 0x91, 0x9f,
	0xc5, 0xaf, 0xe1, 0x77, 0xa0, 0xb5, 0xe3, 0xc6, 0xe1, 0x43, 0x05, 0x21, 0xde, 0x78, 0xdc, 0x99,
	0x33, 0xc7, 0x67, 0xe6, 0x8c, 0x07, 0x76, 0x7c, 0x8f, 0x61, 0x2d, 0xc0, 0x53, 0x2d, 0x4e, 0x28,
	0xa7, 0x08, 0x7c, 0xfa, 0x96, 0x04, 0x0f, 0xcf, 0xbc, 0x98, 0xb4, 0xef, 0x84, 0x94, 0x86, 0x33,
	0x7c, 0x9c, 0x66, 0xae, 0x17, 0xd3, 0x63, 0x4e, 0xe6, 0x98, 0x71, 0x6f, 0x1e, 0x67, 0xe0, 0x76,
	0xd5, 0xa7, 0xf3, 0x39, 0x8d, 0x96, 0xaf, 0x5a, 0x88, 0xe9, 0x8a, 0xe9, 0xe0, 0x63, 0x19, 0x90,
	0xf1, 0x9e, 0x30, 0x4e, 0xa2, 0xb0, 0xef, 0x31, 0xac, 0x63, 0xee, 0x91, 0x19, 0xda, 0x03, 0x59,
	0x7c, 0xd2, 0x5d, 0x2c, 0x48, 0xd0, 0x92, 0x3a, 0xd2, 0xa1, 0x6c, 0x57, 0x44, 0x60, 0xb2, 0x20,
	0x01, 0xda, 0x07, 0x88, 0x3d, 0x4e, 0x70, 0xc4, 0x5d, 0x12, 0xb4, 0x4a, 0x69, 0x56, 0x5e, 0x46,
	0xcc, 0x00, 0x3d, 0x87, 0x5a, 0x82, 0x63, 0x9a, 0x70, 0x1c, 0xb8, 0x42, 0x4b, 0xab, 0xdc, 0x91,
	0x0e, 0x95, 0x93, 0xb6, 0x96, 0x09, 0xd5, 0x72, 0xa1, 0xda, 0x38, 0x17, 0x6a, 0x57, 0xf3, 0x02,
	0x11, 0x42, 0xf7, 0x60, 0x93, 0x71, 0x8f, 0xe3, 0x96, 0xd2, 0x91, 0x0e, 0x77, 0x4e, 0x1a, 0xda,
	0xaa, 0x5b, 0x4d, 0x68, 0x74, 0x44, 0xd2, 0xce, 0x30, 0xe8, 0x29, 0x00, 0x89, 0xa6, 0xd8, 0xe7,
	0x2e, 0x4b, 0xfc, 0x56, 0x35, 0xad, 0xd8, 0x2b, 0x56, 0x98, 0x69, 0x96, 0xd0, 0xc8, 0xa1, 0x8b,
	0xc4, 0xc7, 0xb6, 0x9c, 0xc1, 0x9d, 0xc4, 0x47, 0x3d, 0xd8, 0x89, 0xbc, 0x39, 0x0e, 0xdc, 0x19,
	0xf5, 0x3d, 0x01, 0x69, 0x35, 0x52, 0xa9, 0xff, 0x15, 0xeb, 0x2d, 0x81, 0xb8, 0x5c, 0x02, 0x2e,
	0x36, 0xec, 0x5a, 0x54, 0x0c, 0xa0, 0x67, 0x50, 0x0d, 0x63, 0xb6, 0x62, 0x68, 0xa6, 0x0c, 0xbb,
	0x45, 0x86, 0xc1, 0xc8, 0x29, 0xd4, 0x2b, 0x61, 0xcc, 0xf2, 0x67, 0x0f, 0xa0, 0x92, 0x57, 0x1e,
	0x7c, 0x2e, 0x41, 0xcd, 0xc2, 0xef, 0x0a, 0x2e, 0xfc, 0x1d, 0xf4, 0x1f, 0x19, 0xf4, 0xa7, 0x12,
	0xd4, 0x07, 0x98, 0x8b, 0x0e, 0x99, 0x8d, 0xdf, 0x2c, 0x30, 0xe3, 0xdf, 0x51, 0x28, 0xfd, 0xb6,
	0xc2, 0xd2, 0xaf, 0x28, 0x44, 0xf7, 0x61, 0x2b, 0x1d, 0x34, 0x6b, 0x29, 0x9d, 0xf2, 0x8f, 0xdd,
	0x58, 0x82, 0xd0, 0x19, 0x00, 0xe3, 0x5e, 0xc2, 0x33, 0xe7, 0x1b, 0xb7, 0x3a, 0x2f, 0xa7, 0xe8,
	0xd4, 0xf6, 0xc7, 0x50, 0xc1, 0xd1, 0x72, 0x65, 0x9a, 0xb7, 0x16, 0x6e, 0xe3, 0x28, 0xdd, 0x96,
	0xb5, 0x11, 0x7e, 0x90, 0x40, 0x5d, 0x8d, 0x90, 0xc5, 0x34, 0x62, 0x18, 0x3d, 0x11, 0xeb, 0x1a,
	0x92, 0xa8, 0x38, 0xbf, 0x66, 0xb1, 0x8b, 0xd1, 0x4d, 0xd6, 0x2e, 0x20, 0xd1, 0x29, 0x6c, 0x8a,
	0xdb, 0x92, 0x35, 0xae, 0x9c, 0xfc, 0x5f, 0x2c, 0xf9, 0xf6, 0x36, 0xd9, 0x19, 0xf8, 0xa0, 0x07,
	0xf5, 0x6e, 0x10, 0xac, 0x99, 0x78, 0x9c, 0x13, 0x49, 0x29, 0xd1, 0xba, 0x77, 0xc5, 0x3f, 0x2b,
	0xe7, 0x18, 0x40, 0x7d, 0xb4, 0x58, 0x5f, 0x84, 0xd3, 0x75, 0x8e, 0x9f, 0x14, 0xf3, 0x00, 0xea,
	0x3a, 0x9e, 0xad, 0x11, 0xed, 0x03, 0xdc, 0x9c, 0xd0, 0x8c, 0x4d, 0xb6, 0xe5, 0xfc, 0x86, 0xb2,
	0xa3, 0x08, 0xe4, 0x1b, 0x53, 0x51, 0x1b, 0x9a, 0xee, 0xc4, 0xea, 0x3a, 0x8e, 0x39, 0xb0, 0x0c,
	0xdd, 0xed, 0x77, 0x1d, 0xc3, 0x75, 0xc6, 0xdd, 0xb1, 0xa1, 0x6e, 0x20, 0x05, 0xb6, 0x27, 0xd6,
	0x0b, 0x6b, 0xf8, 0xca, 0x52, 0x15, 0x54, 0x07, 0x65, 0x62, 0xe9, 0xa6, 0xd3, 0xbf, 0x1c, 0x3a,
	0x86, 0xae, 0x56, 0x51, 0x0d, 0xe4, 0xfe, 0xd0, 0x3a, 0x37, 0xed, 0x2b, 0x43, 0x57, 0x1b, 0xe2,
	0x69, 0x1b, 0xfd, 0xe1, 0x4b, 0xc3, 0x36, 0x74, 0xb5, 0x89, 0x64, 0xd8, 0xd4, 0x8d, 0xee, 0xf8,
	0x42, 0xdd, 0x3d, 0x7a, 0x0d, 0xf5, 0xaf, 0x7e, 0x50, 0xb4, 0x0b, 0xff, 0x16, 0xbf, 0x6a, 0x5a,
	0xe7, 0xae, 0x63, 0xf7, 0xd5, 0x0d, 0x54, 0x85, 0x8a, 0x3e, 0xbc, 0x32, 0x9c, 0xb1, 0xd9, 0x57,
	0x25, 0xf4, 0x0f, 0xd4, 0x4c, 0x6b, 0x7c, 0xe9, 0x9a, 0x57, 0xa3, 0xa1, 0x3d, 0x36, 0x74, 0xb5,
	0x24, 0x42, 0xce, 0x85, 0x39, 0x5a, 0x85, 0xca, 0xbd, 0xbb, 0x50, 0xa7, 0x49, 0x98, 0x4f, 0x4b,
	0xf3, 0x62, 0xd2, 0x43, 0xa2, 0x41, 0x33, 0xe2, 0x38, 0x99, 0x7a, 0x3e, 0x1e, 0x89, 0xd5, 0x1a,
	0x49, 0xd7, 0x5b, 0xe9, 0x8e, 0x3d, 0xfa, 0x12, 0x00, 0x00, 0xff, 0xff, 0x47, 0x18, 0x15, 0x28,
	0xd4, 0x06, 0x00, 0x00,
}