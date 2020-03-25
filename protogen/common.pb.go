// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package org_covid19_api

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

// Pagination handles multiple-page results
type Pagination struct {
	TotalResults         uint32   `protobuf:"varint,1,opt,name=total_results,json=totalResults,proto3" json:"total_results,omitempty"`
	TotalPages           uint32   `protobuf:"varint,2,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	CurrentPage          uint32   `protobuf:"varint,3,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	NumPerPage           uint32   `protobuf:"varint,4,opt,name=num_per_page,json=numPerPage,proto3" json:"num_per_page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pagination) Reset()         { *m = Pagination{} }
func (m *Pagination) String() string { return proto.CompactTextString(m) }
func (*Pagination) ProtoMessage()    {}
func (*Pagination) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *Pagination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pagination.Unmarshal(m, b)
}
func (m *Pagination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pagination.Marshal(b, m, deterministic)
}
func (m *Pagination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pagination.Merge(m, src)
}
func (m *Pagination) XXX_Size() int {
	return xxx_messageInfo_Pagination.Size(m)
}
func (m *Pagination) XXX_DiscardUnknown() {
	xxx_messageInfo_Pagination.DiscardUnknown(m)
}

var xxx_messageInfo_Pagination proto.InternalMessageInfo

func (m *Pagination) GetTotalResults() uint32 {
	if m != nil {
		return m.TotalResults
	}
	return 0
}

func (m *Pagination) GetTotalPages() uint32 {
	if m != nil {
		return m.TotalPages
	}
	return 0
}

func (m *Pagination) GetCurrentPage() uint32 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func (m *Pagination) GetNumPerPage() uint32 {
	if m != nil {
		return m.NumPerPage
	}
	return 0
}

func init() {
	proto.RegisterType((*Pagination)(nil), "org.covid19.api.Pagination")
}

func init() {
	proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206)
}

var fileDescriptor_555bd8c177793206 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0xcf, 0x41, 0xaa, 0xc2, 0x30,
	0x10, 0xc6, 0x71, 0xf2, 0x9e, 0xb8, 0x98, 0xa6, 0x14, 0xb2, 0xea, 0xce, 0xaa, 0x1b, 0x57, 0x05,
	0x71, 0xe5, 0xb6, 0x5e, 0x20, 0xf4, 0x02, 0x25, 0xd6, 0x50, 0x02, 0x4d, 0x26, 0xa4, 0x89, 0x57,
	0xf1, 0xba, 0x92, 0x89, 0xdb, 0x3f, 0x3f, 0x86, 0x6f, 0x80, 0xcf, 0x68, 0x2d, 0xba, 0xde, 0x07,
	0x8c, 0x28, 0x1a, 0x0c, 0x4b, 0x3f, 0xe3, 0xdb, 0xbc, 0xae, 0xf7, 0x5e, 0x79, 0x73, 0xfa, 0x30,
	0x00, 0xa9, 0x16, 0xe3, 0x54, 0x34, 0xe8, 0xc4, 0x19, 0xea, 0x88, 0x51, 0xad, 0x53, 0xd0, 0x5b,
	0x5a, 0xe3, 0xd6, 0xb2, 0x8e, 0x5d, 0xea, 0x91, 0x53, 0x1c, 0x4b, 0x13, 0x07, 0xa8, 0x0a, 0xf2,
	0x6a, 0xd1, 0x5b, 0xfb, 0x47, 0x04, 0x28, 0xc9, 0x5c, 0xc4, 0x11, 0xf8, 0x9c, 0x42, 0xd0, 0x2e,
	0x12, 0x69, 0xff, 0x49, 0x54, 0xbf, 0x96, 0x8d, 0xe8, 0x80, 0xbb, 0x64, 0x27, 0xaf, 0x43, 0x21,
	0xbb, 0x72, 0xc4, 0x25, 0x2b, 0x75, 0xc8, 0x62, 0x68, 0x86, 0xea, 0x41, 0xd3, 0x65, 0x5e, 0x2e,
	0xd9, 0x73, 0x4f, 0x2f, 0xdc, 0xbe, 0x01, 0x00, 0x00, 0xff, 0xff, 0x77, 0x5b, 0x54, 0x26, 0xd2,
	0x00, 0x00, 0x00,
}
