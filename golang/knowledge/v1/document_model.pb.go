// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.7.1
// source: document_model.proto

package v1

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_document_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_document_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_document_model_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type DocumentSimple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentId string `protobuf:"bytes,1,opt,name=documentId,proto3" json:"documentId,omitempty"`
	Title      string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	ParentId   string `protobuf:"bytes,3,opt,name=parentId,proto3" json:"parentId,omitempty"`
}

func (x *DocumentSimple) Reset() {
	*x = DocumentSimple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_document_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentSimple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentSimple) ProtoMessage() {}

func (x *DocumentSimple) ProtoReflect() protoreflect.Message {
	mi := &file_document_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentSimple.ProtoReflect.Descriptor instead.
func (*DocumentSimple) Descriptor() ([]byte, []int) {
	return file_document_model_proto_rawDescGZIP(), []int{1}
}

func (x *DocumentSimple) GetDocumentId() string {
	if x != nil {
		return x.DocumentId
	}
	return ""
}

func (x *DocumentSimple) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DocumentSimple) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

type DocumentDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentId string `protobuf:"bytes,1,opt,name=documentId,proto3" json:"documentId,omitempty"`
	Content    string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Title      string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	ParentId   string `protobuf:"bytes,4,opt,name=parentId,proto3" json:"parentId,omitempty"`
	CreateTime int64  `protobuf:"varint,5,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime int64  `protobuf:"varint,6,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	Creator    *User  `protobuf:"bytes,7,opt,name=creator,proto3" json:"creator,omitempty"`
	Updater    *User  `protobuf:"bytes,8,opt,name=updater,proto3" json:"updater,omitempty"`
}

func (x *DocumentDetail) Reset() {
	*x = DocumentDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_document_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentDetail) ProtoMessage() {}

func (x *DocumentDetail) ProtoReflect() protoreflect.Message {
	mi := &file_document_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentDetail.ProtoReflect.Descriptor instead.
func (*DocumentDetail) Descriptor() ([]byte, []int) {
	return file_document_model_proto_rawDescGZIP(), []int{2}
}

func (x *DocumentDetail) GetDocumentId() string {
	if x != nil {
		return x.DocumentId
	}
	return ""
}

func (x *DocumentDetail) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *DocumentDetail) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DocumentDetail) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *DocumentDetail) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *DocumentDetail) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *DocumentDetail) GetCreator() *User {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *DocumentDetail) GetUpdater() *User {
	if x != nil {
		return x.Updater
	}
	return nil
}

type SpaceDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceId string `protobuf:"bytes,1,opt,name=spaceId,proto3" json:"spaceId,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Cover   string `protobuf:"bytes,3,opt,name=cover,proto3" json:"cover,omitempty"`
	Type    int32  `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *SpaceDetail) Reset() {
	*x = SpaceDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_document_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpaceDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpaceDetail) ProtoMessage() {}

func (x *SpaceDetail) ProtoReflect() protoreflect.Message {
	mi := &file_document_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpaceDetail.ProtoReflect.Descriptor instead.
func (*SpaceDetail) Descriptor() ([]byte, []int) {
	return file_document_model_proto_rawDescGZIP(), []int{3}
}

func (x *SpaceDetail) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *SpaceDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SpaceDetail) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *SpaceDetail) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type CommentDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId   int64  `protobuf:"varint,1,opt,name=commentId,proto3" json:"commentId,omitempty"`
	Comment     string `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	MentionUser *User  `protobuf:"bytes,3,opt,name=mentionUser,proto3" json:"mentionUser,omitempty"`
	Creator     *User  `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
	CreateTime  int64  `protobuf:"varint,5,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	UpdateTime  int64  `protobuf:"varint,6,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
}

func (x *CommentDetail) Reset() {
	*x = CommentDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_document_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentDetail) ProtoMessage() {}

func (x *CommentDetail) ProtoReflect() protoreflect.Message {
	mi := &file_document_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentDetail.ProtoReflect.Descriptor instead.
func (*CommentDetail) Descriptor() ([]byte, []int) {
	return file_document_model_proto_rawDescGZIP(), []int{4}
}

func (x *CommentDetail) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *CommentDetail) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *CommentDetail) GetMentionUser() *User {
	if x != nil {
		return x.MentionUser
	}
	return nil
}

