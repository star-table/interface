// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: msg_model.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category  EventCategory `protobuf:"varint,1,opt,name=category,proto3,enum=msg.v1.EventCategory" json:"category,omitempty"`
	Type      EventType     `protobuf:"varint,2,opt,name=type,proto3,enum=msg.v1.EventType" json:"type,omitempty"`
	Timestamp int64         `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TraceId   string        `protobuf:"bytes,4,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// Types that are assignable to PayloadOneof:
	//	*Event_DataEvent
	PayloadOneof isEvent_PayloadOneof `protobuf_oneof:"payload_oneof"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_msg_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_msg_model_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetCategory() EventCategory {
	if x != nil {
		return x.Category
	}
	return EventCategory_Invalid
}

func (x *Event) GetType() EventType {
	if x != nil {
		return x.Type
	}
	return EventType_DataCreated
}

func (x *Event) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Event) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (m *Event) GetPayloadOneof() isEvent_PayloadOneof {
	if m != nil {
		return m.PayloadOneof
	}
	return nil
}

func (x *Event) GetDataEvent() *DataEvent {
	if x, ok := x.GetPayloadOneof().(*Event_DataEvent); ok {
		return x.DataEvent
	}
	return nil
}

type isEvent_PayloadOneof interface {
	isEvent_PayloadOneof()
}

type Event_DataEvent struct {
	DataEvent *DataEvent `protobuf:"bytes,10,opt,name=data_event,json=dataEvent,proto3,oneof"`
}

func (*Event_DataEvent) isEvent_PayloadOneof() {}

type DataEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrgId       int64            `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	AppId       int64            `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	ProjectId   int64            `protobuf:"varint,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	TableId     int64            `protobuf:"varint,4,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
	DataId      int64            `protobuf:"varint,5,opt,name=data_id,json=dataId,proto3" json:"data_id,omitempty"`
	IssueId     int64            `protobuf:"varint,6,opt,name=issue_id,json=issueId,proto3" json:"issue_id,omitempty"`
	Title       string           `protobuf:"bytes,7,opt,name=title,proto3" json:"title,omitempty"`                       // 数据标题
	UserId      int64            `protobuf:"varint,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`      // 操作人ID
	UserName    string           `protobuf:"bytes,9,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"` // 操作人名字
	New         *structpb.Struct `protobuf:"bytes,20,opt,name=new,proto3" json:"new,omitempty"`                          // 字段值-完整新值
	Old         *structpb.Struct `protobuf:"bytes,21,opt,name=old,proto3" json:"old,omitempty"`                          // 字段值-完整旧值
	Incremental *structpb.Struct `protobuf:"bytes,22,opt,name=incremental,proto3" json:"incremental,omitempty"`          // 字段值-增量添加值
	Decremental *structpb.Struct `protobuf:"bytes,23,opt,name=decremental,proto3" json:"decremental,omitempty"`          // 字段值-增量删除值
}

func (x *DataEvent) Reset() {
	*x = DataEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataEvent) ProtoMessage() {}

func (x *DataEvent) ProtoReflect() protoreflect.Message {
	mi := &file_msg_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataEvent.ProtoReflect.Descriptor instead.
func (*DataEvent) Descriptor() ([]byte, []int) {
	return file_msg_model_proto_rawDescGZIP(), []int{1}
}

func (x *DataEvent) GetOrgId() int64 {
	if x != nil {
		return x.OrgId
	}
	return 0
}

func (x *DataEvent) GetAppId() int64 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *DataEvent) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *DataEvent) GetTableId() int64 {
	if x != nil {
		return x.TableId
	}
	return 0
}

func (x *DataEvent) GetDataId() int64 {
	if x != nil {
		return x.DataId
	}
	return 0
}

func (x *DataEvent) GetIssueId() int64 {
	if x != nil {
		return x.IssueId
	}
	return 0
}

func (x *DataEvent) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DataEvent) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DataEvent) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *DataEvent) GetNew() *structpb.Struct {
	if x != nil {
		return x.New
	}
	return nil
}

