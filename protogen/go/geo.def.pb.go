// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: geo.def.proto

package covid19api

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Location describes location information with text-based names
type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country         string  `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`                                            // Country code (in ISO 3166-1 alpha-2 format)
	Region          string  `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`                                              // Region code (state or province, in ISO 3166-2 format)
	City            string  `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`                                                  // City name (in full format)
	PostalCode      string  `protobuf:"bytes,4,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`                    // Postal code
	GeoLat          float64 `protobuf:"fixed64,11,opt,name=geo_lat,json=geoLat,proto3" json:"geo_lat,omitempty"`                             // GPS latitude
	GeoLng          float64 `protobuf:"fixed64,12,opt,name=geo_lng,json=geoLng,proto3" json:"geo_lng,omitempty"`                             // GPS longitude
	GeoPrecisionRad uint64  `protobuf:"varint,13,opt,name=geo_precision_rad,json=geoPrecisionRad,proto3" json:"geo_precision_rad,omitempty"` // Radius precision in centimeters
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geo_def_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_geo_def_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_geo_def_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Location) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Location) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Location) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *Location) GetGeoLat() float64 {
	if x != nil {
		return x.GeoLat
	}
	return 0
}

func (x *Location) GetGeoLng() float64 {
	if x != nil {
		return x.GeoLng
	}
	return 0
}

func (x *Location) GetGeoPrecisionRad() uint64 {
	if x != nil {
		return x.GeoPrecisionRad
	}
	return 0
}

var File_geo_def_proto protoreflect.FileDescriptor

var file_geo_def_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x65, 0x6f, 0x2e, 0x64, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x63, 0x6f, 0x76, 0x69, 0x64, 0x31, 0x39, 0x61, 0x70, 0x69, 0x22, 0xcf, 0x01, 0x0a, 0x08,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x67, 0x65, 0x6f, 0x5f, 0x6c, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x06, 0x67, 0x65, 0x6f, 0x4c, 0x61, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x65, 0x6f, 0x5f,
	0x6c, 0x6e, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x67, 0x65, 0x6f, 0x4c, 0x6e,
	0x67, 0x12, 0x2a, 0x0a, 0x11, 0x67, 0x65, 0x6f, 0x5f, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x67, 0x65,
	0x6f, 0x50, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x64, 0x42, 0x34, 0x0a,
	0x0f, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x6f, 0x76, 0x69, 0x64, 0x31, 0x39, 0x2e, 0x61, 0x70, 0x69,
	0x42, 0x11, 0x47, 0x65, 0x6f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0c, 0x2e, 0x3b, 0x63, 0x6f, 0x76, 0x69, 0x64, 0x31, 0x39,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_geo_def_proto_rawDescOnce sync.Once
	file_geo_def_proto_rawDescData = file_geo_def_proto_rawDesc
)

func file_geo_def_proto_rawDescGZIP() []byte {
	file_geo_def_proto_rawDescOnce.Do(func() {
		file_geo_def_proto_rawDescData = protoimpl.X.CompressGZIP(file_geo_def_proto_rawDescData)
	})
	return file_geo_def_proto_rawDescData
}

var file_geo_def_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_geo_def_proto_goTypes = []interface{}{
	(*Location)(nil), // 0: covid19api.Location
}
var file_geo_def_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_geo_def_proto_init() }
func file_geo_def_proto_init() {
	if File_geo_def_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_geo_def_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_geo_def_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_geo_def_proto_goTypes,
		DependencyIndexes: file_geo_def_proto_depIdxs,
		MessageInfos:      file_geo_def_proto_msgTypes,
	}.Build()
	File_geo_def_proto = out.File
	file_geo_def_proto_rawDesc = nil
	file_geo_def_proto_goTypes = nil
	file_geo_def_proto_depIdxs = nil
}
