// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_id" {
}

variable "alarm_body" {
  default = "High CPU utilization reached"
}

variable "alarm_compartment_id_in_subtree" {
  default = false
}

variable "alarm_defined_tags_value" {
  default = "value"
}

variable "alarm_destinations" {
  default = []
}

variable "alarm_display_name" {
  default = "High CPU Utilization"
}

variable "alarm_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "alarm_is_enabled" {
  default = false
}

variable "alarm_message_format" {
  default = "ONS_OPTIMIZED"
}

variable "alarm_metric_compartment_id_in_subtree" {
  default = false
}

variable "alarm_namespace" {
  default = "oci_computeagent"
}

variable "alarm_pending_duration" {
  default = "PT5M"
}

variable "alarm_query" {
  default = "CpuUtilization[10m].percentile(0.9) < 85"
}

variable "alarm_repeat_notification_duration" {
  default = "PT2H"
}

variable "alarm_resolution" {
  default = "1m"
}

variable "alarm_resource_group" {
  default = "resourceGroup"
}

variable "alarm_severity" {
  default = "WARNING"
}

variable "alarm_state" {
  default = "ACTIVE"
}

variable "alarm_suppression_description" {
  default = "System Maintenance"
}

variable "alarm_suppression_time_suppress_from" {
  default = "2029-02-01T18:00:00.000Z"
}

variable "alarm_suppression_time_suppress_until" {
  default = "2029-02-01T19:00:00.000Z"
}

variable "alarm_history_collection_alarm_historytype" {
  default = "STATE_TRANSITION_HISTORY"
}

variable "alarm_history_collection_timestamp_greater_than_or_equal_to" {
  default = "2028-12-01T01:00:00.001Z"
}

variable "alarm_history_collection_timestamp_less_than" {
  default = "2035-01-01T01:00:00.001Z"
}

variable "alarm_status_compartment_id_in_subtree" {
  default = false
}

variable "alarm_status_display_name" {
  default = "High CPU Utilization"
}

variable "tag_namespace_description" {
  default = "Just a test"
}

variable "tag_namespace_name" {
  default = "testexamples-tag-namespace"
}

variable "is_notifications_per_metric_dimension_enabled" {
  default = false
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "random_string" "topicname" {
  length  = 10
  special = false
}

resource "oci_ons_notification_topic" "test_notification_topic" {
  #Required
  compartment_id = var.compartment_id
  name           = random_string.topicname.result
}

resource "oci_monitoring_alarm" "test_alarm" {
  #Required
  compartment_id        = var.compartment_id
  destinations          = [oci_ons_notification_topic.test_notification_topic.id]
  display_name          = var.alarm_display_name
  is_enabled            = var.alarm_is_enabled
  metric_compartment_id = var.compartment_id
  namespace             = var.alarm_namespace
  query                 = var.alarm_query
  severity              = var.alarm_severity

  #Optional
  body                             = var.alarm_body
  message_format                   = var.alarm_message_format
  metric_compartment_id_in_subtree = var.alarm_metric_compartment_id_in_subtree
  pending_duration                 = var.alarm_pending_duration
  repeat_notification_duration     = var.alarm_repeat_notification_duration
  resolution                       = var.alarm_resolution
  resource_group                   = var.alarm_resource_group
  is_notifications_per_metric_dimension_enabled = var.is_notifications_per_metric_dimension_enabled

  suppression {
    #Required
    time_suppress_from  = var.alarm_suppression_time_suppress_from
    time_suppress_until = var.alarm_suppression_time_suppress_until

    #Optional
    description = var.alarm_suppression_description
  }
}

data "oci_monitoring_alarms" "test_alarms" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  compartment_id_in_subtree = var.alarm_compartment_id_in_subtree
  display_name              = var.alarm_display_name
  state                     = var.alarm_state
}

data "oci_monitoring_alarm_history_collection" "test_alarm_history_collection" {
  #Required
  alarm_id = oci_monitoring_alarm.test_alarm.id

  #Optional
  alarm_historytype                  = var.alarm_history_collection_alarm_historytype
  timestamp_greater_than_or_equal_to = var.alarm_history_collection_timestamp_greater_than_or_equal_to
  timestamp_less_than                = var.alarm_history_collection_timestamp_less_than
}

data "oci_monitoring_alarm_statuses" "test_alarm_statuses" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  compartment_id_in_subtree = var.alarm_status_compartment_id_in_subtree
  display_name              = var.alarm_status_display_name
}

data "oci_monitoring_alarm" "test_alarm" {
  #Required
  alarm_id = oci_monitoring_alarm.test_alarm.id
}

