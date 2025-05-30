// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "authz_compartment_id" {
}
variable "privileged_api_control_privileged_operation_list_api_name" {
}
variable "notification_topic_id" {
}
variable "resource_id" {
}

variable "privileged_api_request_defined_tags_value" {
  default = "value"
}

variable "privileged_api_request_display_name" {
  default = "displayName"
}

variable "privileged_api_request_duration_in_hrs" {
  default = 1
}

variable "privileged_api_request_freeform_tags" {
  default = { "Department" = "db" }
}

variable "privileged_api_request_id" {
  default = "id"
}

variable "privileged_api_request_privileged_operation_list_attribute_names" {
  default = []
}

variable "privileged_api_request_reason_detail" {
  default = "reasonDetail"
}

variable "privileged_api_request_reason_summary" {
  default = "TestPrivilegedApiControl"
}

variable "privileged_api_request_resource_type" {
  default = "EXADATAINFRASTRUCTURE"
}

variable "privileged_api_request_severity" {
  default = "SEV_3"
}

variable "privileged_api_request_state" {
  default = "APPROVED"
}

variable "privileged_api_request_sub_resource_name_list" {
  default = []
}

variable "privileged_api_request_ticket_numbers" {
  default = []
}

variable "privileged_api_request_time_requested_for_future_access" {
  default = ""
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_apiaccesscontrol_privileged_api_request" "test_privileged_api_request" {
  #Required
  compartment_id                   = var.authz_compartment_id
  privileged_operation_list {
    #Required
    api_name = var.privileged_api_control_privileged_operation_list_api_name

    #Optional
    attribute_names = var.privileged_api_request_privileged_operation_list_attribute_names
  }
  reason_summary = var.privileged_api_request_reason_summary
  resource_id    = var.resource_id

  #Optional
  duration_in_hrs                  = var.privileged_api_request_duration_in_hrs
  freeform_tags                    = var.privileged_api_request_freeform_tags
  notification_topic_id            = var.notification_topic_id
  reason_detail                    = var.privileged_api_request_reason_detail
  severity                         = var.privileged_api_request_severity
  sub_resource_name_list           = var.privileged_api_request_sub_resource_name_list
  ticket_numbers                   = var.privileged_api_request_ticket_numbers
  time_requested_for_future_access = var.privileged_api_request_time_requested_for_future_access
}

data "oci_apiaccesscontrol_privileged_api_requests" "test_privileged_api_requests" {
  compartment_id = var.authz_compartment_id
  #Optional
  display_name   = var.privileged_api_request_display_name
  id             = var.privileged_api_request_id
  resource_id    = var.resource_id
  resource_type  = var.privileged_api_request_resource_type
  state          = var.privileged_api_request_state
}

