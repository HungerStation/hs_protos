// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: vendor-search/home-module/home-module.proto

package home_module

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

type HomeModule_Metadata_ActionType int32

const (
	HomeModule_Metadata_CREATE HomeModule_Metadata_ActionType = 0
	HomeModule_Metadata_UPDATE HomeModule_Metadata_ActionType = 1
	HomeModule_Metadata_DELETE HomeModule_Metadata_ActionType = 2
)

// Enum value maps for HomeModule_Metadata_ActionType.
var (
	HomeModule_Metadata_ActionType_name = map[int32]string{
		0: "CREATE",
		1: "UPDATE",
		2: "DELETE",
	}
	HomeModule_Metadata_ActionType_value = map[string]int32{
		"CREATE": 0,
		"UPDATE": 1,
		"DELETE": 2,
	}
)

func (x HomeModule_Metadata_ActionType) Enum() *HomeModule_Metadata_ActionType {
	p := new(HomeModule_Metadata_ActionType)
	*p = x
	return p
}

func (x HomeModule_Metadata_ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HomeModule_Metadata_ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_vendor_search_home_module_home_module_proto_enumTypes[0].Descriptor()
}

func (HomeModule_Metadata_ActionType) Type() protoreflect.EnumType {
	return &file_vendor_search_home_module_home_module_proto_enumTypes[0]
}

func (x HomeModule_Metadata_ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HomeModule_Metadata_ActionType.Descriptor instead.
func (HomeModule_Metadata_ActionType) EnumDescriptor() ([]byte, []int) {
	return file_vendor_search_home_module_home_module_proto_rawDescGZIP(), []int{0, 0, 0}
}

type HomeModule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata                  *HomeModule_Metadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Id                        uint32                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Slug                      string                 `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Redirection               uint32                 `protobuf:"varint,4,opt,name=redirection,proto3" json:"redirection,omitempty"`
	Enabled                   bool                   `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Weight                    int32                  `protobuf:"zigzag32,6,opt,name=weight,proto3" json:"weight,omitempty"`
	RenderFilter              bool                   `protobuf:"varint,7,opt,name=render_filter,json=renderFilter,proto3" json:"render_filter,omitempty"`
	SwimlaneConfig            string                 `protobuf:"bytes,8,opt,name=swimlane_config,json=swimlaneConfig,proto3" json:"swimlane_config,omitempty"`
	Title                     string                 `protobuf:"bytes,9,opt,name=title,proto3" json:"title,omitempty"`
	TitleAr                   string                 `protobuf:"bytes,10,opt,name=title_ar,json=titleAr,proto3" json:"title_ar,omitempty"`
	Description               string                 `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	CoverImage_3X2Url         string                 `protobuf:"bytes,12,opt,name=cover_image_3x2_url,json=coverImage3x2Url,proto3" json:"cover_image_3x2_url,omitempty"`
	CoverImage_3X3Url         string                 `protobuf:"bytes,13,opt,name=cover_image_3x3_url,json=coverImage3x3Url,proto3" json:"cover_image_3x3_url,omitempty"`
	CoverImage_4X3Url         string                 `protobuf:"bytes,14,opt,name=cover_image_4x3_url,json=coverImage4x3Url,proto3" json:"cover_image_4x3_url,omitempty"`
	CoverImage_4X4Url         string                 `protobuf:"bytes,15,opt,name=cover_image_4x4_url,json=coverImage4x4Url,proto3" json:"cover_image_4x4_url,omitempty"`
	CoverImage_7X2Url         string                 `protobuf:"bytes,16,opt,name=cover_image_7x2_url,json=coverImage7x2Url,proto3" json:"cover_image_7x2_url,omitempty"`
	CoverImage_7X3Url         string                 `protobuf:"bytes,17,opt,name=cover_image_7x3_url,json=coverImage7x3Url,proto3" json:"cover_image_7x3_url,omitempty"`
	CoverImage_7X4Url         string                 `protobuf:"bytes,18,opt,name=cover_image_7x4_url,json=coverImage7x4Url,proto3" json:"cover_image_7x4_url,omitempty"`
	DisabledCoverImage_3X2Url string                 `protobuf:"bytes,19,opt,name=disabled_cover_image_3x2_url,json=disabledCoverImage3x2Url,proto3" json:"disabled_cover_image_3x2_url,omitempty"`
	DisabledCoverImage_3X3Url string                 `protobuf:"bytes,20,opt,name=disabled_cover_image_3x3_url,json=disabledCoverImage3x3Url,proto3" json:"disabled_cover_image_3x3_url,omitempty"`
	DisabledCoverImage_4X3Url string                 `protobuf:"bytes,21,opt,name=disabled_cover_image_4x3_url,json=disabledCoverImage4x3Url,proto3" json:"disabled_cover_image_4x3_url,omitempty"`
	DisabledCoverImage_4X4Url string                 `protobuf:"bytes,22,opt,name=disabled_cover_image_4x4_url,json=disabledCoverImage4x4Url,proto3" json:"disabled_cover_image_4x4_url,omitempty"`
	DisabledCoverImage_7X2Url string                 `protobuf:"bytes,23,opt,name=disabled_cover_image_7x2_url,json=disabledCoverImage7x2Url,proto3" json:"disabled_cover_image_7x2_url,omitempty"`
	DisabledCoverImage_7X3Url string                 `protobuf:"bytes,24,opt,name=disabled_cover_image_7x3_url,json=disabledCoverImage7x3Url,proto3" json:"disabled_cover_image_7x3_url,omitempty"`
	DisabledCoverImage_7X4Url string                 `protobuf:"bytes,25,opt,name=disabled_cover_image_7x4_url,json=disabledCoverImage7x4Url,proto3" json:"disabled_cover_image_7x4_url,omitempty"`
	CreatedAt                 *timestamppb.Timestamp `protobuf:"bytes,26,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt                 *timestamppb.Timestamp `protobuf:"bytes,27,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Configs                   []string               `protobuf:"bytes,28,rep,name=configs,proto3" json:"configs,omitempty"`
}

