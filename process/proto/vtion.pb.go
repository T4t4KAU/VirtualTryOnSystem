// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: vtion.proto

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

type DataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cloth      []byte `protobuf:"bytes,1,opt,name=cloth,proto3" json:"cloth,omitempty"`
	ClothMask  []byte `protobuf:"bytes,2,opt,name=cloth_mask,json=clothMask,proto3" json:"cloth_mask,omitempty"`
	Image      []byte `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	ImageParse []byte `protobuf:"bytes,4,opt,name=image_parse,json=imageParse,proto3" json:"image_parse,omitempty"`
	ImagePose  []byte `protobuf:"bytes,5,opt,name=image_pose,json=imagePose,proto3" json:"image_pose,omitempty"`
	PoseJson   string `protobuf:"bytes,6,opt,name=pose_json,json=poseJson,proto3" json:"pose_json,omitempty"`
	Name       string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DataRequest) Reset() {
	*x = DataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vtion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataRequest) ProtoMessage() {}

func (x *DataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vtion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataRequest.ProtoReflect.Descriptor instead.
func (*DataRequest) Descriptor() ([]byte, []int) {
	return file_vtion_proto_rawDescGZIP(), []int{0}
}

func (x *DataRequest) GetCloth() []byte {
	if x != nil {
		return x.Cloth
	}
	return nil
}

func (x *DataRequest) GetClothMask() []byte {
	if x != nil {
		return x.ClothMask
	}
	return nil
}

func (x *DataRequest) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *DataRequest) GetImageParse() []byte {
	if x != nil {
		return x.ImageParse
	}
	return nil
}

func (x *DataRequest) GetImagePose() []byte {
	if x != nil {
		return x.ImagePose
	}
	return nil
}

func (x *DataRequest) GetPoseJson() string {
	if x != nil {
		return x.PoseJson
	}
	return ""
}

func (x *DataRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DataReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []byte `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DataReply) Reset() {
	*x = DataReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vtion_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataReply) ProtoMessage() {}

func (x *DataReply) ProtoReflect() protoreflect.Message {
	mi := &file_vtion_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataReply.ProtoReflect.Descriptor instead.
func (*DataReply) Descriptor() ([]byte, []int) {
	return file_vtion_proto_rawDescGZIP(), []int{1}
}

func (x *DataReply) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_vtion_proto protoreflect.FileDescriptor

var file_vtion_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x01,
	0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6c, 0x6f, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x6c,
	0x6f, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x6f, 0x74, 0x68, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x6c, 0x6f, 0x74, 0x68, 0x4d, 0x61,
	0x73, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x50, 0x61, 0x72, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x50, 0x6f, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x65,
	0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73,
	0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x09, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x2f,
	0x0a, 0x05, 0x56, 0x69, 0x74, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x08, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0a, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_vtion_proto_rawDescOnce sync.Once
	file_vtion_proto_rawDescData = file_vtion_proto_rawDesc
)

func file_vtion_proto_rawDescGZIP() []byte {
	file_vtion_proto_rawDescOnce.Do(func() {
		file_vtion_proto_rawDescData = protoimpl.X.CompressGZIP(file_vtion_proto_rawDescData)
	})
	return file_vtion_proto_rawDescData
}

var file_vtion_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_vtion_proto_goTypes = []interface{}{
	(*DataRequest)(nil), // 0: DataRequest
	(*DataReply)(nil),   // 1: DataReply
}
var file_vtion_proto_depIdxs = []int32{
	0, // 0: Viton.Generate:input_type -> DataRequest
	1, // 1: Viton.Generate:output_type -> DataReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vtion_proto_init() }
func file_vtion_proto_init() {
	if File_vtion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vtion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataRequest); i {
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
		file_vtion_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataReply); i {
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
			RawDescriptor: file_vtion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vtion_proto_goTypes,
		DependencyIndexes: file_vtion_proto_depIdxs,
		MessageInfos:      file_vtion_proto_msgTypes,
	}.Build()
	File_vtion_proto = out.File
	file_vtion_proto_rawDesc = nil
	file_vtion_proto_goTypes = nil
	file_vtion_proto_depIdxs = nil
}