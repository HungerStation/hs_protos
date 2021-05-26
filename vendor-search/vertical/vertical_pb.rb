# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protobuf/vendor-search/vertical/vertical.proto

require 'google/protobuf'

require 'google/protobuf/timestamp_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("protobuf/vendor-search/vertical/vertical.proto", :syntax => :proto3) do
    add_message "Vertical" do
      optional :metadata, :message, 1, "Vertical.Metadata"
      optional :id, :uint32, 2
      optional :vertical_key, :string, 3
      optional :title, :string, 4
      optional :title_ar, :string, 5
      optional :description, :string, 6
      optional :weight, :uint32, 7
      optional :enabled, :bool, 8
      optional :created_at, :message, 9, "google.protobuf.Timestamp"
      optional :updated_at, :message, 10, "google.protobuf.Timestamp"
      optional :category, :uint32, 11
      optional :salesforce_name, :uint32, 12
    end
    add_message "Vertical.Metadata" do
      optional :guid, :string, 1
      optional :source, :string, 2
      optional :action, :enum, 3, "Vertical.Metadata.ActionType"
      optional :event_timestamp, :message, 4, "google.protobuf.Timestamp"
    end
    add_enum "Vertical.Metadata.ActionType" do
      value :CREATE, 0
      value :UPDATE, 1
      value :DELETE, 2
    end
  end
end

Vertical = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("Vertical").msgclass
Vertical::Metadata = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("Vertical.Metadata").msgclass
Vertical::Metadata::ActionType = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("Vertical.Metadata.ActionType").enummodule