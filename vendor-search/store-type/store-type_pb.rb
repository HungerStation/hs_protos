# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: vendor-search/store-type/store-type.proto

require 'google/protobuf'

require 'google/protobuf/timestamp_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("vendor-search/store-type/store-type.proto", :syntax => :proto3) do
    add_message "StoreType" do
      optional :metadata, :message, 1, "StoreType.Metadata"
      optional :id, :uint32, 2
      optional :store_type, :uint32, 3
      optional :name, :string, 4
      optional :name_ar, :string, 5
      optional :enabled, :bool, 6
      optional :weight, :uint32, 7
      optional :view_style, :uint32, 8
      optional :description, :string, 9
      optional :description_ar, :string, 10
      optional :created_at, :message, 11, "google.protobuf.Timestamp"
      optional :updated_at, :message, 12, "google.protobuf.Timestamp"
      optional :background_image_url, :string, 13
      optional :unavailable_description, :string, 14
      optional :unavailable_description_ar, :string, 15
      optional :coming_soon_description, :string, 16
      optional :coming_soon_description_ar, :string, 17
      optional :action_label_en, :string, 18
      optional :action_label_ar, :string, 19
    end
    add_message "StoreType.Metadata" do
      optional :guid, :string, 1
      optional :source, :string, 2
      optional :action, :enum, 3, "StoreType.Metadata.ActionType"
      optional :event_timestamp, :message, 4, "google.protobuf.Timestamp"
    end
    add_enum "StoreType.Metadata.ActionType" do
      value :CREATE, 0
      value :UPDATE, 1
      value :DELETE, 2
    end
  end
end

module PubsubSynchronizer
  module Protobufs
    module Protos
      StoreType = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("StoreType").msgclass
      StoreType::Metadata = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("StoreType.Metadata").msgclass
      StoreType::Metadata::ActionType = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("StoreType.Metadata.ActionType").enummodule
    end
  end
end