func (x *HomeModule) Reset() {
	*x = HomeModule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vendor_search_home_module_home_module_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeModule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeModule) ProtoMessage() {}

func (x *HomeModule) ProtoReflect() protoreflect.Message {
	mi := &file_vendor_search_home_module_home_module_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeModule.ProtoReflect.Descriptor instead.
func (*HomeModule) Descriptor() ([]byte, []int) {
	return file_vendor_search_home_module_home_module_proto_rawDescGZIP(), []int{0}
}

func (x *HomeModule) GetMetadata() *HomeModule_Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *HomeModule) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HomeModule) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *HomeModule) GetRedirection() uint32 {
	if x != nil {
		return x.Redirection
	}
	return 0
}

func (x *HomeModule) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *HomeModule) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *HomeModule) GetRenderFilter() bool {
	if x != nil {
		return x.RenderFilter
	}
	return false
}

func (x *HomeModule) GetSwimlaneConfig() string {
	if x != nil {
		return x.SwimlaneConfig
	}
	return ""
}

func (x *HomeModule) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *HomeModule) GetTitleAr() string {
	if x != nil {
		return x.TitleAr
	}
	return ""
}

func (x *HomeModule) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *HomeModule) GetCoverImage_3X2Url() string {
	if x != nil {
		return x.CoverImage_3X2Url
	}
	return ""
}

func (x *HomeModule) GetCoverImage_3X3Url() string {
	if x != nil {
		return x.CoverImage_3X3Url
	}
	return ""
}

func (x *HomeModule) GetCoverImage_4X3Url() string {
	if x != nil {
		return x.CoverImage_4X3Url
	}
	return ""
}

func (x *HomeModule) GetCoverImage_4X4Url() string {
	if x != nil {
		return x.CoverImage_4X4Url
	}
	return ""
}

func (x *HomeModule) GetCoverImage_7X2Url() string {
	if x != nil {
		return x.CoverImage_7X2Url
	}
	return ""
}

func (x *HomeModule) GetCoverImage_7X3Url() string {
	if x != nil {
		return x.CoverImage_7X3Url
	}
	return ""
}

func (x *HomeModule) GetCoverImage_7X4Url() string {
	if x != nil {
		return x.CoverImage_7X4Url
	}
	return ""
}

func (x *HomeModule) GetDisabledCoverImage_3X2Url() string {
	if x != nil {
		return x.DisabledCoverImage_3X2Url
	}
	return ""
}

