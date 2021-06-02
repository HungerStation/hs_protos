# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: vendor-search/week-time-rule/week-time-rule.proto

require 'google/protobuf'

require 'google/protobuf/timestamp_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("vendor-search/week-time-rule/week-time-rule.proto", :syntax => :proto3) do
    add_message "WeekTimeRule" do
      optional :metadata, :message, 1, "WeekTimeRule.Metadata"
      optional :id, :uint32, 2
      optional :name, :string, 3
      optional :restaurant_id, :uint32, 4
      optional :created_at, :message, 5, "google.protobuf.Timestamp"
      optional :updated_at, :message, 6, "google.protobuf.Timestamp"
    end
    add_message "WeekTimeRule.Metadata" do
      optional :guid, :string, 1
      optional :source, :string, 2
      optional :action, :enum, 3, "WeekTimeRule.Metadata.ActionType"
      optional :event_timestamp, :message, 4, "google.protobuf.Timestamp"
    end
    add_enum "WeekTimeRule.Metadata.ActionType" do
      value :CREATE, 0
      value :UPDATE, 1
      value :DELETE, 2
    end
  end
end

module PubsubSynchronizer
  module Protobufs
    module Protos
      WeekTimeRule = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("WeekTimeRule").msgclass
      WeekTimeRule::Metadata = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("WeekTimeRule.Metadata").msgclass
      WeekTimeRule::Metadata::ActionType = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("WeekTimeRule.Metadata.ActionType").enummodule
    end
  end
end
