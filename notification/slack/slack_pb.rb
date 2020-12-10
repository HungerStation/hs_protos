# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: notification/slack/slack.proto

require 'google/protobuf'

require 'google/protobuf/timestamp_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("notification/slack/slack.proto", :syntax => :proto3) do
    add_message "protos.notification.slack.Request" do
      optional :guid, :string, 1
      optional :title, :string, 2
      optional :source, :string, 3
      optional :webhook, :string, 4
      optional :level, :enum, 5, "protos.notification.slack.Request.SeverityLevel"
      optional :payload, :message, 6, "protos.notification.slack.Request.Message"
      optional :event_timestamp, :message, 7, "google.protobuf.Timestamp"
    end
    add_message "protos.notification.slack.Request.Field" do
      optional :title, :string, 1
      optional :value, :string, 2
      optional :short, :bool, 3
    end
    add_message "protos.notification.slack.Request.Attachment" do
      optional :color, :string, 1
      optional :fallback, :string, 2
      optional :title, :string, 3
      optional :title_link, :string, 4
      optional :pretext, :string, 5
      optional :text, :string, 6
      repeated :fields, :message, 7, "protos.notification.slack.Request.Field"
    end
    add_message "protos.notification.slack.Request.Message" do
      optional :text, :string, 1
      repeated :attachments, :message, 2, "protos.notification.slack.Request.Attachment"
    end
    add_enum "protos.notification.slack.Request.SeverityLevel" do
      value :INFO, 0
      value :WARN, 1
      value :ERROR, 2
    end
    add_message "protos.notification.slack.Response" do
      optional :error, :string, 1
    end
  end
end

module Protos
  module Notification
    module Slack
      Request = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("protos.notification.slack.Request").msgclass
      Request::Field = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("protos.notification.slack.Request.Field").msgclass
      Request::Attachment = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("protos.notification.slack.Request.Attachment").msgclass
      Request::Message = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("protos.notification.slack.Request.Message").msgclass
      Request::SeverityLevel = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("protos.notification.slack.Request.SeverityLevel").enummodule
      Response = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("protos.notification.slack.Response").msgclass
    end
  end
end
