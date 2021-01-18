// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: notification/push/ios.proto

package ios

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Request_Priority int32

const (
	Request_LOW  Request_Priority = 0
	Request_HIGH Request_Priority = 1
)

// Enum value maps for Request_Priority.
var (
	Request_Priority_name = map[int32]string{
		0: "LOW",
		1: "HIGH",
	}
	Request_Priority_value = map[string]int32{
		"LOW":  0,
		"HIGH": 1,
	}
)

func (x Request_Priority) Enum() *Request_Priority {
	p := new(Request_Priority)
	*p = x
	return p
}

func (x Request_Priority) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Request_Priority) Descriptor() protoreflect.EnumDescriptor {
	return file_notification_push_ios_proto_enumTypes[0].Descriptor()
}

func (Request_Priority) Type() protoreflect.EnumType {
	return &file_notification_push_ios_proto_enumTypes[0]
}

func (x Request_Priority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Request_Priority.Descriptor instead.
func (Request_Priority) EnumDescriptor() ([]byte, []int) {
	return file_notification_push_ios_proto_rawDescGZIP(), []int{0, 0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guid   string `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	// device tokens/recipient
	Tokens         []string             `protobuf:"bytes,4,rep,name=tokens,proto3" json:"tokens,omitempty"`
	Payload        *Request_Message     `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	EventTimestamp *timestamp.Timestamp `protobuf:"bytes,6,opt,name=event_timestamp,json=eventTimestamp,proto3" json:"event_timestamp,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_push_ios_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_notification_push_ios_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_notification_push_ios_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetGuid() string {
	if x != nil {
		return x.Guid
	}
	return ""
}

func (x *Request) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Request) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Request) GetTokens() []string {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *Request) GetPayload() *Request_Message {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Request) GetEventTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.EventTimestamp
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_push_ios_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_notification_push_ios_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_notification_push_ios_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Request_Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alert  string            `protobuf:"bytes,1,opt,name=alert,proto3" json:"alert,omitempty"`
	Sound  string            `protobuf:"bytes,2,opt,name=sound,proto3" json:"sound,omitempty"`
	Custom map[string]string `protobuf:"bytes,3,rep,name=custom,proto3" json:"custom,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Request_Payload) Reset() {
	*x = Request_Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_push_ios_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Payload) ProtoMessage() {}

func (x *Request_Payload) ProtoReflect() protoreflect.Message {
	mi := &file_notification_push_ios_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Payload.ProtoReflect.Descriptor instead.
func (*Request_Payload) Descriptor() ([]byte, []int) {
	return file_notification_push_ios_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Request_Payload) GetAlert() string {
	if x != nil {
		return x.Alert
	}
	return ""
}

func (x *Request_Payload) GetSound() string {
	if x != nil {
		return x.Sound
	}
	return ""
}

func (x *Request_Payload) GetCustom() map[string]string {
	if x != nil {
		return x.Custom
	}
	return nil
}

type Request_Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic      string           `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	CollapseId string           `protobuf:"bytes,2,opt,name=collapse_id,json=collapseId,proto3" json:"collapse_id,omitempty"`
	Payload    *Request_Payload `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Priority   Request_Priority `protobuf:"varint,4,opt,name=priority,proto3,enum=protos.notification.push.ios.Request_Priority" json:"priority,omitempty"`
}

func (x *Request_Message) Reset() {
	*x = Request_Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_push_ios_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Message) ProtoMessage() {}

func (x *Request_Message) ProtoReflect() protoreflect.Message {
	mi := &file_notification_push_ios_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Message.ProtoReflect.Descriptor instead.
func (*Request_Message) Descriptor() ([]byte, []int) {
	return file_notification_push_ios_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Request_Message) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Request_Message) GetCollapseId() string {
	if x != nil {
		return x.CollapseId
	}
	return ""
}

func (x *Request_Message) GetPayload() *Request_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Request_Message) GetPriority() Request_Priority {
	if x != nil {
		return x.Priority
	}
	return Request_LOW
}

var File_notification_push_ios_proto protoreflect.FileDescriptor

var file_notification_push_ios_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70,
	0x75, 0x73, 0x68, 0x2f, 0x69, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x69, 0x6f, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x05, 0x0a,
	0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x12, 0x47, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x69,
	0x6f, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x43, 0x0a, 0x0f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x1a, 0xc3, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x51, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x75, 0x73, 0x68, 0x2e, 0x69, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x1a, 0x39, 0x0a, 0x0b, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xd5, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c,
	0x61, 0x70, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x6f, 0x6c, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x69, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x4a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e,
	0x69, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x1d,
	0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x4f,
	0x57, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x49, 0x47, 0x48, 0x10, 0x01, 0x22, 0x20, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42,
	0x17, 0x5a, 0x15, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x70, 0x75, 0x73, 0x68, 0x2f, 0x69, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notification_push_ios_proto_rawDescOnce sync.Once
	file_notification_push_ios_proto_rawDescData = file_notification_push_ios_proto_rawDesc
)

func file_notification_push_ios_proto_rawDescGZIP() []byte {
	file_notification_push_ios_proto_rawDescOnce.Do(func() {
		file_notification_push_ios_proto_rawDescData = protoimpl.X.CompressGZIP(file_notification_push_ios_proto_rawDescData)
	})
	return file_notification_push_ios_proto_rawDescData
}

var file_notification_push_ios_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_notification_push_ios_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_notification_push_ios_proto_goTypes = []interface{}{
	(Request_Priority)(0),       // 0: protos.notification.push.ios.Request.Priority
	(*Request)(nil),             // 1: protos.notification.push.ios.Request
	(*Response)(nil),            // 2: protos.notification.push.ios.Response
	(*Request_Payload)(nil),     // 3: protos.notification.push.ios.Request.Payload
	(*Request_Message)(nil),     // 4: protos.notification.push.ios.Request.Message
	nil,                         // 5: protos.notification.push.ios.Request.Payload.CustomEntry
	(*timestamp.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_notification_push_ios_proto_depIdxs = []int32{
	4, // 0: protos.notification.push.ios.Request.payload:type_name -> protos.notification.push.ios.Request.Message
	6, // 1: protos.notification.push.ios.Request.event_timestamp:type_name -> google.protobuf.Timestamp
	5, // 2: protos.notification.push.ios.Request.Payload.custom:type_name -> protos.notification.push.ios.Request.Payload.CustomEntry
	3, // 3: protos.notification.push.ios.Request.Message.payload:type_name -> protos.notification.push.ios.Request.Payload
	0, // 4: protos.notification.push.ios.Request.Message.priority:type_name -> protos.notification.push.ios.Request.Priority
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_notification_push_ios_proto_init() }
func file_notification_push_ios_proto_init() {
	if File_notification_push_ios_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notification_push_ios_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_notification_push_ios_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_notification_push_ios_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Payload); i {
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
		file_notification_push_ios_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Message); i {
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
			RawDescriptor: file_notification_push_ios_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notification_push_ios_proto_goTypes,
		DependencyIndexes: file_notification_push_ios_proto_depIdxs,
		EnumInfos:         file_notification_push_ios_proto_enumTypes,
		MessageInfos:      file_notification_push_ios_proto_msgTypes,
	}.Build()
	File_notification_push_ios_proto = out.File
	file_notification_push_ios_proto_rawDesc = nil
	file_notification_push_ios_proto_goTypes = nil
	file_notification_push_ios_proto_depIdxs = nil
}