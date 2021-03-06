// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: notification/slack/slack.proto

package slack

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request_SeverityLevel int32

const (
	Request_INFO  Request_SeverityLevel = 0
	Request_WARN  Request_SeverityLevel = 1
	Request_ERROR Request_SeverityLevel = 2
)

// Enum value maps for Request_SeverityLevel.
var (
	Request_SeverityLevel_name = map[int32]string{
		0: "INFO",
		1: "WARN",
		2: "ERROR",
	}
	Request_SeverityLevel_value = map[string]int32{
		"INFO":  0,
		"WARN":  1,
		"ERROR": 2,
	}
)

func (x Request_SeverityLevel) Enum() *Request_SeverityLevel {
	p := new(Request_SeverityLevel)
	*p = x
	return p
}

func (x Request_SeverityLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Request_SeverityLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_notification_slack_slack_proto_enumTypes[0].Descriptor()
}

func (Request_SeverityLevel) Type() protoreflect.EnumType {
	return &file_notification_slack_slack_proto_enumTypes[0]
}

func (x Request_SeverityLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Request_SeverityLevel.Descriptor instead.
func (Request_SeverityLevel) EnumDescriptor() ([]byte, []int) {
	return file_notification_slack_slack_proto_rawDescGZIP(), []int{0, 0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guid           string                 `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Title          string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Source         string                 `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Webhook        string                 `protobuf:"bytes,4,opt,name=webhook,proto3" json:"webhook,omitempty"`
	Level          Request_SeverityLevel  `protobuf:"varint,5,opt,name=level,proto3,enum=protos.notification.slack.Request_SeverityLevel" json:"level,omitempty"`
	Payload        *Request_Message       `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
	EventTimestamp *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=event_timestamp,json=eventTimestamp,proto3" json:"event_timestamp,omitempty"`
	ServiceToken   string                 `protobuf:"bytes,8,opt,name=service_token,json=serviceToken,proto3" json:"service_token,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_slack_slack_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_notification_slack_slack_proto_msgTypes[0]
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
	return file_notification_slack_slack_proto_rawDescGZIP(), []int{0}
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

func (x *Request) GetWebhook() string {
	if x != nil {
		return x.Webhook
	}
	return ""
}

func (x *Request) GetLevel() Request_SeverityLevel {
	if x != nil {
		return x.Level
	}
	return Request_INFO
}

func (x *Request) GetPayload() *Request_Message {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Request) GetEventTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.EventTimestamp
	}
	return nil
}

func (x *Request) GetServiceToken() string {
	if x != nil {
		return x.ServiceToken
	}
	return ""
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
		mi := &file_notification_slack_slack_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_notification_slack_slack_proto_msgTypes[1]
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
	return file_notification_slack_slack_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Request_Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Short bool   `protobuf:"varint,3,opt,name=short,proto3" json:"short,omitempty"`
}

func (x *Request_Field) Reset() {
	*x = Request_Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_slack_slack_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Field) ProtoMessage() {}

func (x *Request_Field) ProtoReflect() protoreflect.Message {
	mi := &file_notification_slack_slack_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Field.ProtoReflect.Descriptor instead.
func (*Request_Field) Descriptor() ([]byte, []int) {
	return file_notification_slack_slack_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Request_Field) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Request_Field) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Request_Field) GetShort() bool {
	if x != nil {
		return x.Short
	}
	return false
}

type Request_Attachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color     string           `protobuf:"bytes,1,opt,name=color,proto3" json:"color,omitempty"`
	Fallback  string           `protobuf:"bytes,2,opt,name=fallback,proto3" json:"fallback,omitempty"`
	Title     string           `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	TitleLink string           `protobuf:"bytes,4,opt,name=title_link,json=titleLink,proto3" json:"title_link,omitempty"`
	Pretext   string           `protobuf:"bytes,5,opt,name=pretext,proto3" json:"pretext,omitempty"`
	Text      string           `protobuf:"bytes,6,opt,name=text,proto3" json:"text,omitempty"`
	Fields    []*Request_Field `protobuf:"bytes,7,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *Request_Attachment) Reset() {
	*x = Request_Attachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_slack_slack_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Attachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Attachment) ProtoMessage() {}

func (x *Request_Attachment) ProtoReflect() protoreflect.Message {
	mi := &file_notification_slack_slack_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Attachment.ProtoReflect.Descriptor instead.
func (*Request_Attachment) Descriptor() ([]byte, []int) {
	return file_notification_slack_slack_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Request_Attachment) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Request_Attachment) GetFallback() string {
	if x != nil {
		return x.Fallback
	}
	return ""
}

func (x *Request_Attachment) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Request_Attachment) GetTitleLink() string {
	if x != nil {
		return x.TitleLink
	}
	return ""
}

func (x *Request_Attachment) GetPretext() string {
	if x != nil {
		return x.Pretext
	}
	return ""
}

func (x *Request_Attachment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Request_Attachment) GetFields() []*Request_Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

type Request_Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text        string                `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Attachments []*Request_Attachment `protobuf:"bytes,2,rep,name=attachments,proto3" json:"attachments,omitempty"`
	Channel     string                `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
	Username    string                `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	IconEmoji   string                `protobuf:"bytes,5,opt,name=icon_emoji,json=iconEmoji,proto3" json:"icon_emoji,omitempty"`
	IconUrl     string                `protobuf:"bytes,6,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url,omitempty"`
}

func (x *Request_Message) Reset() {
	*x = Request_Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_slack_slack_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Message) ProtoMessage() {}

func (x *Request_Message) ProtoReflect() protoreflect.Message {
	mi := &file_notification_slack_slack_proto_msgTypes[4]
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
	return file_notification_slack_slack_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Request_Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Request_Message) GetAttachments() []*Request_Attachment {
	if x != nil {
		return x.Attachments
	}
	return nil
}

func (x *Request_Message) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *Request_Message) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Request_Message) GetIconEmoji() string {
	if x != nil {
		return x.IconEmoji
	}
	return ""
}

