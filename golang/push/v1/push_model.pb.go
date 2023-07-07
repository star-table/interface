// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: push_model.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PushType int32

const (
	PushType_Card  PushType = 0
	PushType_Email PushType = 1
	PushType_Sms   PushType = 2
)

// Enum value maps for PushType.
var (
	PushType_name = map[int32]string{
		0: "Card",
		1: "Email",
		2: "Sms",
	}
	PushType_value = map[string]int32{
		"Card":  0,
		"Email": 1,
		"Sms":   2,
	}
)

func (x PushType) Enum() *PushType {
	p := new(PushType)
	*p = x
	return p
}

func (x PushType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PushType) Descriptor() protoreflect.EnumDescriptor {
	return file_push_model_proto_enumTypes[0].Descriptor()
}

func (PushType) Type() protoreflect.EnumType {
	return &file_push_model_proto_enumTypes[0]
}

func (x PushType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PushType.Descriptor instead.
func (PushType) EnumDescriptor() ([]byte, []int) {
	return file_push_model_proto_rawDescGZIP(), []int{0}
}

type PushMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    PushType   `protobuf:"varint,1,opt,name=type,proto3,enum=push.v1.PushType" json:"type,omitempty"`
	Payload *anypb.Any `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *PushMessage) Reset() {
	*x = PushMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMessage) ProtoMessage() {}

func (x *PushMessage) ProtoReflect() protoreflect.Message {
	mi := &file_push_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMessage.ProtoReflect.Descriptor instead.
func (*PushMessage) Descriptor() ([]byte, []int) {
	return file_push_model_proto_rawDescGZIP(), []int{0}
}

func (x *PushMessage) GetType() PushType {
	if x != nil {
		return x.Type
	}
	return PushType_Card
}

func (x *PushMessage) GetPayload() *anypb.Any {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_push_model_proto protoreflect.FileDescriptor

var file_push_model_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75,
	0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2a, 0x28, 0x0a, 0x08,
	0x50, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x10, 0x01, 0x12, 0x07, 0x0a,
	0x03, 0x53, 0x6d, 0x73, 0x10, 0x02, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x65, 0x61, 0x2e,
	0x62, 0x6a, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x4c, 0x65, 0x73, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_push_model_proto_rawDescOnce sync.Once
	file_push_model_proto_rawDescData = file_push_model_proto_rawDesc
)

func file_push_model_proto_rawDescGZIP() []byte {
	file_push_model_proto_rawDescOnce.Do(func() {
		file_push_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_push_model_proto_rawDescData)
	})
	return file_push_model_proto_rawDescData
}

var file_push_model_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_push_model_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_push_model_proto_goTypes = []interface{}{
	(PushType)(0),       // 0: push.v1.PushType
	(*PushMessage)(nil), // 1: push.v1.PushMessage
	(*anypb.Any)(nil),   // 2: google.protobuf.Any
}
var file_push_model_proto_depIdxs = []int32{
	0, // 0: push.v1.PushMessage.type:type_name -> push.v1.PushType
	2, // 1: push.v1.PushMessage.payload:type_name -> google.protobuf.Any
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_push_model_proto_init() }
func file_push_model_proto_init() {
	if File_push_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_push_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMessage); i {
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
			RawDescriptor: file_push_model_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_push_model_proto_goTypes,
		DependencyIndexes: file_push_model_proto_depIdxs,
		EnumInfos:         file_push_model_proto_enumTypes,
		MessageInfos:      file_push_model_proto_msgTypes,
	}.Build()
	File_push_model_proto = out.File
	file_push_model_proto_rawDesc = nil
	file_push_model_proto_goTypes = nil
	file_push_model_proto_depIdxs = nil
}