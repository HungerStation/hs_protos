// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: notification/slack/slack.proto

package notification_slack

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

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

	Guid           string                `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Title          string                `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Source         string                `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Webhook        string                `protobuf:"bytes,4,opt,name=webhook,proto3" json:"webhook,omitempty"`
	Level          Request_SeverityLevel `protobuf:"varint,5,opt,name=level,proto3,enum=notification.slack.Request_SeverityLevel" json:"level,omitempty"`
	Payload        map[string]string     `protobuf:"bytes,6,rep,name=payload,proto3" json:"payload,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	EventTimestamp *timestamp.Timestamp  `protobuf:"bytes,7,opt,name=event_timestamp,json=eventTimestamp,proto3" json:"event_timestamp,omitempty"`
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

func (x *Request) GetPayload() map[string]string {
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

var File_notification_slack_slack_proto protoreflect.FileDescriptor

var file_notification_slack_slack_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73,
	0x6c, 0x61, 0x63, 0x6b, 0x2f, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73,
	0x6c, 0x61, 0x63, 0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x03, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67,
	0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x3f, 0x0a, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x42, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x6c, 0x61,
	0x63, 0x6b, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x43, 0x0a, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x1a, 0x3a, 0x0a, 0x0c, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x2e, 0x0a, 0x0d, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x57, 0x41, 0x52, 0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x02, 0x22, 0x20, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x32, 0x72, 0x0a, 0x0c, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x1b, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x6c, 0x61, 0x63, 0x6b,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14,
	0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x6c, 0x61, 0x63, 0x6b, 0x3a, 0x01, 0x2a, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_notification_slack_slack_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_notification_slack_slack_proto_goTypes = []interface{}{
	(Request_SeverityLevel)(0),  // 0: notification.slack.Request.SeverityLevel
	(*Request)(nil),             // 1: notification.slack.Request
	(*Response)(nil),            // 2: notification.slack.Response
	nil,                         // 3: notification.slack.Request.PayloadEntry
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_notification_slack_slack_proto_depIdxs = []int32{
	0, // 0: notification.slack.Request.level:type_name -> notification.slack.Request.SeverityLevel
	3, // 1: notification.slack.Request.payload:type_name -> notification.slack.Request.PayloadEntry
	4, // 2: notification.slack.Request.event_timestamp:type_name -> google.protobuf.Timestamp
	1, // 3: notification.slack.SlackService.Send:input_type -> notification.slack.Request
	2, // 4: notification.slack.SlackService.Send:output_type -> notification.slack.Response
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_notification_slack_slack_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
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