func (x *DataEvent) GetOld() *structpb.Struct {
	if x != nil {
		return x.Old
	}
	return nil
}

func (x *DataEvent) GetIncremental() *structpb.Struct {
	if x != nil {
		return x.Incremental
	}
	return nil
}

func (x *DataEvent) GetDecremental() *structpb.Struct {
	if x != nil {
		return x.Decremental
	}
	return nil
}

var File_msg_model_proto protoreflect.FileDescriptor

var file_msg_model_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x73, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x6d, 0x73, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0e, 0x6d, 0x73, 0x67, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe8, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6d,
	0x73, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x25, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x6d, 0x73,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x73, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52,
	0x09, 0x64, 0x61, 0x74, 0x61, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x0f, 0x0a, 0x0d, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x22, 0xfe, 0x03, 0x0a, 0x09,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x06, 0x6f, 0x72, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02,
	0x20, 0x00, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x06, 0x61, 0x70, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02,
	0x28, 0x00, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x22, 0x02, 0x28, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x22, 0x0a, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x28, 0x00, 0x52, 0x07, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52,
	0x06, 0x64, 0x61, 0x74, 0x61, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02,
	0x20, 0x00, 0x52, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x29, 0x0a, 0x03, 0x6e, 0x65, 0x77, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x03, 0x6e, 0x65, 0x77, 0x12, 0x29, 0x0a, 0x03, 0x6f,
	0x6c, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x03, 0x6f, 0x6c, 0x64, 0x12, 0x39, 0x0a, 0x0b, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x0b, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x12, 0x39, 0x0a, 0x0b, 0x64, 0x65, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x0b, 0x64, 0x65, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x42, 0x35, 0x5a, 0x33,
	0x67, 0x69, 0x74, 0x65, 0x61, 0x2e, 0x62, 0x6a, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x4c, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6d, 0x73, 0x67, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_model_proto_rawDescOnce sync.Once
	file_msg_model_proto_rawDescData = file_msg_model_proto_rawDesc
)

func file_msg_model_proto_rawDescGZIP() []byte {
	file_msg_model_proto_rawDescOnce.Do(func() {
		file_msg_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_model_proto_rawDescData)
	})
	return file_msg_model_proto_rawDescData
}

var file_msg_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_msg_model_proto_goTypes = []interface{}{
	(*Event)(nil),           // 0: msg.v1.Event
	(*DataEvent)(nil),       // 1: msg.v1.DataEvent
	(EventCategory)(0),      // 2: msg.v1.EventCategory
	(EventType)(0),          // 3: msg.v1.EventType
	(*structpb.Struct)(nil), // 4: google.protobuf.Struct
}
var file_msg_model_proto_depIdxs = []int32{
	2, // 0: msg.v1.Event.category:type_name -> msg.v1.EventCategory
	3, // 1: msg.v1.Event.type:type_name -> msg.v1.EventType
	1, // 2: msg.v1.Event.data_event:type_name -> msg.v1.DataEvent
	4, // 3: msg.v1.DataEvent.new:type_name -> google.protobuf.Struct
	4, // 4: msg.v1.DataEvent.old:type_name -> google.protobuf.Struct
	4, // 5: msg.v1.DataEvent.incremental:type_name -> google.protobuf.Struct
	4, // 6: msg.v1.DataEvent.decremental:type_name -> google.protobuf.Struct
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_msg_model_proto_init() }
func file_msg_model_proto_init() {
	if File_msg_model_proto != nil {
		return
	}
	file_msg_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_msg_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_msg_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataEvent); i {
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
	file_msg_model_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Event_DataEvent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msg_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_model_proto_goTypes,
		DependencyIndexes: file_msg_model_proto_depIdxs,
		MessageInfos:      file_msg_model_proto_msgTypes,
	}.Build()
	File_msg_model_proto = out.File
	file_msg_model_proto_rawDesc = nil
	file_msg_model_proto_goTypes = nil
	file_msg_model_proto_depIdxs = nil
}
