// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: vendor-search/vertical/vertical.proto

package vertical

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

type Vertical_Metadata_ActionType int32

const (
	Vertical_Metadata_CREATE Vertical_Metadata_ActionType = 0
	Vertical_Metadata_UPDATE Vertical_Metadata_ActionType = 1
	Vertical_Metadata_DELETE Vertical_Metadata_ActionType = 2
)

// Enum value maps for Vertical_Metadata_ActionType.
var (
	Vertical_Metadata_ActionType_name = map[int32]string{
		0: "CREATE",
		1: "UPDATE",
		2: "DELETE",
	}
	Vertical_Metadata_ActionType_value = map[string]int32{
		"CREATE": 0,
		"UPDATE": 1,
		"DELETE": 2,
	}
)

func (x Vertical_Metadata_ActionType) Enum() *Vertical_Metadata_ActionType {
	p := new(Vertical_Metadata_ActionType)
	*p = x
	return p
}

func (x Vertical_Metadata_ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Vertical_Metadata_ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_vendor_search_vertical_vertical_proto_enumTypes[0].Descriptor()
}

func (Vertical_Metadata_ActionType) Type() protoreflect.EnumType {
	return &file_vendor_search_vertical_vertical_proto_enumTypes[0]
}

func (x Vertical_Metadata_ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Vertical_Metadata_ActionType.Descriptor instead.
func (Vertical_Metadata_ActionType) EnumDescriptor() ([]byte, []int) {
	return file_vendor_search_vertical_vertical_proto_rawDescGZIP(), []int{0, 0, 0}
}

type Vertical struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata       *Vertical_Metadata     `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Id             uint32                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	VerticalKey    string                 `protobuf:"bytes,3,opt,name=vertical_key,json=verticalKey,proto3" json:"vertical_key,omitempty"`
	Title          string                 `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	TitleAr        string                 `protobuf:"bytes,5,opt,name=title_ar,json=titleAr,proto3" json:"title_ar,omitempty"`
	Description    string                 `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Weight         int32                  `protobuf:"zigzag32,7,opt,name=weight,proto3" json:"weight,omitempty"`
	Enabled        bool                   `protobuf:"varint,8,opt,name=enabled,proto3" json:"enabled,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Category       uint32                 `protobuf:"varint,11,opt,name=category,proto3" json:"category,omitempty"`
	SalesforceName string                 `protobuf:"bytes,12,opt,name=salesforce_name,json=salesforceName,proto3" json:"salesforce_name,omitempty"`
}

func (x *Vertical) Reset() {
	*x = Vertical{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vendor_search_vertical_vertical_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vertical) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vertical) ProtoMessage() {}

func (x *Vertical) ProtoReflect() protoreflect.Message {
	mi := &file_vendor_search_vertical_vertical_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vertical.ProtoReflect.Descriptor instead.
func (*Vertical) Descriptor() ([]byte, []int) {
	return file_vendor_search_vertical_vertical_proto_rawDescGZIP(), []int{0}
}

func (x *Vertical) GetMetadata() *Vertical_Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Vertical) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Vertical) GetVerticalKey() string {
	if x != nil {
		return x.VerticalKey
	}
	return ""
}

func (x *Vertical) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Vertical) GetTitleAr() string {
	if x != nil {
		return x.TitleAr
	}
	return ""
}

func (x *Vertical) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Vertical) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Vertical) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Vertical) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Vertical) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Vertical) GetCategory() uint32 {
	if x != nil {
		return x.Category
	}
	return 0
}

func (x *Vertical) GetSalesforceName() string {
	if x != nil {
		return x.SalesforceName
	}
	return ""
}

type Vertical_Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guid           string                       `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Source         string                       `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Action         Vertical_Metadata_ActionType `protobuf:"varint,3,opt,name=action,proto3,enum=Vertical_Metadata_ActionType" json:"action,omitempty"`
	EventTimestamp *timestamppb.Timestamp       `protobuf:"bytes,4,opt,name=event_timestamp,json=eventTimestamp,proto3" json:"event_timestamp,omitempty"`
}

func (x *Vertical_Metadata) Reset() {
	*x = Vertical_Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vendor_search_vertical_vertical_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vertical_Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vertical_Metadata) ProtoMessage() {}

