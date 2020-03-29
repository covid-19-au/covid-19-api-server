// Code generated by protoc-gen-go. DO NOT EDIT.
// source: geo.def.proto

package covid19api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// NamedLocation describes location information with text-based names
type NamedLocation struct {
	Country              string   `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	Region               string   `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	City                 string   `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	PostalCode           string   `protobuf:"bytes,4,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NamedLocation) Reset()         { *m = NamedLocation{} }
func (m *NamedLocation) String() string { return proto.CompactTextString(m) }
func (*NamedLocation) ProtoMessage()    {}
func (*NamedLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1446070098fba2b, []int{0}
}

func (m *NamedLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamedLocation.Unmarshal(m, b)
}
func (m *NamedLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamedLocation.Marshal(b, m, deterministic)
}
func (m *NamedLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamedLocation.Merge(m, src)
}
func (m *NamedLocation) XXX_Size() int {
	return xxx_messageInfo_NamedLocation.Size(m)
}
func (m *NamedLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_NamedLocation.DiscardUnknown(m)
}

var xxx_messageInfo_NamedLocation proto.InternalMessageInfo

func (m *NamedLocation) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *NamedLocation) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *NamedLocation) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *NamedLocation) GetPostalCode() string {
	if m != nil {
		return m.PostalCode
	}
	return ""
}

// GPSLocation describes location information with GPS coordinates
type GPSLocation struct {
	GeoLat               float64  `protobuf:"fixed64,1,opt,name=geo_lat,json=geoLat,proto3" json:"geo_lat,omitempty"`
	GeoLng               float64  `protobuf:"fixed64,2,opt,name=geo_lng,json=geoLng,proto3" json:"geo_lng,omitempty"`
	PrecisionRad         uint64   `protobuf:"varint,3,opt,name=precision_rad,json=precisionRad,proto3" json:"precision_rad,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GPSLocation) Reset()         { *m = GPSLocation{} }
func (m *GPSLocation) String() string { return proto.CompactTextString(m) }
func (*GPSLocation) ProtoMessage()    {}
func (*GPSLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1446070098fba2b, []int{1}
}

func (m *GPSLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GPSLocation.Unmarshal(m, b)
}
func (m *GPSLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GPSLocation.Marshal(b, m, deterministic)
}
func (m *GPSLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GPSLocation.Merge(m, src)
}
func (m *GPSLocation) XXX_Size() int {
	return xxx_messageInfo_GPSLocation.Size(m)
}
func (m *GPSLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_GPSLocation.DiscardUnknown(m)
}

var xxx_messageInfo_GPSLocation proto.InternalMessageInfo

func (m *GPSLocation) GetGeoLat() float64 {
	if m != nil {
		return m.GeoLat
	}
	return 0
}

func (m *GPSLocation) GetGeoLng() float64 {
	if m != nil {
		return m.GeoLng
	}
	return 0
}

func (m *GPSLocation) GetPrecisionRad() uint64 {
	if m != nil {
		return m.PrecisionRad
	}
	return 0
}

func init() {
	proto.RegisterType((*NamedLocation)(nil), "covid19api.NamedLocation")
	proto.RegisterType((*GPSLocation)(nil), "covid19api.GPSLocation")
}

func init() {
	proto.RegisterFile("geo.def.proto", fileDescriptor_b1446070098fba2b)
}

var fileDescriptor_b1446070098fba2b = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x59, 0x5d, 0xb6, 0x38, 0x75, 0x11, 0x73, 0xd0, 0xdc, 0x94, 0x0a, 0xe2, 0x69, 0x41,
	0x3c, 0x79, 0xad, 0x87, 0x22, 0x14, 0x59, 0xe2, 0x0f, 0x58, 0xc6, 0x64, 0x1a, 0x02, 0x35, 0x13,
	0xd2, 0x58, 0xd8, 0x7f, 0x2f, 0x8e, 0xdb, 0xbd, 0xcd, 0x7b, 0xdf, 0xe1, 0x63, 0x1e, 0xb4, 0x9e,
	0xb8, 0x73, 0xb4, 0xeb, 0x52, 0xe6, 0xc2, 0x0a, 0x2c, 0x1f, 0x83, 0x7b, 0x7e, 0xc5, 0x14, 0x56,
	0x47, 0x68, 0x3f, 0xf0, 0x9b, 0xdc, 0x96, 0x2d, 0x96, 0xc0, 0x51, 0x69, 0x58, 0x58, 0xfe, 0x89,
	0x25, 0x8f, 0xba, 0xba, 0xaf, 0x9e, 0x2e, 0xcc, 0x29, 0xaa, 0x1b, 0x68, 0x32, 0xf9, 0xc0, 0x51,
	0x9f, 0x09, 0x98, 0x92, 0x52, 0x50, 0xdb, 0x50, 0x46, 0x7d, 0x2e, 0xad, 0xdc, 0xea, 0x0e, 0x96,
	0x89, 0x0f, 0x05, 0xf7, 0x83, 0x65, 0x47, 0xba, 0x16, 0x04, 0xff, 0xd5, 0x1b, 0x3b, 0x5a, 0x39,
	0x58, 0x6e, 0xfa, 0xcf, 0xd9, 0x7a, 0x0b, 0x0b, 0x4f, 0x3c, 0xec, 0xb1, 0x88, 0xb5, 0x32, 0x8d,
	0x27, 0xde, 0x62, 0x99, 0x41, 0xf4, 0x62, 0x9d, 0x40, 0xf4, 0xea, 0x01, 0xda, 0x94, 0xc9, 0x86,
	0x43, 0xe0, 0x38, 0x64, 0x74, 0xa2, 0xaf, 0xcd, 0xe5, 0x5c, 0x1a, 0x74, 0xeb, 0x47, 0xb8, 0xe2,
	0xec, 0xbb, 0xe9, 0xdf, 0x0e, 0x53, 0x58, 0x5f, 0x6f, 0x88, 0xdf, 0x63, 0xa1, 0xbc, 0x43, 0x4b,
	0xfd, 0xdf, 0x1e, 0x7d, 0xf5, 0xd5, 0xc8, 0x30, 0x2f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x33,
	0x90, 0x32, 0x56, 0x29, 0x01, 0x00, 0x00,
}