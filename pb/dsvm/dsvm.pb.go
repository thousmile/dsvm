// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: dsvm.proto

package dsvm

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

type InitSchemasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InitSchemasRequest) Reset() {
	*x = InitSchemasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dsvm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitSchemasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitSchemasRequest) ProtoMessage() {}

func (x *InitSchemasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dsvm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitSchemasRequest.ProtoReflect.Descriptor instead.
func (*InitSchemasRequest) Descriptor() ([]byte, []int) {
	return file_dsvm_proto_rawDescGZIP(), []int{0}
}

type InitSchemasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *InitSchemasResponse) Reset() {
	*x = InitSchemasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dsvm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitSchemasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitSchemasResponse) ProtoMessage() {}

func (x *InitSchemasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dsvm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitSchemasResponse.ProtoReflect.Descriptor instead.
func (*InitSchemasResponse) Descriptor() ([]byte, []int) {
	return file_dsvm_proto_rawDescGZIP(), []int{1}
}

func (x *InitSchemasResponse) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

var File_dsvm_proto protoreflect.FileDescriptor

var file_dsvm_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x73, 0x76, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x73,
	0x76, 0x6d, 0x22, 0x14, 0x0a, 0x12, 0x49, 0x6e, 0x69, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2b, 0x0a, 0x13, 0x49, 0x6e, 0x69, 0x74,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x32, 0x4a, 0x0a, 0x04, 0x44, 0x73, 0x76, 0x6d, 0x12, 0x42, 0x0a,
	0x0b, 0x49, 0x6e, 0x69, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x12, 0x18, 0x2e, 0x64,
	0x73, 0x76, 0x6d, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x73, 0x76, 0x6d, 0x2e, 0x49, 0x6e,
	0x69, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x73, 0x76, 0x6d, 0x3b, 0x64,
	0x73, 0x76, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dsvm_proto_rawDescOnce sync.Once
	file_dsvm_proto_rawDescData = file_dsvm_proto_rawDesc
)

func file_dsvm_proto_rawDescGZIP() []byte {
	file_dsvm_proto_rawDescOnce.Do(func() {
		file_dsvm_proto_rawDescData = protoimpl.X.CompressGZIP(file_dsvm_proto_rawDescData)
	})
	return file_dsvm_proto_rawDescData
}

var file_dsvm_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dsvm_proto_goTypes = []interface{}{
	(*InitSchemasRequest)(nil),  // 0: dsvm.InitSchemasRequest
	(*InitSchemasResponse)(nil), // 1: dsvm.InitSchemasResponse
}
var file_dsvm_proto_depIdxs = []int32{
	0, // 0: dsvm.Dsvm.InitSchemas:input_type -> dsvm.InitSchemasRequest
	1, // 1: dsvm.Dsvm.InitSchemas:output_type -> dsvm.InitSchemasResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dsvm_proto_init() }
func file_dsvm_proto_init() {
	if File_dsvm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dsvm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitSchemasRequest); i {
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
		file_dsvm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitSchemasResponse); i {
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
			RawDescriptor: file_dsvm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dsvm_proto_goTypes,
		DependencyIndexes: file_dsvm_proto_depIdxs,
		MessageInfos:      file_dsvm_proto_msgTypes,
	}.Build()
	File_dsvm_proto = out.File
	file_dsvm_proto_rawDesc = nil
	file_dsvm_proto_goTypes = nil
	file_dsvm_proto_depIdxs = nil
}