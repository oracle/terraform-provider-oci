// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


locals {
  alarm_suppression_from_offset_in_days = 50
  alarm_suppression_to_offset_in_days =  51
}


variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "alarm_id" {
}

variable "dimensions" {
  default = {
    key1 = "value1"
    key2 = "value2"
  }
}

variable "display_name" {
  default =  "test suppression"
}

variable "time_suppress_from" {

}

variable "time_suppress_until" {

}

variable "defined_tags" {
  default = "value"
}

variable "description" {
  default = "A description for the suppression."
}

variable "freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}


resource "oci_monitoring_alarm_suppression" "test_alarm_suppression" {
  #Required
  alarm_suppression_target {
    alarm_id     = oci_monitoring_alarm.test_suppression_alarm.id
    target_type = "ALARM"
  }
  dimensions = var.dimensions
  display_name         = var.display_name
  time_suppress_from   = timeadd(timestamp(),"${local.alarm_suppression_from_offset_in_days * 24}h")
  time_suppress_until  = timeadd(timestamp(), "${local.alarm_suppression_to_offset_in_days * 24}h")

  #Optional
  description = var.description
  freeform_tags = var.freeform_tags
}

variable "compartment_id" {
}

variable "alarm_display_name" {
  default = "High CPU Utilization"
}


resource "random_string" "topicname" {
  length  = 10
  special = false
}

resource "oci_ons_notification_topic" "test_alarm_notification_topic" {
  #Required
  compartment_id = var.compartment_id
  name           = random_string.topicname.result
}


variable "alarm_is_enabled" {
  default = false
}

variable "alarm_namespace" {
  default = "oci_computeagent"
}

variable "alarm_query" {
  default = "CpuUtilization[10m].percentile(0.9) < 85"
}

variable "alarm_severity" {
  default = "WARNING"
}

variable "is_notifications_per_metric_dimension_enabled" {
  default = true
}

resource "oci_monitoring_alarm" "test_suppression_alarm" {
  #Required
  compartment_id        = var.compartment_id
  destinations          = [oci_ons_notification_topic.test_alarm_notification_topic.id]
  display_name          = var.alarm_display_name
  is_enabled            = var.alarm_is_enabled
  metric_compartment_id = var.compartment_id
  namespace             = var.alarm_namespace
  query                 = var.alarm_query
  severity              = var.alarm_severity
  is_notifications_per_metric_dimension_enabled = var.is_notifications_per_metric_dimension_enabled
}