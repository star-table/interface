// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: officeapi.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WOPIRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName    string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	AccessToken string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Permission  string `protobuf:"bytes,3,opt,name=permission,proto3" json:"permission,omitempty"`
}

func (x *WOPIRequest) Reset() {
	*x = WOPIRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_officeapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WOPIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WOPIRequest) ProtoMessage() {}

func (x *WOPIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_officeapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WOPIRequest.ProtoReflect.Descriptor instead.
func (*WOPIRequest) Descriptor() ([]byte, []int) {
	return file_officeapi_proto_rawDescGZIP(), []int{0}
}

func (x *WOPIRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *WOPIRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *WOPIRequest) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

type WOPIRequestFileContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName    string          `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	AccessToken string          `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Permission  string          `protobuf:"bytes,3,opt,name=permission,proto3" json:"permission,omitempty"`
	Data        *structpb.Value `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *WOPIRequestFileContent) Reset() {
	*x = WOPIRequestFileContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_officeapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WOPIRequestFileContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WOPIRequestFileContent) ProtoMessage() {}

func (x *WOPIRequestFileContent) ProtoReflect() protoreflect.Message {
	mi := &file_officeapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WOPIRequestFileContent.ProtoReflect.Descriptor instead.
func (*WOPIRequestFileContent) Descriptor() ([]byte, []int) {
	return file_officeapi_proto_rawDescGZIP(), []int{1}
}

func (x *WOPIRequestFileContent) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *WOPIRequestFileContent) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *WOPIRequestFileContent) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *WOPIRequestFileContent) GetData() *structpb.Value {
	if x != nil {
		return x.Data
	}
	return nil
}

type FormConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FormConfigResponse) Reset() {
	*x = FormConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_officeapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FormConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FormConfigResponse) ProtoMessage() {}

func (x *FormConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_officeapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FormConfigResponse.ProtoReflect.Descriptor instead.
func (*FormConfigResponse) Descriptor() ([]byte, []int) {
	return file_officeapi_proto_rawDescGZIP(), []int{2}
}

type WOPIEditableResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseFileName            string `protobuf:"bytes,1,opt,name=BaseFileName,proto3" json:"BaseFileName,omitempty"`
	OwnerId                 string `protobuf:"bytes,2,opt,name=OwnerId,proto3" json:"OwnerId,omitempty"`
	UserId                  string `protobuf:"bytes,3,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Size                    int32  `protobuf:"varint,4,opt,name=Size,proto3" json:"Size,omitempty"`
	UserFriendlyName        string `protobuf:"bytes,5,opt,name=UserFriendlyName,proto3" json:"UserFriendlyName,omitempty"`
	UserCanWrite            bool   `protobuf:"varint,6,opt,name=UserCanWrite,proto3" json:"UserCanWrite,omitempty"`
	UserCanNotWriteRelative bool   `protobuf:"varint,7,opt,name=UserCanNotWriteRelative,proto3" json:"UserCanNotWriteRelative,omitempty"`
	Status                  string `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	LastModifiedTime        string `protobuf:"bytes,9,opt,name=LastModifiedTime,proto3" json:"LastModifiedTime,omitempty"`
}

func (x *WOPIEditableResp) Reset() {
	*x = WOPIEditableResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_officeapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WOPIEditableResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WOPIEditableResp) ProtoMessage() {}

func (x *WOPIEditableResp) ProtoReflect() protoreflect.Message {
	mi := &file_officeapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WOPIEditableResp.ProtoReflect.Descriptor instead.
func (*WOPIEditableResp) Descriptor() ([]byte, []int) {
	return file_officeapi_proto_rawDescGZIP(), []int{3}
}

func (x *WOPIEditableResp) GetBaseFileName() string {
	if x != nil {
		return x.BaseFileName
	}
	return ""
}

func (x *WOPIEditableResp) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *WOPIEditableResp) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *WOPIEditableResp) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *WOPIEditableResp) GetUserFriendlyName() string {
	if x != nil {
		return x.UserFriendlyName
	}
	return ""
}

func (x *WOPIEditableResp) GetUserCanWrite() bool {
	if x != nil {
		return x.UserCanWrite
	}
	return false
}

func (x *WOPIEditableResp) GetUserCanNotWriteRelative() bool {
	if x != nil {
		return x.UserCanNotWriteRelative
	}
	return false
}

func (x *WOPIEditableResp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *WOPIEditableResp) GetLastModifiedTime() string {
	if x != nil {
		return x.LastModifiedTime
	}
	return ""
}

type GetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetConfigRequest) Reset() {
	*x = GetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_officeapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigRequest) ProtoMessage() {}

func (x *GetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_officeapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigRequest.ProtoReflect.Descriptor instead.
func (*GetConfigRequest) Descriptor() ([]byte, []int) {
	return file_officeapi_proto_rawDescGZIP(), []int{4}
}

type GetConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type           string `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`                     // 类型 type: collabora, microsoft
	UrlPrefix      string `protobuf:"bytes,2,opt,name=UrlPrefix,proto3" json:"UrlPrefix,omitempty"`           // OfficeUrl前缀
	FileExtensions string `protobuf:"bytes,3,opt,name=FileExtensions,proto3" json:"FileExtensions,omitempty"` // 文件后嘴
}

func (x *GetConfigResponse) Reset() {
	*x = GetConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_officeapi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigResponse) ProtoMessage() {}

func (x *GetConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_officeapi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigResponse.ProtoReflect.Descriptor instead.
func (*GetConfigResponse) Descriptor() ([]byte, []int) {
	return file_officeapi_proto_rawDescGZIP(), []int{5}
}

func (x *GetConfigResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetConfigResponse) GetUrlPrefix() string {
	if x != nil {
		return x.UrlPrefix
	}
	return ""
}

func (x *GetConfigResponse) GetFileExtensions() string {
	if x != nil {
		return x.FileExtensions
	}
	return ""
}

var File_officeapi_proto protoreflect.FileDescriptor

var file_officeapi_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x0b, 0x57,
	0x4f, 0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xa4, 0x01, 0x0a, 0x16, 0x57,
	0x4f, 0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x14, 0x0a, 0x12, 0x46, 0x6f, 0x72, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xca, 0x02, 0x0a, 0x10, 0x57, 0x4f, 0x50, 0x49,
	0x45, 0x64, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x0c,
	0x42, 0x61, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x43, 0x61,
	0x6e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72, 0x43, 0x61,
	0x6e, 0x4e, 0x6f, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x55, 0x73, 0x65, 0x72, 0x43, 0x61, 0x6e,
	0x4e, 0x6f, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x4c, 0x61, 0x73, 0x74,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x72, 0x6c, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x72, 0x6c, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12,
	0x26, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xe9, 0x01, 0x0a, 0x09, 0x4f, 0x66, 0x66, 0x69,
	0x63, 0x65, 0x41, 0x70, 0x69, 0x12, 0x61, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x1e, 0x2e, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x79, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x57,
	0x6f, 0x70, 0x69, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x64, 0x69, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x57, 0x4f, 0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x4f,
	0x50, 0x49, 0x45, 0x64, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x23,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x6f, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x7d, 0x42, 0x1f, 0x5a, 0x1d, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_officeapi_proto_rawDescOnce sync.Once
	file_officeapi_proto_rawDescData = file_officeapi_proto_rawDesc
)

func file_officeapi_proto_rawDescGZIP() []byte {
	file_officeapi_proto_rawDescOnce.Do(func() {
		file_officeapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_officeapi_proto_rawDescData)
	})
	return file_officeapi_proto_rawDescData
}

var file_officeapi_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_officeapi_proto_goTypes = []interface{}{
	(*WOPIRequest)(nil),            // 0: officeapi.v1.WOPIRequest
	(*WOPIRequestFileContent)(nil), // 1: officeapi.v1.WOPIRequestFileContent
	(*FormConfigResponse)(nil),     // 2: officeapi.v1.FormConfigResponse
	(*WOPIEditableResp)(nil),       // 3: officeapi.v1.WOPIEditableResp
	(*GetConfigRequest)(nil),       // 4: officeapi.v1.GetConfigRequest
	(*GetConfigResponse)(nil),      // 5: officeapi.v1.GetConfigResponse
	(*structpb.Value)(nil),         // 6: google.protobuf.Value
}
var file_officeapi_proto_depIdxs = []int32{
	6, // 0: officeapi.v1.WOPIRequestFileContent.data:type_name -> google.protobuf.Value
	4, // 1: officeapi.v1.OfficeApi.GetConfig:input_type -> officeapi.v1.GetConfigRequest
	0, // 2: officeapi.v1.OfficeApi.GetWopiFileInfoEditable:input_type -> officeapi.v1.WOPIRequest
	5, // 3: officeapi.v1.OfficeApi.GetConfig:output_type -> officeapi.v1.GetConfigResponse
	3, // 4: officeapi.v1.OfficeApi.GetWopiFileInfoEditable:output_type -> officeapi.v1.WOPIEditableResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_officeapi_proto_init() }
func file_officeapi_proto_init() {
	if File_officeapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_officeapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WOPIRequest); i {
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
		file_officeapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WOPIRequestFileContent); i {
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
		file_officeapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FormConfigResponse); i {
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
		file_officeapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WOPIEditableResp); i {
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
		file_officeapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigRequest); i {
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
		file_officeapi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigResponse); i {
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
			RawDescriptor: file_officeapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_officeapi_proto_goTypes,
		DependencyIndexes: file_officeapi_proto_depIdxs,
		MessageInfos:      file_officeapi_proto_msgTypes,
	}.Build()
	File_officeapi_proto = out.File
	file_officeapi_proto_rawDesc = nil
	file_officeapi_proto_goTypes = nil
	file_officeapi_proto_depIdxs = nil
}