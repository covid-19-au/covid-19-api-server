// Code generated by protoc-gen-go. DO NOT EDIT.
// source: admin.svc.proto

package covid19api

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
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x41, 0x48, 0x08, 0x3c, 0xa4, 0xc8, 0x03, 0x43, 0x2a, 0x18, 0x60, 0xbf, 0x0a, 0x98,
	0x18, 0x13, 0xd2, 0x4a, 0xdd, 0xa2, 0x22, 0xb1, 0x5f, 0xe3, 0x4b, 0xb0, 0xe4, 0x62, 0x63, 0x3b,
	0x41, 0x7d, 0x39, 0x9e, 0x0d, 0x25, 0x8d, 0x89, 0x54, 0x22, 0x91, 0x31, 0xdf, 0x9f, 0xfb, 0x7e,
	0x9f, 0xcd, 0x66, 0x28, 0x76, 0xf2, 0x03, 0x5c, 0x53, 0x80, 0xb1, 0xda, 0x6b, 0xce, 0x0a, 0xdd,
	0x48, 0xf1, 0xf0, 0x8c, 0x46, 0xc6, 0xf3, 0x4a, 0xeb, 0x4a, 0xd1, 0xa2, 0x4b, 0xb6, 0x75, 0xb9,
	0xa0, 0x9d, 0xf1, 0xfb, 0xc3, 0x8f, 0xf1, 0xed, 0x71, 0xf8, 0x65, 0xd1, 0x18, 0xb2, 0xae, 0xcf,
	0xa3, 0x02, 0x1d, 0x81, 0xa0, 0xb2, 0xff, 0xbe, 0x2a, 0x95, 0xac, 0xde, 0xfd, 0x40, 0x1e, 0xbf,
	0xcf, 0xd8, 0x65, 0x86, 0x1e, 0x93, 0xf6, 0x08, 0x7c, 0xc5, 0xa2, 0xb5, 0x7b, 0x25, 0xdb, 0xc8,
	0x82, 0x36, 0x84, 0x62, 0xcf, 0xaf, 0xe1, 0x50, 0x01, 0xa1, 0x02, 0x96, 0x6d, 0x7f, 0x1c, 0xff,
	0xe1, 0xa9, 0xd6, 0xea, 0x0d, 0x55, 0x4d, 0x77, 0x27, 0x7c, 0xc9, 0x2e, 0x12, 0x21, 0x5e, 0xd0,
	0x91, 0xe3, 0x73, 0x18, 0xb6, 0x81, 0x40, 0x37, 0xf4, 0x59, 0x93, 0xf3, 0xff, 0x6b, 0xf2, 0xda,
	0x8f, 0x68, 0x02, 0x9d, 0xac, 0xc9, 0x48, 0x8d, 0x68, 0x02, 0x9d, 0xa6, 0x59, 0x33, 0x96, 0x08,
	0xb1, 0xea, 0x6e, 0xd0, 0xf1, 0x9b, 0xa3, 0xb5, 0x7a, 0x3e, 0x59, 0x95, 0x91, 0x1a, 0x55, 0x0d,
	0x7c, 0x92, 0x2a, 0xbd, 0x67, 0x33, 0x6d, 0xab, 0x60, 0x00, 0x34, 0x32, 0x8d, 0x7e, 0x1f, 0x34,
	0x6f, 0x27, 0xf2, 0xd3, 0xed, 0x79, 0x37, 0xfa, 0xf4, 0x13, 0x00, 0x00, 0xff, 0xff, 0x57, 0x9f,
	0x4d, 0x2d, 0x6a, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/covid19api.DataAdmin/IsServiceReady", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataAdminClient) AddCases(ctx context.Context, in *AddCasesRequest, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/covid19api.DataAdmin/AddCases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataAdminClient) PutCases(ctx context.Context, in *PutCasesRequest, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/covid19api.DataAdmin/PutCases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataAdminClient) DelCases(ctx context.Context, in *DelCasesRequest, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/covid19api.DataAdmin/DelCases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataAdminClient) AddFlights(ctx context.Context, in *AddFlightsRequest, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/covid19api.DataAdmin/AddFlights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataAdminClient) DelFlights(ctx context.Context, in *DelFlightsRequest, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/covid19api.DataAdmin/DelFlights", in, out, opts...)
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
		FullMethod: "/covid19api.DataAdmin/IsServiceReady",
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
		FullMethod: "/covid19api.DataAdmin/AddCases",
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
		FullMethod: "/covid19api.DataAdmin/PutCases",
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
		FullMethod: "/covid19api.DataAdmin/DelCases",
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
		FullMethod: "/covid19api.DataAdmin/AddFlights",
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
		FullMethod: "/covid19api.DataAdmin/DelFlights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataAdminServer).DelFlights(ctx, req.(*DelFlightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "covid19api.DataAdmin",
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