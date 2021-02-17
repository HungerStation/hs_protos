// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: notification/push/android/android.proto

package android

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

type Request_Priority int32

const (
	Request_NORMAL Request_Priority = 0
	Request_HIGH   Request_Priority = 1
)

// Enum value maps for Request_Priority.
var (
	Request_Priority_name = map[int32]string{
		0: "NORMAL",
		1: "HIGH",
	}
	Request_Priority_value = map[string]int32{
		"NORMAL": 0,
		"HIGH":   1,
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
	return file_notification_push_android_android_proto_enumTypes[0].Descriptor()
}

func (Request_Priority) Type() protoreflect.EnumType {
	return &file_notification_push_android_android_proto_enumTypes[0]
}

func (x Request_Priority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Request_Priority.Descriptor instead.
func (Request_Priority) EnumDescriptor() ([]byte, []int) {
	return file_notification_push_android_android_proto_rawDescGZIP(), []int{0, 0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guid   string `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	// device tokens/recipient
	Tokens         []string               `protobuf:"bytes,4,rep,name=tokens,proto3" json:"tokens,omitempty"`
	Payload        *Request_Message       `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	EventTimestamp *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=event_timestamp,json=eventTimestamp,proto3" json:"event_timestamp,omitempty"`
	ServiceToken   string                 `protobuf:"bytes,7,opt,name=service_token,json=serviceToken,proto3" json:"service_token,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_push_android_android_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_notification_push_android_android_proto_msgTypes[0]
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
	return file_notification_push_android_android_proto_rawDescGZIP(), []int{0}
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
		mi := &file_notification_push_android_android_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_notification_push_android_android_proto_msgTypes[1]
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
	return file_notification_push_android_android_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Request_Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic       string            `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	CollapseKey string            `protobuf:"bytes,2,opt,name=collapse_key,json=collapseKey,proto3" json:"collapse_key,omitempty"`
	Priority    Request_Priority  `protobuf:"varint,3,opt,name=priority,proto3,enum=protos.notification.push.android.Request_Priority" json:"priority,omitempty"`
	Data        map[string]string `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Request_Message) Reset() {
	*x = Request_Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_push_android_android_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Message) ProtoMessage() {}

func (x *Request_Message) ProtoReflect() protoreflect.Message {
	mi := &file_notification_push_android_android_proto_msgTypes[2]
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
	return file_notification_push_android_android_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Request_Message) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Request_Message) GetCollapseKey() string {
	if x != nil {
		return x.CollapseKey
	}
	return ""
}

func (x *Request_Message) GetPriority() Request_Priority {
	if x != nil {
		return x.Priority
	}
	return Request_NORMAL
}

func (x *Request_Message) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_notification_push_android_android_proto protoreflect.FileDescriptor

var file_notification_push_android_android_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70,
	0x75, 0x73, 0x68, 0x2f, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x2f, 0x61, 0x6e, 0x64, 0x72,
	0x6f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x75, 0x73, 0x68, 0x2e, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x04, 0x0a,
	0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x12, 0x4b, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x61,
	0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x43, 0x0a, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x9c, 0x02, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6f, 0x6c, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x4e,
	0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x61, 0x6e, 0x64, 0x72,
	0x6f, 0x69, 0x64, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x4f,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x20, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x48, 0x49, 0x47, 0x48, 0x10, 0x01, 0x22, 0x20, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x1b, 0x5a, 0x19,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x75, 0x73,
	0x68, 0x2f, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_notification_push_android_android_proto_rawDescOnce sync.Once
	file_notification_push_android_android_proto_rawDescData = file_notification_push_android_android_proto_rawDesc
)

func file_notification_push_android_android_proto_rawDescGZIP() []byte {
	file_notification_push_android_android_proto_rawDescOnce.Do(func() {
		file_notification_push_android_android_proto_rawDescData = protoimpl.X.CompressGZIP(file_notification_push_android_android_proto_rawDescData)
	})
	return file_notification_push_android_android_proto_rawDescData
}

var file_notification_push_android_android_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_notification_push_android_android_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_notification_push_android_android_proto_goTypes = []interface{}{
	(Request_Priority)(0),         // 0: protos.notification.push.android.Request.Priority
	(*Request)(nil),               // 1: protos.notification.push.android.Request
	(*Response)(nil),              // 2: protos.notification.push.android.Response
	(*Request_Message)(nil),       // 3: protos.notification.push.android.Request.Message
	nil,                           // 4: protos.notification.push.android.Request.Message.DataEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_notification_push_android_android_proto_depIdxs = []int32{
	3, // 0: protos.notification.push.android.Request.payload:type_name -> protos.notification.push.android.Request.Message
	5, // 1: protos.notification.push.android.Request.event_timestamp:type_name -> google.protobuf.Timestamp
	0, // 2: protos.notification.push.android.Request.Message.priority:type_name -> protos.notification.push.android.Request.Priority
	4, // 3: protos.notification.push.android.Request.Message.data:type_name -> protos.notification.push.android.Request.Message.DataEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_notification_push_android_android_proto_init() }
func file_notification_push_android_android_proto_init() {
	if File_notification_push_android_android_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notification_push_android_android_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_notification_push_android_android_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_notification_push_android_android_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_notification_push_android_android_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notification_push_android_android_proto_goTypes,
		DependencyIndexes: file_notification_push_android_android_proto_depIdxs,
		EnumInfos:         file_notification_push_android_android_proto_enumTypes,
		MessageInfos:      file_notification_push_android_android_proto_msgTypes,
	}.Build()
	File_notification_push_android_android_proto = out.File
	file_notification_push_android_android_proto_rawDesc = nil
	file_notification_push_android_android_proto_goTypes = nil
	file_notification_push_android_android_proto_depIdxs = nil
}