func (x *HomeModule) GetDisabledCoverImage_3X3Url() string {
	if x != nil {
		return x.DisabledCoverImage_3X3Url
	}
	return ""
}

func (x *HomeModule) GetDisabledCoverImage_4X3Url() string {
	if x != nil {
		return x.DisabledCoverImage_4X3Url
	}
	return ""
}

func (x *HomeModule) GetDisabledCoverImage_4X4Url() string {
	if x != nil {
		return x.DisabledCoverImage_4X4Url
	}
	return ""
}

func (x *HomeModule) GetDisabledCoverImage_7X2Url() string {
	if x != nil {
		return x.DisabledCoverImage_7X2Url
	}
	return ""
}

func (x *HomeModule) GetDisabledCoverImage_7X3Url() string {
	if x != nil {
		return x.DisabledCoverImage_7X3Url
	}
	return ""
}

func (x *HomeModule) GetDisabledCoverImage_7X4Url() string {
	if x != nil {
		return x.DisabledCoverImage_7X4Url
	}
	return ""
}

func (x *HomeModule) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *HomeModule) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *HomeModule) GetConfigs() []string {
	if x != nil {
		return x.Configs
	}
	return nil
}

type HomeModule_Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guid           string                         `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Source         string                         `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Action         HomeModule_Metadata_ActionType `protobuf:"varint,3,opt,name=action,proto3,enum=HomeModule_Metadata_ActionType" json:"action,omitempty"`
	EventTimestamp *timestamppb.Timestamp         `protobuf:"bytes,4,opt,name=event_timestamp,json=eventTimestamp,proto3" json:"event_timestamp,omitempty"`
}

func (x *HomeModule_Metadata) Reset() {
	*x = HomeModule_Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vendor_search_home_module_home_module_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeModule_Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeModule_Metadata) ProtoMessage() {}

func (x *HomeModule_Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_vendor_search_home_module_home_module_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeModule_Metadata.ProtoReflect.Descriptor instead.
func (*HomeModule_Metadata) Descriptor() ([]byte, []int) {
	return file_vendor_search_home_module_home_module_proto_rawDescGZIP(), []int{0, 0}
}

func (x *HomeModule_Metadata) GetGuid() string {
	if x != nil {
		return x.Guid
	}
	return ""
}

func (x *HomeModule_Metadata) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *HomeModule_Metadata) GetAction() HomeModule_Metadata_ActionType {
	if x != nil {
		return x.Action
	}
	return HomeModule_Metadata_CREATE
}

func (x *HomeModule_Metadata) GetEventTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.EventTimestamp
	}
	return nil
}

var File_vendor_search_home_module_home_module_proto protoreflect.FileDescriptor