func (x *CommentDetail) GetCreator() *User {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *CommentDetail) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *CommentDetail) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

type RecycleDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecycleId  int64  `protobuf:"varint,1,opt,name=recycleId,proto3" json:"recycleId,omitempty"`
	Title      string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Creator    *User  `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	CreateTime int64  `protobuf:"varint,4,opt,name=createTime,proto3" json:"createTime,omitempty"`
	LeftDay    int64  `protobuf:"varint,5,opt,name=leftDay,proto3" json:"leftDay,omitempty"` // 剩余删除天数
}

func (x *RecycleDetail) Reset() {
	*x = RecycleDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_document_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecycleDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecycleDetail) ProtoMessage() {}

func (x *RecycleDetail) ProtoReflect() protoreflect.Message {
	mi := &file_document_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecycleDetail.ProtoReflect.Descriptor instead.
func (*RecycleDetail) Descriptor() ([]byte, []int) {
	return file_document_model_proto_rawDescGZIP(), []int{5}
}

func (x *RecycleDetail) GetRecycleId() int64 {
	if x != nil {
		return x.RecycleId
	}
	return 0
}

func (x *RecycleDetail) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *RecycleDetail) GetCreator() *User {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *RecycleDetail) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *RecycleDetail) GetLeftDay() int64 {
	if x != nil {
		return x.LeftDay
	}
	return 0
}

var File_document_model_proto protoreflect.FileDescriptor

var file_document_model_proto_rawDesc = []byte{
	0x0a, 0x14, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x22, 0x4a, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x22, 0x62, 0x0a, 0x0e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0x98, 0x02, 0x0a, 0x0e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x2c, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x22,
	0x65, 0x0a, 0x0b, 0x53, 0x70, 0x61, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xeb, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x34, 0x0a, 0x0b, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x0b, 0x6d, 0x65, 0x6e, 0x74, 0x69,
	0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0xab, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x65, 0x66, 0x74,
	0x44, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6c, 0x65, 0x66, 0x74, 0x44,
	0x61, 0x79, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x65, 0x61, 0x2e, 0x62, 0x6a, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x4c, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f,
	0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_document_model_proto_rawDescOnce sync.Once
	file_document_model_proto_rawDescData = file_document_model_proto_rawDesc
)

func file_document_model_proto_rawDescGZIP() []byte {
	file_document_model_proto_rawDescOnce.Do(func() {
		file_document_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_document_model_proto_rawDescData)
	})
	return file_document_model_proto_rawDescData
}

var file_document_model_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_document_model_proto_goTypes = []interface{}{
	(*User)(nil),           // 0: knowledge.v1.User
	(*DocumentSimple)(nil), // 1: knowledge.v1.DocumentSimple
	(*DocumentDetail)(nil), // 2: knowledge.v1.DocumentDetail
	(*SpaceDetail)(nil),    // 3: knowledge.v1.SpaceDetail
	(*CommentDetail)(nil),  // 4: knowledge.v1.CommentDetail
	(*RecycleDetail)(nil),  // 5: knowledge.v1.RecycleDetail
}
var file_document_model_proto_depIdxs = []int32{
	0, // 0: knowledge.v1.DocumentDetail.creator:type_name -> knowledge.v1.User
	0, // 1: knowledge.v1.DocumentDetail.updater:type_name -> knowledge.v1.User
	0, // 2: knowledge.v1.CommentDetail.mentionUser:type_name -> knowledge.v1.User
	0, // 3: knowledge.v1.CommentDetail.creator:type_name -> knowledge.v1.User
	0, // 4: knowledge.v1.RecycleDetail.creator:type_name -> knowledge.v1.User
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_document_model_proto_init() }
func file_document_model_proto_init() {
	if File_document_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_document_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_document_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentSimple); i {
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
		file_document_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentDetail); i {
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
		file_document_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpaceDetail); i {
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
		file_document_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentDetail); i {
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
		file_document_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecycleDetail); i {
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
			RawDescriptor: file_document_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_document_model_proto_goTypes,
		DependencyIndexes: file_document_model_proto_depIdxs,
		MessageInfos:      file_document_model_proto_msgTypes,
	}.Build()
	File_document_model_proto = out.File
	file_document_model_proto_rawDesc = nil
	file_document_model_proto_goTypes = nil
	file_document_model_proto_depIdxs = nil
}
