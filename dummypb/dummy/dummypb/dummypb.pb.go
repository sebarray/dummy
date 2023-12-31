// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.5.1
// source: dummypb.proto

package dummypb

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

type DummyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *DummyRequest) Reset() {
	*x = DummyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dummypb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DummyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DummyRequest) ProtoMessage() {}

func (x *DummyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dummypb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DummyRequest.ProtoReflect.Descriptor instead.
func (*DummyRequest) Descriptor() ([]byte, []int) {
	return file_dummypb_proto_rawDescGZIP(), []int{0}
}

func (x *DummyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DummyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *DummyResponse) Reset() {
	*x = DummyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dummypb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DummyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DummyResponse) ProtoMessage() {}

func (x *DummyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dummypb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DummyResponse.ProtoReflect.Descriptor instead.
func (*DummyResponse) Descriptor() ([]byte, []int) {
	return file_dummypb_proto_rawDescGZIP(), []int{1}
}

func (x *DummyResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_dummypb_proto protoreflect.FileDescriptor

var file_dummypb_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x22, 0x22, 0x0a, 0x0c, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x0d, 0x44, 0x75,
	0x6d, 0x6d, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x32,
	0x42, 0x0a, 0x0c, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x32, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x13, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79,
	0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2f, 0x64, 0x75, 0x6d,
	0x6d, 0x79, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dummypb_proto_rawDescOnce sync.Once
	file_dummypb_proto_rawDescData = file_dummypb_proto_rawDesc
)

func file_dummypb_proto_rawDescGZIP() []byte {
	file_dummypb_proto_rawDescOnce.Do(func() {
		file_dummypb_proto_rawDescData = protoimpl.X.CompressGZIP(file_dummypb_proto_rawDescData)
	})
	return file_dummypb_proto_rawDescData
}

var file_dummypb_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dummypb_proto_goTypes = []interface{}{
	(*DummyRequest)(nil),  // 0: dummy.DummyRequest
	(*DummyResponse)(nil), // 1: dummy.DummyResponse
}
var file_dummypb_proto_depIdxs = []int32{
	0, // 0: dummy.DummyService.Hello:input_type -> dummy.DummyRequest
	1, // 1: dummy.DummyService.Hello:output_type -> dummy.DummyResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dummypb_proto_init() }
func file_dummypb_proto_init() {
	if File_dummypb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dummypb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DummyRequest); i {
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
		file_dummypb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DummyResponse); i {
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
			RawDescriptor: file_dummypb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dummypb_proto_goTypes,
		DependencyIndexes: file_dummypb_proto_depIdxs,
		MessageInfos:      file_dummypb_proto_msgTypes,
	}.Build()
	File_dummypb_proto = out.File
	file_dummypb_proto_rawDesc = nil
	file_dummypb_proto_goTypes = nil
	file_dummypb_proto_depIdxs = nil
}