var file_vendor_search_home_module_home_module_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f,
	0x68, 0x6f, 0x6d, 0x65, 0x2d, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x68, 0x6f, 0x6d, 0x65,
	0x2d, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9,
	0x0b, 0x0a, 0x0a, 0x48, 0x6f, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x30, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x6c, 0x75, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x11, 0x52,
	0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f,
	0x73, 0x77, 0x69, 0x6d, 0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x77, 0x69, 0x6d, 0x6c, 0x61, 0x6e, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x5f, 0x61, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x41, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x13, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x33, 0x78, 0x32, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x33, 0x78, 0x32, 0x55, 0x72, 0x6c, 0x12, 0x2d, 0x0a, 0x13, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x33, 0x78, 0x33, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x33, 0x78, 0x33, 0x55, 0x72, 0x6c, 0x12, 0x2d, 0x0a, 0x13, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x34, 0x78, 0x33, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x34,
	0x78, 0x33, 0x55, 0x72, 0x6c, 0x12, 0x2d, 0x0a, 0x13, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x5f, 0x34, 0x78, 0x34, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x34, 0x78,
	0x34, 0x55, 0x72, 0x6c, 0x12, 0x2d, 0x0a, 0x13, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x37, 0x78, 0x32, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x37, 0x78, 0x32,
	0x55, 0x72, 0x6c, 0x12, 0x2d, 0x0a, 0x13, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x37, 0x78, 0x33, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x37, 0x78, 0x33, 0x55,
	0x72, 0x6c, 0x12, 0x2d, 0x0a, 0x13, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x37, 0x78, 0x34, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x37, 0x78, 0x34, 0x55, 0x72,
	0x6c, 0x12, 0x3e, 0x0a, 0x1c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x33, 0x78, 0x32, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x33, 0x78, 0x32, 0x55, 0x72,
	0x6c, 0x12, 0x3e, 0x0a, 0x1c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x33, 0x78, 0x33, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x33, 0x78, 0x33, 0x55, 0x72,
	0x6c, 0x12, 0x3e, 0x0a, 0x1c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x34, 0x78, 0x33, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x34, 0x78, 0x33, 0x55, 0x72,
	0x6c, 0x12, 0x3e, 0x0a, 0x1c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x34, 0x78, 0x34, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x34, 0x78, 0x34, 0x55, 0x72,
	0x6c, 0x12, 0x3e, 0x0a, 0x1c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x37, 0x78, 0x32, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x37, 0x78, 0x32, 0x55, 0x72,
	0x6c, 0x12, 0x3e, 0x0a, 0x1c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x37, 0x78, 0x33, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x37, 0x78, 0x33, 0x55, 0x72,
	0x6c, 0x12, 0x3e, 0x0a, 0x1c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x37, 0x78, 0x34, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x37, 0x78, 0x34, 0x55, 0x72,
	0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x1a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x18, 0x1c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x1a, 0xe6, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x75,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x48, 0x6f, 0x6d,
	0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x30, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x42, 0x4f, 0x5a, 0x25, 0x76, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x68, 0x6f, 0x6d, 0x65,
	0x2d, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x3b, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0xea, 0x02, 0x25, 0x50, 0x75, 0x62, 0x73, 0x75, 0x62, 0x53, 0x79, 0x6e, 0x63,
	0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x73, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_vendor_search_home_module_home_module_proto_rawDescOnce sync.Once
	file_vendor_search_home_module_home_module_proto_rawDescData = file_vendor_search_home_module_home_module_proto_rawDesc
)

func file_vendor_search_home_module_home_module_proto_rawDescGZIP() []byte {
	file_vendor_search_home_module_home_module_proto_rawDescOnce.Do(func() {
		file_vendor_search_home_module_home_module_proto_rawDescData = protoimpl.X.CompressGZIP(file_vendor_search_home_module_home_module_proto_rawDescData)
	})
	return file_vendor_search_home_module_home_module_proto_rawDescData
}

var file_vendor_search_home_module_home_module_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_vendor_search_home_module_home_module_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_vendor_search_home_module_home_module_proto_goTypes = []interface{}{
	(HomeModule_Metadata_ActionType)(0), // 0: HomeModule.Metadata.ActionType
	(*HomeModule)(nil),                  // 1: HomeModule
	(*HomeModule_Metadata)(nil),         // 2: HomeModule.Metadata
	(*timestamppb.Timestamp)(nil),       // 3: google.protobuf.Timestamp
}
var file_vendor_search_home_module_home_module_proto_depIdxs = []int32{
	2, // 0: HomeModule.metadata:type_name -> HomeModule.Metadata
	3, // 1: HomeModule.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: HomeModule.updated_at:type_name -> google.protobuf.Timestamp
	0, // 3: HomeModule.Metadata.action:type_name -> HomeModule.Metadata.ActionType
	3, // 4: HomeModule.Metadata.event_timestamp:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_vendor_search_home_module_home_module_proto_init() }
func file_vendor_search_home_module_home_module_proto_init() {
	if File_vendor_search_home_module_home_module_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vendor_search_home_module_home_module_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeModule); i {
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
		file_vendor_search_home_module_home_module_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeModule_Metadata); i {
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
			RawDescriptor: file_vendor_search_home_module_home_module_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vendor_search_home_module_home_module_proto_goTypes,
		DependencyIndexes: file_vendor_search_home_module_home_module_proto_depIdxs,
		EnumInfos:         file_vendor_search_home_module_home_module_proto_enumTypes,
		MessageInfos:      file_vendor_search_home_module_home_module_proto_msgTypes,
	}.Build()
	File_vendor_search_home_module_home_module_proto = out.File
	file_vendor_search_home_module_home_module_proto_rawDesc = nil
	file_vendor_search_home_module_home_module_proto_goTypes = nil
	file_vendor_search_home_module_home_module_proto_depIdxs = nil
}
