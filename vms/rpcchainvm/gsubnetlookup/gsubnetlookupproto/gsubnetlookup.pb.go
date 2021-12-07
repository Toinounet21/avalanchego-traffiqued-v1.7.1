// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: gsubnetlookup.proto

package gsubnetlookupproto

import (
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

type SubnetIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainID []byte `protobuf:"bytes,1,opt,name=chainID,proto3" json:"chainID,omitempty"`
}

func (x *SubnetIDRequest) Reset() {
	*x = SubnetIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gsubnetlookup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubnetIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubnetIDRequest) ProtoMessage() {}

func (x *SubnetIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gsubnetlookup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubnetIDRequest.ProtoReflect.Descriptor instead.
func (*SubnetIDRequest) Descriptor() ([]byte, []int) {
	return file_gsubnetlookup_proto_rawDescGZIP(), []int{0}
}

func (x *SubnetIDRequest) GetChainID() []byte {
	if x != nil {
		return x.ChainID
	}
	return nil
}

type SubnetIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SubnetIDResponse) Reset() {
	*x = SubnetIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gsubnetlookup_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubnetIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubnetIDResponse) ProtoMessage() {}

func (x *SubnetIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gsubnetlookup_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubnetIDResponse.ProtoReflect.Descriptor instead.
func (*SubnetIDResponse) Descriptor() ([]byte, []int) {
	return file_gsubnetlookup_proto_rawDescGZIP(), []int{1}
}

func (x *SubnetIDResponse) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

var File_gsubnetlookup_proto protoreflect.FileDescriptor

var file_gsubnetlookup_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x6c, 0x6f,
	0x6f, 0x6b, 0x75, 0x70, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0f, 0x53, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x22, 0x22, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x32, 0x65, 0x0a, 0x0c, 0x53, 0x75,
	0x62, 0x6e, 0x65, 0x74, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x12, 0x55, 0x0a, 0x08, 0x53, 0x75,
	0x62, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x12, 0x23, 0x2e, 0x67, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74,
	0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x6e,
	0x65, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67, 0x73,
	0x75, 0x62, 0x6e, 0x65, 0x74, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x76, 0x61, 0x2d, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x61, 0x76, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x68, 0x65, 0x67, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x76, 0x6d, 0x2f,
	0x67, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2f, 0x67, 0x73,
	0x75, 0x62, 0x6e, 0x65, 0x74, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gsubnetlookup_proto_rawDescOnce sync.Once
	file_gsubnetlookup_proto_rawDescData = file_gsubnetlookup_proto_rawDesc
)

func file_gsubnetlookup_proto_rawDescGZIP() []byte {
	file_gsubnetlookup_proto_rawDescOnce.Do(func() {
		file_gsubnetlookup_proto_rawDescData = protoimpl.X.CompressGZIP(file_gsubnetlookup_proto_rawDescData)
	})
	return file_gsubnetlookup_proto_rawDescData
}

var file_gsubnetlookup_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gsubnetlookup_proto_goTypes = []interface{}{
	(*SubnetIDRequest)(nil),  // 0: gsubnetlookupproto.SubnetIDRequest
	(*SubnetIDResponse)(nil), // 1: gsubnetlookupproto.SubnetIDResponse
}
var file_gsubnetlookup_proto_depIdxs = []int32{
	0, // 0: gsubnetlookupproto.SubnetLookup.SubnetID:input_type -> gsubnetlookupproto.SubnetIDRequest
	1, // 1: gsubnetlookupproto.SubnetLookup.SubnetID:output_type -> gsubnetlookupproto.SubnetIDResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gsubnetlookup_proto_init() }
func file_gsubnetlookup_proto_init() {
	if File_gsubnetlookup_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gsubnetlookup_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubnetIDRequest); i {
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
		file_gsubnetlookup_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubnetIDResponse); i {
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
			RawDescriptor: file_gsubnetlookup_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gsubnetlookup_proto_goTypes,
		DependencyIndexes: file_gsubnetlookup_proto_depIdxs,
		MessageInfos:      file_gsubnetlookup_proto_msgTypes,
	}.Build()
	File_gsubnetlookup_proto = out.File
	file_gsubnetlookup_proto_rawDesc = nil
	file_gsubnetlookup_proto_goTypes = nil
	file_gsubnetlookup_proto_depIdxs = nil
}