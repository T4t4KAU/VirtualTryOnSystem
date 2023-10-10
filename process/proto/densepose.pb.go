// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: densepose.proto

package proto

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

type DensePoseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *DensePoseRequest) Reset() {
	*x = DensePoseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_densepose_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DensePoseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DensePoseRequest) ProtoMessage() {}

func (x *DensePoseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_densepose_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DensePoseRequest.ProtoReflect.Descriptor instead.
func (*DensePoseRequest) Descriptor() ([]byte, []int) {
	return file_densepose_proto_rawDescGZIP(), []int{0}
}

func (x *DensePoseRequest) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

type DensePoseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DensePoseReply) Reset() {
	*x = DensePoseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_densepose_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DensePoseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DensePoseReply) ProtoMessage() {}

func (x *DensePoseReply) ProtoReflect() protoreflect.Message {
	mi := &file_densepose_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DensePoseReply.ProtoReflect.Descriptor instead.
func (*DensePoseReply) Descriptor() ([]byte, []int) {
	return file_densepose_proto_rawDescGZIP(), []int{1}
}

func (x *DensePoseReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_densepose_proto protoreflect.FileDescriptor

var file_densepose_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x65, 0x6e, 0x73, 0x65, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x28, 0x0a, 0x10, 0x44, 0x65, 0x6e, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x0e, 0x44,
	0x65, 0x6e, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x3d, 0x0a, 0x09, 0x44, 0x65, 0x6e, 0x73, 0x65,
	0x70, 0x6f, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x12, 0x11, 0x2e, 0x44, 0x65, 0x6e, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x44, 0x65, 0x6e, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_densepose_proto_rawDescOnce sync.Once
	file_densepose_proto_rawDescData = file_densepose_proto_rawDesc
)

func file_densepose_proto_rawDescGZIP() []byte {
	file_densepose_proto_rawDescOnce.Do(func() {
		file_densepose_proto_rawDescData = protoimpl.X.CompressGZIP(file_densepose_proto_rawDescData)
	})
	return file_densepose_proto_rawDescData
}

var file_densepose_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_densepose_proto_goTypes = []interface{}{
	(*DensePoseRequest)(nil), // 0: DensePoseRequest
	(*DensePoseReply)(nil),   // 1: DensePoseReply
}
var file_densepose_proto_depIdxs = []int32{
	0, // 0: Densepose.Generate:input_type -> DensePoseRequest
	1, // 1: Densepose.Generate:output_type -> DensePoseReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_densepose_proto_init() }
func file_densepose_proto_init() {
	if File_densepose_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_densepose_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DensePoseRequest); i {
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
		file_densepose_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DensePoseReply); i {
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
			RawDescriptor: file_densepose_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_densepose_proto_goTypes,
		DependencyIndexes: file_densepose_proto_depIdxs,
		MessageInfos:      file_densepose_proto_msgTypes,
	}.Build()
	File_densepose_proto = out.File
	file_densepose_proto_rawDesc = nil
	file_densepose_proto_goTypes = nil
	file_densepose_proto_depIdxs = nil
}
