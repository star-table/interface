// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: batch.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The request message containing the user's name.
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OrgId int64  `protobuf:"varint,2,opt,name=orgId,proto3" json:"orgId,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_batch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_batch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_batch_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloRequest) GetOrgId() int64 {
	if x != nil {
		return x.OrgId
	}
	return 0
}

// The response message containing the greetings
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_batch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_batch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_batch_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_batch_proto protoreflect.FileDescriptor

var file_batch_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x67,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x22,
	0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xb1, 0x15, 0x0a, 0x05, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x57, 0x0a, 0x04, 0x54, 0x65, 0x6d, 0x70, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12,
	0x19, 0x2f, 0x7a, 0x6f, 0x34, 0x77, 0x72, 0x72, 0x4e, 0x59, 0x42, 0x43, 0x78, 0x39, 0x39, 0x37,
	0x41, 0x53, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x7d, 0x12, 0x5f, 0x0a, 0x0c, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x50, 0x47, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b,
	0x12, 0x19, 0x2f, 0x54, 0x33, 0x33, 0x52, 0x6f, 0x68, 0x37, 0x65, 0x71, 0x79, 0x54, 0x67, 0x6b,
	0x48, 0x39, 0x45, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x7d, 0x12, 0x63, 0x0a, 0x10, 0x53,
	0x70, 0x6c, 0x69, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x16, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x31, 0x72, 0x30, 0x66, 0x79, 0x41, 0x32, 0x74,
	0x61, 0x62, 0x55, 0x46, 0x6b, 0x65, 0x58, 0x62, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x7d,
	0x12, 0x62, 0x0a, 0x0f, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x57, 0x59, 0x48, 0x4e,
	0x36, 0x6e, 0x41, 0x45, 0x75, 0x6e, 0x76, 0x61, 0x53, 0x33, 0x69, 0x51, 0x2f, 0x7b, 0x6f, 0x72,
	0x67, 0x49, 0x64, 0x7d, 0x12, 0x5a, 0x0a, 0x0f, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x57,
	0x6f, 0x72, 0x6b, 0x48, 0x6f, 0x75, 0x72, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f,
	0x4f, 0x38, 0x49, 0x75, 0x49, 0x53, 0x54, 0x42, 0x65, 0x41, 0x66, 0x4d, 0x6c, 0x57, 0x76, 0x58,
	0x12, 0x66, 0x0a, 0x13, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x61,
	0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f,
	0x6d, 0x7a, 0x67, 0x30, 0x45, 0x41, 0x65, 0x39, 0x6f, 0x41, 0x4f, 0x58, 0x33, 0x41, 0x55, 0x64,
	0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x7d, 0x12, 0x6a, 0x0a, 0x17, 0x4d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x65, 0x44, 0x72, 0x6f, 0x70, 0x55, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x79, 0x6c, 0x52, 0x68,
	0x4e, 0x67, 0x4f, 0x79, 0x39, 0x58, 0x44, 0x51, 0x6b, 0x4e, 0x68, 0x45, 0x2f, 0x7b, 0x6f, 0x72,
	0x67, 0x49, 0x64, 0x7d, 0x12, 0x6d, 0x0a, 0x1a, 0x46, 0x69, 0x78, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x6f, 0x73, 0x74, 0x57, 0x68, 0x65, 0x6e, 0x4d, 0x6f,
	0x76, 0x65, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x66, 0x50, 0x36, 0x67, 0x46,
	0x5a, 0x56, 0x33, 0x68, 0x36, 0x61, 0x59, 0x71, 0x48, 0x77, 0x5a, 0x2f, 0x7b, 0x6f, 0x72, 0x67,
	0x49, 0x64, 0x7d, 0x12, 0x5d, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x50, 0x47, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x74, 0x67, 0x43, 0x45, 0x34, 0x64,
	0x67, 0x65, 0x32, 0x34, 0x68, 0x35, 0x78, 0x65, 0x5a, 0x4f, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x49,
	0x64, 0x7d, 0x12, 0x5c, 0x0a, 0x11, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x47, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f,
	0x35, 0x73, 0x53, 0x50, 0x4f, 0x4a, 0x79, 0x31, 0x69, 0x69, 0x70, 0x4c, 0x4a, 0x68, 0x65, 0x50,
	0x12, 0x5d, 0x0a, 0x0a, 0x46, 0x69, 0x78, 0x50, 0x47, 0x4f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x16,
	0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x39, 0x44, 0x44, 0x61, 0x72, 0x55, 0x74, 0x34, 0x57,
	0x34, 0x67, 0x48, 0x41, 0x6c, 0x63, 0x6d, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x7d, 0x12,
	0x5e, 0x0a, 0x0b, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x16,
	0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x4c, 0x47, 0x52, 0x68, 0x43, 0x59, 0x57, 0x62, 0x46,
	0x67, 0x49, 0x69, 0x47, 0x46, 0x6c, 0x68, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x7d, 0x12,
	0x6a, 0x0a, 0x17, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43,
	0x68, 0x61, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b,
	0x12, 0x19, 0x2f, 0x79, 0x6f, 0x31, 0x32, 0x71, 0x68, 0x4a, 0x79, 0x54, 0x39, 0x48, 0x33, 0x77,
	0x6c, 0x6a, 0x61, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x7d, 0x12, 0x65, 0x0a, 0x12, 0x46,
	0x69, 0x78, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x45, 0x34, 0x42, 0x36, 0x6e,
	0x74, 0x33, 0x55, 0x6c, 0x32, 0x39, 0x76, 0x35, 0x4c, 0x57, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x49,
	0x64, 0x7d, 0x12, 0x64, 0x0a, 0x11, 0x46, 0x69, 0x78, 0x49, 0x73, 0x73, 0x75, 0x65, 0x46, 0x6f,
	0x72, 0x44, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f,
	0x6d, 0x62, 0x6f, 0x35, 0x43, 0x55, 0x6a, 0x39, 0x65, 0x43, 0x35, 0x4a, 0x35, 0x67, 0x4b, 0x46,
	0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x7d, 0x12, 0x64, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x48,
	0x61, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x64, 0x75, 0x65, 0x56, 0x69, 0x65, 0x77, 0x12, 0x16, 0x2e,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x77, 0x6b, 0x78, 0x4e, 0x72, 0x49, 0x47, 0x42, 0x6b, 0x6f,
	0x47, 0x66, 0x72, 0x6d, 0x59, 0x54, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x7d, 0x12, 0x67,
	0x0a, 0x14, 0x41, 0x64, 0x64, 0x50, 0x67, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x42,
	0x69, 0x57, 0x53, 0x73, 0x53, 0x75, 0x45, 0x78, 0x58, 0x47, 0x66, 0x57, 0x46, 0x73, 0x58, 0x2f,
	0x7b, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x7d, 0x12, 0x5e, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x44, 0x61,
	0x73, 0x68, 0x42, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x6a,
	0x57, 0x59, 0x34, 0x39, 0x46, 0x57, 0x46, 0x68, 0x4f, 0x67, 0x36, 0x67, 0x59, 0x36, 0x30, 0x2f,
	0x7b, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x7d, 0x12, 0x5f, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f,
	0x75, 0x6c, 0x55, 0x6e, 0x74, 0x4e, 0x51, 0x62, 0x71, 0x4e, 0x54, 0x4e, 0x4e, 0x6a, 0x48, 0x55,
	0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x7d, 0x12, 0x61, 0x0a, 0x0e, 0x50, 0x67, 0x44, 0x61,
	0x74, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b,
	0x12, 0x19, 0x2f, 0x50, 0x67, 0x67, 0x5a, 0x6e, 0x35, 0x56, 0x32, 0x44, 0x38, 0x55, 0x6e, 0x67,
	0x41, 0x45, 0x77, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x7d, 0x12, 0x60, 0x0a, 0x0d, 0x43,
	0x6c, 0x65, 0x61, 0x72, 0x53, 0x61, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1b, 0x12, 0x19, 0x2f, 0x65, 0x78, 0x64, 0x49, 0x46, 0x56, 0x54, 0x59, 0x66, 0x33, 0x44,
	0x37, 0x43, 0x62, 0x61, 0x33, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x7d, 0x12, 0x65, 0x0a,
	0x12, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x41, 0x75, 0x64, 0x69, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x6c, 0x74, 0x69, 0x76,
	0x4d, 0x52, 0x53, 0x64, 0x74, 0x63, 0x67, 0x4e, 0x66, 0x4d, 0x65, 0x48, 0x2f, 0x7b, 0x6f, 0x72,
	0x67, 0x49, 0x64, 0x7d, 0x12, 0x5f, 0x0a, 0x0c, 0x44, 0x62, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x33, 0x31, 0x4c,
	0x4d, 0x4d, 0x38, 0x53, 0x49, 0x36, 0x31, 0x4b, 0x52, 0x6d, 0x4b, 0x61, 0x75, 0x2f, 0x7b, 0x6f,
	0x72, 0x67, 0x49, 0x64, 0x7d, 0x12, 0x6a, 0x0a, 0x17, 0x46, 0x69, 0x78, 0x32, 0x30, 0x32, 0x32,
	0x31, 0x31, 0x30, 0x34, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x16, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x32, 0x63, 0x6c, 0x35, 0x62, 0x5a, 0x35,
	0x67, 0x57, 0x61, 0x41, 0x35, 0x49, 0x45, 0x30, 0x4e, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x49, 0x64,
	0x7d, 0x12, 0x71, 0x0a, 0x1e, 0x46, 0x69, 0x78, 0x32, 0x30, 0x32, 0x32, 0x31, 0x31, 0x30, 0x37,
	0x41, 0x70, 0x70, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x36, 0x35, 0x53, 0x6a,
	0x68, 0x76, 0x46, 0x53, 0x6a, 0x74, 0x47, 0x37, 0x66, 0x63, 0x68, 0x44, 0x2f, 0x7b, 0x6f, 0x72,
	0x67, 0x49, 0x64, 0x7d, 0x12, 0x6a, 0x0a, 0x17, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12,
	0x16, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x33, 0x79, 0x30, 0x67, 0x56, 0x51, 0x56, 0x41,
	0x75, 0x38, 0x68, 0x42, 0x7a, 0x68, 0x31, 0x4a, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x7d,
	0x12, 0x69, 0x0a, 0x16, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b,
	0x12, 0x19, 0x2f, 0x45, 0x66, 0x49, 0x77, 0x5a, 0x35, 0x6f, 0x63, 0x30, 0x64, 0x42, 0x41, 0x47,
	0x42, 0x50, 0x37, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x7d, 0x42, 0x37, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x65, 0x61, 0x2e, 0x62, 0x6a, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x4c,
	0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_batch_proto_rawDescOnce sync.Once
	file_batch_proto_rawDescData = file_batch_proto_rawDesc
)

func file_batch_proto_rawDescGZIP() []byte {
	file_batch_proto_rawDescOnce.Do(func() {
		file_batch_proto_rawDescData = protoimpl.X.CompressGZIP(file_batch_proto_rawDescData)
	})
	return file_batch_proto_rawDescData
}

var file_batch_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_batch_proto_goTypes = []interface{}{
	(*HelloRequest)(nil), // 0: batch.v1.HelloRequest
	(*HelloReply)(nil),   // 1: batch.v1.HelloReply
}
var file_batch_proto_depIdxs = []int32{
	0,  // 0: batch.v1.Batch.Temp:input_type -> batch.v1.HelloRequest
	0,  // 1: batch.v1.Batch.ExportPGData:input_type -> batch.v1.HelloRequest
	0,  // 2: batch.v1.Batch.SplitTableHeader:input_type -> batch.v1.HelloRequest
	0,  // 3: batch.v1.Batch.MigrateRelating:input_type -> batch.v1.HelloRequest
	0,  // 4: batch.v1.Batch.MigrateWorkHour:input_type -> batch.v1.HelloRequest
	0,  // 5: batch.v1.Batch.MigrateCollaborator:input_type -> batch.v1.HelloRequest
	0,  // 6: batch.v1.Batch.MigrateDropUnusedTables:input_type -> batch.v1.HelloRequest
	0,  // 7: batch.v1.Batch.FixIssueStatusLostWhenMove:input_type -> batch.v1.HelloRequest
	0,  // 8: batch.v1.Batch.AddPGIndex:input_type -> batch.v1.HelloRequest
	0,  // 9: batch.v1.Batch.MigrateGlobalUser:input_type -> batch.v1.HelloRequest
	0,  // 10: batch.v1.Batch.FixPGOrgId:input_type -> batch.v1.HelloRequest
	0,  // 11: batch.v1.Batch.MigratePath:input_type -> batch.v1.HelloRequest
	0,  // 12: batch.v1.Batch.MigrateGroupChatSetting:input_type -> batch.v1.HelloRequest
	0,  // 13: batch.v1.Batch.FixIssueStatusName:input_type -> batch.v1.HelloRequest
	0,  // 14: batch.v1.Batch.FixIssueForDelPro:input_type -> batch.v1.HelloRequest
	0,  // 15: batch.v1.Batch.AddHasOverdueView:input_type -> batch.v1.HelloRequest
	0,  // 16: batch.v1.Batch.AddPgStatisticsTable:input_type -> batch.v1.HelloRequest
	0,  // 17: batch.v1.Batch.AddDashBord:input_type -> batch.v1.HelloRequest
	0,  // 18: batch.v1.Batch.ChangeDomain:input_type -> batch.v1.HelloRequest
	0,  // 19: batch.v1.Batch.PgDataTransfer:input_type -> batch.v1.HelloRequest
	0,  // 20: batch.v1.Batch.ClearSameUser:input_type -> batch.v1.HelloRequest
	0,  // 21: batch.v1.Batch.MigrateAuditStatus:input_type -> batch.v1.HelloRequest
	0,  // 22: batch.v1.Batch.DbMigrations:input_type -> batch.v1.HelloRequest
	0,  // 23: batch.v1.Batch.Fix20221104TemplatePath:input_type -> batch.v1.HelloRequest
	0,  // 24: batch.v1.Batch.Fix20221107AppCreatorAddMember:input_type -> batch.v1.HelloRequest
	0,  // 25: batch.v1.Batch.AutomationRebuildActive:input_type -> batch.v1.HelloRequest
	0,  // 26: batch.v1.Batch.ExchangeRelatingDataId:input_type -> batch.v1.HelloRequest
	1,  // 27: batch.v1.Batch.Temp:output_type -> batch.v1.HelloReply
	1,  // 28: batch.v1.Batch.ExportPGData:output_type -> batch.v1.HelloReply
	1,  // 29: batch.v1.Batch.SplitTableHeader:output_type -> batch.v1.HelloReply
	1,  // 30: batch.v1.Batch.MigrateRelating:output_type -> batch.v1.HelloReply
	1,  // 31: batch.v1.Batch.MigrateWorkHour:output_type -> batch.v1.HelloReply
	1,  // 32: batch.v1.Batch.MigrateCollaborator:output_type -> batch.v1.HelloReply
	1,  // 33: batch.v1.Batch.MigrateDropUnusedTables:output_type -> batch.v1.HelloReply
	1,  // 34: batch.v1.Batch.FixIssueStatusLostWhenMove:output_type -> batch.v1.HelloReply
	1,  // 35: batch.v1.Batch.AddPGIndex:output_type -> batch.v1.HelloReply
	1,  // 36: batch.v1.Batch.MigrateGlobalUser:output_type -> batch.v1.HelloReply
	1,  // 37: batch.v1.Batch.FixPGOrgId:output_type -> batch.v1.HelloReply
	1,  // 38: batch.v1.Batch.MigratePath:output_type -> batch.v1.HelloReply
	1,  // 39: batch.v1.Batch.MigrateGroupChatSetting:output_type -> batch.v1.HelloReply
	1,  // 40: batch.v1.Batch.FixIssueStatusName:output_type -> batch.v1.HelloReply
	1,  // 41: batch.v1.Batch.FixIssueForDelPro:output_type -> batch.v1.HelloReply
	1,  // 42: batch.v1.Batch.AddHasOverdueView:output_type -> batch.v1.HelloReply
	1,  // 43: batch.v1.Batch.AddPgStatisticsTable:output_type -> batch.v1.HelloReply
	1,  // 44: batch.v1.Batch.AddDashBord:output_type -> batch.v1.HelloReply
	1,  // 45: batch.v1.Batch.ChangeDomain:output_type -> batch.v1.HelloReply
	1,  // 46: batch.v1.Batch.PgDataTransfer:output_type -> batch.v1.HelloReply
	1,  // 47: batch.v1.Batch.ClearSameUser:output_type -> batch.v1.HelloReply
	1,  // 48: batch.v1.Batch.MigrateAuditStatus:output_type -> batch.v1.HelloReply
	1,  // 49: batch.v1.Batch.DbMigrations:output_type -> batch.v1.HelloReply
	1,  // 50: batch.v1.Batch.Fix20221104TemplatePath:output_type -> batch.v1.HelloReply
	1,  // 51: batch.v1.Batch.Fix20221107AppCreatorAddMember:output_type -> batch.v1.HelloReply
	1,  // 52: batch.v1.Batch.AutomationRebuildActive:output_type -> batch.v1.HelloReply
	1,  // 53: batch.v1.Batch.ExchangeRelatingDataId:output_type -> batch.v1.HelloReply
	27, // [27:54] is the sub-list for method output_type
	0,  // [0:27] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_batch_proto_init() }
func file_batch_proto_init() {
	if File_batch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_batch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_batch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
			RawDescriptor: file_batch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_batch_proto_goTypes,
		DependencyIndexes: file_batch_proto_depIdxs,
		MessageInfos:      file_batch_proto_msgTypes,
	}.Build()
	File_batch_proto = out.File
	file_batch_proto_rawDesc = nil
	file_batch_proto_goTypes = nil
	file_batch_proto_depIdxs = nil
}