func (x *Vertical_Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_vendor_search_vertical_vertical_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vertical_Metadata.ProtoReflect.Descriptor instead.
func (*Vertical_Metadata) Descriptor() ([]byte, []int) {
	return file_vendor_search_vertical_vertical_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Vertical_Metadata) GetGuid() string {
	if x != nil {
		return x.Guid
	}
	return ""
}

func (x *Vertical_Metadata) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Vertical_Metadata) GetAction() Vertical_Metadata_ActionType {
	if x != nil {
		return x.Action
	}
	return Vertical_Metadata_CREATE
}

func (x *Vertical_Metadata) GetEventTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.EventTimestamp
	}
	return nil
}

var File_vendor_search_vertical_vertical_proto protoreflect.FileDescriptor

var file_vendor_search_vertical_vertical_proto_rawDesc = []byte{
	0x0a, 0x25, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f,
	0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x2f, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x05, 0x0a, 0x08, 0x56, 0x65, 0x72,
	0x74, 0x69, 0x63, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x56, 0x65, 0x72, 0x74, 0x69, 0x63,
	0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61,
	0x6c, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x65, 0x72,
	0x74, 0x69, 0x63, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x41, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x27, 0x0a, 0x0f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x66,
	0x6f, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0xe4, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x35, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1d, 0x2e, 0x56, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x0f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x30, 0x0a,
	0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x42,
	0x40, 0x5a, 0x16, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2f, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0xea, 0x02, 0x25, 0x50, 0x75, 0x62, 0x73,
	0x75, 0x62, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x3a, 0x3a,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vendor_search_vertical_vertical_proto_rawDescOnce sync.Once
	file_vendor_search_vertical_vertical_proto_rawDescData = file_vendor_search_vertical_vertical_proto_rawDesc
)

func file_vendor_search_vertical_vertical_proto_rawDescGZIP() []byte {
	file_vendor_search_vertical_vertical_proto_rawDescOnce.Do(func() {
		file_vendor_search_vertical_vertical_proto_rawDescData = protoimpl.X.CompressGZIP(file_vendor_search_vertical_vertical_proto_rawDescData)
	})
	return file_vendor_search_vertical_vertical_proto_rawDescData
}

var file_vendor_search_vertical_vertical_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_vendor_search_vertical_vertical_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_vendor_search_vertical_vertical_proto_goTypes = []interface{}{
	(Vertical_Metadata_ActionType)(0), // 0: Vertical.Metadata.ActionType
	(*Vertical)(nil),                  // 1: Vertical
	(*Vertical_Metadata)(nil),         // 2: Vertical.Metadata
	(*timestamppb.Timestamp)(nil),     // 3: google.protobuf.Timestamp
}
var file_vendor_search_vertical_vertical_proto_depIdxs = []int32{
	2, // 0: Vertical.metadata:type_name -> Vertical.Metadata
	3, // 1: Vertical.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: Vertical.updated_at:type_name -> google.protobuf.Timestamp
	0, // 3: Vertical.Metadata.action:type_name -> Vertical.Metadata.ActionType
	3, // 4: Vertical.Metadata.event_timestamp:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_vendor_search_vertical_vertical_proto_init() }
func file_vendor_search_vertical_vertical_proto_init() {
	if File_vendor_search_vertical_vertical_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vendor_search_vertical_vertical_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vertical); i {
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
		file_vendor_search_vertical_vertical_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vertical_Metadata); i {
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
			RawDescriptor: file_vendor_search_vertical_vertical_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vendor_search_vertical_vertical_proto_goTypes,
		DependencyIndexes: file_vendor_search_vertical_vertical_proto_depIdxs,
		EnumInfos:         file_vendor_search_vertical_vertical_proto_enumTypes,
		MessageInfos:      file_vendor_search_vertical_vertical_proto_msgTypes,
	}.Build()
	File_vendor_search_vertical_vertical_proto = out.File
	file_vendor_search_vertical_vertical_proto_rawDesc = nil
	file_vendor_search_vertical_vertical_proto_goTypes = nil
	file_vendor_search_vertical_vertical_proto_depIdxs = nil
}