func (x *Request_Message) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

var File_notification_slack_slack_proto protoreflect.FileDescriptor

var file_notification_slack_slack_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73,
	0x6c, 0x61, 0x63, 0x6b, 0x2f, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x07, 0x0a,
	0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x46, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x44, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x43, 0x0a, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x49, 0x0a, 0x05,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x1a, 0xe3, 0x01, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x72, 0x65, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x72, 0x65, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x40, 0x0a, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0xde, 0x01,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x4f, 0x0a,
	0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x65, 0x6d, 0x6f,
	0x6a, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x63, 0x6f, 0x6e, 0x45, 0x6d,
	0x6f, 0x6a, 0x69, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x22, 0x2e,
	0x0a, 0x0d, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x41, 0x52,
	0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x22, 0x20,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x42, 0x14, 0x5a, 0x12, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notification_slack_slack_proto_rawDescOnce sync.Once
	file_notification_slack_slack_proto_rawDescData = file_notification_slack_slack_proto_rawDesc
)

func file_notification_slack_slack_proto_rawDescGZIP() []byte {
	file_notification_slack_slack_proto_rawDescOnce.Do(func() {
		file_notification_slack_slack_proto_rawDescData = protoimpl.X.CompressGZIP(file_notification_slack_slack_proto_rawDescData)
	})
	return file_notification_slack_slack_proto_rawDescData
}

var file_notification_slack_slack_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_notification_slack_slack_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_notification_slack_slack_proto_goTypes = []interface{}{
	(Request_SeverityLevel)(0),    // 0: protos.notification.slack.Request.SeverityLevel
	(*Request)(nil),               // 1: protos.notification.slack.Request
	(*Response)(nil),              // 2: protos.notification.slack.Response
	(*Request_Field)(nil),         // 3: protos.notification.slack.Request.Field
	(*Request_Attachment)(nil),    // 4: protos.notification.slack.Request.Attachment
	(*Request_Message)(nil),       // 5: protos.notification.slack.Request.Message
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_notification_slack_slack_proto_depIdxs = []int32{
	0, // 0: protos.notification.slack.Request.level:type_name -> protos.notification.slack.Request.SeverityLevel
	5, // 1: protos.notification.slack.Request.payload:type_name -> protos.notification.slack.Request.Message
	6, // 2: protos.notification.slack.Request.event_timestamp:type_name -> google.protobuf.Timestamp
	3, // 3: protos.notification.slack.Request.Attachment.fields:type_name -> protos.notification.slack.Request.Field
	4, // 4: protos.notification.slack.Request.Message.attachments:type_name -> protos.notification.slack.Request.Attachment
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_notification_slack_slack_proto_init() }
func file_notification_slack_slack_proto_init() {
	if File_notification_slack_slack_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notification_slack_slack_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_notification_slack_slack_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_notification_slack_slack_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Field); i {
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
		file_notification_slack_slack_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Attachment); i {
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
		file_notification_slack_slack_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_notification_slack_slack_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notification_slack_slack_proto_goTypes,
		DependencyIndexes: file_notification_slack_slack_proto_depIdxs,
		EnumInfos:         file_notification_slack_slack_proto_enumTypes,
		MessageInfos:      file_notification_slack_slack_proto_msgTypes,
	}.Build()
	File_notification_slack_slack_proto = out.File
	file_notification_slack_slack_proto_rawDesc = nil
	file_notification_slack_slack_proto_goTypes = nil
	file_notification_slack_slack_proto_depIdxs = nil
}
