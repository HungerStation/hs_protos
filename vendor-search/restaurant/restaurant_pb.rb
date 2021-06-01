# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: restaurant/restaurant.proto

require 'google/protobuf'

require 'google/protobuf/timestamp_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("restaurant/restaurant.proto", :syntax => :proto3) do
    add_message "Restaurant" do
      optional :metadata, :message, 1, "Restaurant.Metadata"
      optional :id, :int32, 2
      optional :name, :string, 3
      optional :created_at, :message, 4, "google.protobuf.Timestamp"
      optional :updated_at, :message, 5, "google.protobuf.Timestamp"
      optional :enabled, :bool, 6
      optional :logo_url, :string, 7
      optional :country_id, :int32, 8
      optional :name_ar, :string, 9
      optional :minordervalue, :float, 10
      optional :contact, :string, 11
      optional :description, :string, 12
      optional :description_ar, :string, 13
      optional :slug, :string, 14
      optional :featured, :bool, 15
      optional :weight, :int32, 16
      optional :company_id, :int32, 17
      optional :working_time_id, :int32, 18
      optional :billing_group_id, :int32, 19
      optional :external_service_host_id, :int32, 20
      optional :sub_working_time_id, :int32, 21
      optional :logo_ar_url, :string, 22
      optional :grade, :double, 23
      optional :bias, :int32, 24
      optional :background_url, :string, 25
      optional :sponsorship, :int32, 26
      optional :sla_dashboard, :bool, 27
      optional :cover_photo_url, :string, 28
      optional :slug_ar, :string, 29
      optional :vat_applied, :bool, 30
      optional :brand_id, :int32, 31
      optional :store_type_id, :int32, 32
      optional :ios_cover_photo_url, :string, 33
      optional :android_cover_photo_url, :string, 34
      optional :enabled_reason, :int32, 35
      optional :restaurant_type, :int32, 36
      optional :rate_average, :float, 37
      optional :rate_count, :int32, 38
      optional :vertical_id, :int32, 39
      optional :boolean, :bool, 40
    end
    add_message "Restaurant.Metadata" do
      optional :guid, :string, 1
      optional :source, :string, 2
      optional :action, :enum, 3, "Restaurant.Metadata.ActionType"
      optional :event_timestamp, :message, 4, "google.protobuf.Timestamp"
    end
    add_enum "Restaurant.Metadata.ActionType" do
      value :CREATE, 0
      value :UPDATE, 1
      value :DELETE, 2
    end
  end
end

module Pubsub
  module Protos
    module VendorSearch
      Restaurant = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("Restaurant").msgclass
      Restaurant::Metadata = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("Restaurant.Metadata").msgclass
      Restaurant::Metadata::ActionType = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("Restaurant.Metadata.ActionType").enummodule
    end
  end
end
