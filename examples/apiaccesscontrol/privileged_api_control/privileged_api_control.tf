// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "authz_compartment_id" {
}
variable "privileged_api_control_resources" {
}
variable "notification_topic_id" {
}

variable "privileged_api_control_privileged_operation_list_api_name" {
}

variable "privileged_api_control_privileged_operation_list_entity_type" {
}

variable "privileged_api_control_approver_group_id_list" {
  default = ["use_iam_policy"]
}

variable "privileged_api_control_defined_tags_value" {
  default = "value"
}

variable "privileged_api_control_description" {
  default = "Control for pre approving the apis"
}

variable "privileged_api_control_display_name" {
  default = "TestPrivilegedApiControl"
}

variable "privileged_api_control_freeform_tags" {
  default = { "Department" = "db" }
}

variable "privileged_api_control_id" {
  default = "id"
}

variable "privileged_api_control_number_of_approvers" {
  default = 1
}

variable "privileged_api_control_privileged_operation_list_attribute_names" {
  default = []
}


variable "privileged_api_control_resource_type" {
  default = "EXADATAINFRASTRUCTURE"
}

variable "privileged_api_control_state" {
  default = "ACTIVE"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_apiaccesscontrol_privileged_api_control" "test_privileged_api_control" {
  #Required
  approver_group_id_list = var.privileged_api_control_approver_group_id_list
  compartment_id         = var.authz_compartment_id
  notification_topic_id  = var.notification_topic_id
  privileged_operation_list {
    #Required
    api_name = var.privileged_api_control_privileged_operation_list_api_name
    entity_type     = var.privileged_api_control_privileged_operation_list_entity_type
    #Optional
    attribute_names = var.privileged_api_control_privileged_operation_list_attribute_names
  }
  resource_type = var.privileged_api_control_resource_type
  resources     = var.privileged_api_control_resources
  number_of_approvers = var.privileged_api_control_number_of_approvers
  description         = var.privileged_api_control_description
  display_name        = var.privileged_api_control_display_name
  #Optional
  freeform_tags       = var.privileged_api_control_freeform_tags
}

data "oci_apiaccesscontrol_privileged_api_controls" "test_privileged_api_controls" {
  compartment_id = var.authz_compartment_id
  #Optional
  display_name   = var.privileged_api_control_display_name
  id             = var.privileged_api_control_id
  resource_type  = var.privileged_api_control_resource_type
  state          = var.privileged_api_control_state
}

