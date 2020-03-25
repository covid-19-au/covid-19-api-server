// Code generated by protoc-gen-go. DO NOT EDIT.
// source: admin.svc.proto

package org_covid19_api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() {
	proto.RegisterFile("admin.svc.proto", fileDescriptor_bd637287d1829a88)
}

var fileDescriptor_bd637287d1829a88 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x11, 0x41, 0x34, 0x87, 0x2d, 0xe6, 0xe0, 0xa1, 0x82, 0x88, 0x0f, 0x30, 0x8b, 0x7a,
	0xf2, 0xd8, 0x5a, 0x45, 0xc1, 0x43, 0x59, 0xc1, 0x7b, 0xb6, 0x33, 0xad, 0x81, 0xac, 0x13, 0x93,
	0xb4, 0xb2, 0xcf, 0xe8, 0x4b, 0x49, 0xb7, 0xad, 0xca, 0x6a, 0xd9, 0x5e, 0xff, 0x6f, 0xf8, 0xfe,
	0x99, 0x44, 0x44, 0x0a, 0x57, 0xfa, 0x0d, 0x7c, 0x53, 0x80, 0x75, 0x1c, 0x58, 0x46, 0xec, 0x2a,
	0x28, 0xb8, 0xd1, 0x78, 0x79, 0x03, 0xca, 0xea, 0xf8, 0xb4, 0x62, 0xae, 0x0c, 0xcd, 0x37, 0x78,
	0x59, 0x97, 0x73, 0x5a, 0xd9, 0xb0, 0xee, 0xa6, 0xe3, 0xb3, 0x6d, 0xf8, 0xe1, 0x94, 0xb5, 0xe4,
	0x7c, 0xcf, 0xa3, 0x42, 0x79, 0xf2, 0x80, 0x54, 0xf6, 0xc1, 0x71, 0x69, 0x74, 0xf5, 0x1a, 0x7e,
	0x45, 0x57, 0x9f, 0xfb, 0xe2, 0x28, 0x53, 0x41, 0x25, 0xed, 0x26, 0x32, 0x13, 0xb3, 0x47, 0xff,
	0x4c, 0xae, 0xd1, 0x05, 0x2d, 0x48, 0xe1, 0x5a, 0x9e, 0x40, 0x57, 0x02, 0x43, 0x09, 0xdc, 0xb5,
	0x1b, 0xc4, 0xf1, 0x9f, 0x3c, 0x65, 0x36, 0x2f, 0xca, 0xd4, 0x24, 0x1f, 0xc4, 0x61, 0x82, 0x78,
	0xdb, 0x96, 0xcb, 0x73, 0xd8, 0x3a, 0x09, 0x06, 0xb4, 0xa0, 0xf7, 0x9a, 0x7c, 0xd8, 0x65, 0xca,
	0xeb, 0x30, 0x66, 0x1a, 0xd0, 0x44, 0x53, 0x46, 0x66, 0xcc, 0x34, 0xa0, 0x29, 0xa6, 0x27, 0x21,
	0x12, 0xc4, 0xfb, 0xee, 0x25, 0xe5, 0xc5, 0x7f, 0xf7, 0xf5, 0x70, 0xa2, 0x2d, 0x23, 0x33, 0x6e,
	0xfb, 0x81, 0x13, 0x6c, 0xa9, 0x4c, 0x67, 0xdf, 0x9f, 0x99, 0xb7, 0x38, 0xdf, 0x5b, 0x1e, 0x6c,
	0xe6, 0xae, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x9f, 0x00, 0xb2, 0x6d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DataAdminClient is the client API for DataAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataAdminClient interface {
	IsServiceReady(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
	AddCases(ctx context.Context, in *AddCasesRequest, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
	PutCases(ctx context.Context, in *PutCasesRequest, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
	DelCases(ctx context.Context, in *DelCasesRequest, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
	AddFlights(ctx context.Context, in *AddFlightsRequest, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
	DelFlights(ctx context.Context, in *DelFlightsRequest, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
}

type dataAdminClient struct {
	cc grpc.ClientConnInterface
}

func NewDataAdminClient(cc grpc.ClientConnInterface) DataAdminClient {
	return &dataAdminClient{cc}
}

func (c *dataAdminClient) IsServiceReady(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/org.covid19.api.DataAdmin/IsServiceReady", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataAdminClient) AddCases(ctx context.Context, in *AddCasesRequest, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/org.covid19.api.DataAdmin/AddCases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataAdminClient) PutCases(ctx context.Context, in *PutCasesRequest, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/org.covid19.api.DataAdmin/PutCases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataAdminClient) DelCases(ctx context.Context, in *DelCasesRequest, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/org.covid19.api.DataAdmin/DelCases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataAdminClient) AddFlights(ctx context.Context, in *AddFlightsRequest, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/org.covid19.api.DataAdmin/AddFlights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataAdminClient) DelFlights(ctx context.Context, in *DelFlightsRequest, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/org.covid19.api.DataAdmin/DelFlights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataAdminServer is the server API for DataAdmin service.
type DataAdminServer interface {
	IsServiceReady(context.Context, *empty.Empty) (*wrappers.BoolValue, error)
	AddCases(context.Context, *AddCasesRequest) (*wrappers.BoolValue, error)
	PutCases(context.Context, *PutCasesRequest) (*wrappers.BoolValue, error)
	DelCases(context.Context, *DelCasesRequest) (*wrappers.BoolValue, error)
	AddFlights(context.Context, *AddFlightsRequest) (*wrappers.BoolValue, error)
	DelFlights(context.Context, *DelFlightsRequest) (*wrappers.BoolValue, error)
}

// UnimplementedDataAdminServer can be embedded to have forward compatible implementations.
type UnimplementedDataAdminServer struct {
}

func (*UnimplementedDataAdminServer) IsServiceReady(ctx context.Context, req *empty.Empty) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsServiceReady not implemented")
}
func (*UnimplementedDataAdminServer) AddCases(ctx context.Context, req *AddCasesRequest) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCases not implemented")
}
func (*UnimplementedDataAdminServer) PutCases(ctx context.Context, req *PutCasesRequest) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutCases not implemented")
}
func (*UnimplementedDataAdminServer) DelCases(ctx context.Context, req *DelCasesRequest) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelCases not implemented")
}
func (*UnimplementedDataAdminServer) AddFlights(ctx context.Context, req *AddFlightsRequest) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFlights not implemented")
}
func (*UnimplementedDataAdminServer) DelFlights(ctx context.Context, req *DelFlightsRequest) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelFlights not implemented")
}

func RegisterDataAdminServer(s *grpc.Server, srv DataAdminServer) {
	s.RegisterService(&_DataAdmin_serviceDesc, srv)
}

func _DataAdmin_IsServiceReady_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataAdminServer).IsServiceReady(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.covid19.api.DataAdmin/IsServiceReady",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataAdminServer).IsServiceReady(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataAdmin_AddCases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataAdminServer).AddCases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.covid19.api.DataAdmin/AddCases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataAdminServer).AddCases(ctx, req.(*AddCasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataAdmin_PutCases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutCasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataAdminServer).PutCases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.covid19.api.DataAdmin/PutCases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataAdminServer).PutCases(ctx, req.(*PutCasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataAdmin_DelCases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelCasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataAdminServer).DelCases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.covid19.api.DataAdmin/DelCases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataAdminServer).DelCases(ctx, req.(*DelCasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataAdmin_AddFlights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFlightsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataAdminServer).AddFlights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.covid19.api.DataAdmin/AddFlights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataAdminServer).AddFlights(ctx, req.(*AddFlightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataAdmin_DelFlights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelFlightsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataAdminServer).DelFlights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.covid19.api.DataAdmin/DelFlights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataAdminServer).DelFlights(ctx, req.(*DelFlightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "org.covid19.api.DataAdmin",
	HandlerType: (*DataAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsServiceReady",
			Handler:    _DataAdmin_IsServiceReady_Handler,
		},
		{
			MethodName: "AddCases",
			Handler:    _DataAdmin_AddCases_Handler,
		},
		{
			MethodName: "PutCases",
			Handler:    _DataAdmin_PutCases_Handler,
		},
		{
			MethodName: "DelCases",
			Handler:    _DataAdmin_DelCases_Handler,
		},
		{
			MethodName: "AddFlights",
			Handler:    _DataAdmin_AddFlights_Handler,
		},
		{
			MethodName: "DelFlights",
			Handler:    _DataAdmin_DelFlights_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.svc.proto",
}
