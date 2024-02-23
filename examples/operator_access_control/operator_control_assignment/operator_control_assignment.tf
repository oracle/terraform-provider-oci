// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "operator_control_approvers_list" {
  default = []
}

variable "operator_control_assignment_comment" {
  default = "comment"
}

variable "operator_control_assignment_defined_tags_value" {
  default = "definedTags"
}

variable "operator_control_assignment_freeform_tags" {
  default = "freeformTags"
}

variable "operator_control_assignment_is_auto_approve_during_maintenance" {
  default = false
}

variable "operator_control_assignment_is_enforced_always" {
  default = true
}

variable "operator_control_assignment_is_log_forwarded" {
  default = false
}

variable "operator_control_assignment_is_hypervisor_log_forwarded" {
  default = false
}

variable "operator_control_assignment_remote_syslog_server_address" {
  default = "remoteSyslogServerAddress"
}

variable "operator_control_assignment_remote_syslog_server_ca_cert" {
  default = "remoteSyslogServerCACert"
}

variable "operator_control_assignment_remote_syslog_server_port" {
  default = 10
}

variable "operator_control_assignment_resource_name" {
  default = "resourceName"
}

variable "operator_control_assignment_resource_type" {
  default = "EXADATAINFRASTRUCTURE"
}

variable "operator_control_assignment_state" {
  default = "CREATED"
}

variable "operator_control_assignment_time_assignment_from" {
  default = "timeAssignmentFrom"
}

variable "operator_control_assignment_time_assignment_to" {
  default = "timeAssignmentTo"
}

variable "operator_control_description" {
  default = "description"
}

variable "operator_control_display_name" {
  default = "displayName"
}

variable "operator_control_email_id_list" {
  default = []
}

variable "operator_control_freeform_tags" {
  default = "freeformTags"
}

variable "operator_control_is_fully_pre_approved" {
  default = false
}

variable "operator_control_pre_approved_op_action_list" {
  default = []
}

variable "operator_control_resource_type" {
  default = "EXADATAINFRASTRUCTURE"
}

variable "operator_control_state" {
  default = "CREATED"
}

variable "operator_control_system_message" {
  default = "systemMessage"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_identity_groups" "get_admin_approver_group" {
    #Required
    compartment_id = var.tenancy_ocid

    #Optional
    name = "Administrators"
}

resource "random_string" "random_suffix" {
  length  = 9
  special = false
  upper   = false
}

resource "oci_operator_access_control_operator_control" "test_operator_control" {
  #Required
  compartment_id        = var.compartment_id
  operator_control_name = "tfexample-${random_string.random_suffix.result}"

  #Optional
  #approver_groups_list        = [data.oci_identity_groups.get_admin_approver_group.groups[0].id]
  approver_groups_list        = ["use_iam_policy"]
  approvers_list              = var.operator_control_approvers_list
  #defined_tags                = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.operator_control_defined_tags_value)
  description                 = var.operator_control_description
  email_id_list               = var.operator_control_email_id_list
  #freeform_tags               = var.operator_control_freeform_tags
  is_fully_pre_approved       = var.operator_control_is_fully_pre_approved
  pre_approved_op_action_list = var.operator_control_pre_approved_op_action_list
  resource_type               = var.operator_control_resource_type
  system_message              = var.operator_control_system_message
}

resource "oci_operator_access_control_operator_control_assignment" "test_operator_control_assignment" {
  #Required
  compartment_id          = var.compartment_id
  is_enforced_always      = var.operator_control_assignment_is_enforced_always
  operator_control_id     = oci_operator_access_control_operator_control.test_operator_control.id
  resource_compartment_id = var.compartment_id
  resource_id             = "ocid1.exadatainfrastructure.test..${random_string.random_suffix.result}"
  resource_name           = var.operator_control_assignment_resource_name
  resource_type           = var.operator_control_assignment_resource_type

  #Optional
  comment                            = var.operator_control_assignment_comment
  #defined_tags                       = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.operator_control_assignment_defined_tags_value)
  #freeform_tags                      = var.operator_control_assignment_freeform_tags
  is_auto_approve_during_maintenance = var.operator_control_assignment_is_auto_approve_during_maintenance
  is_log_forwarded                   = var.operator_control_assignment_is_log_forwarded
  is_hypervisor_log_forwarded        = var.operator_control_assignment_is_hypervisor_log_forwarded
  remote_syslog_server_address       = var.operator_control_assignment_remote_syslog_server_address
  remote_syslog_server_ca_cert       = var.operator_control_assignment_remote_syslog_server_ca_cert
  remote_syslog_server_port          = var.operator_control_assignment_remote_syslog_server_port
  #time_assignment_from               = "2006-01-02T15:04:05Z"
  #time_assignment_to                 = "2022-01-02T15:04:05Z"
}

data "oci_operator_access_control_operator_control_assignments" "test_operator_control_assignments" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  operator_control_name = oci_operator_access_control_operator_control.test_operator_control.operator_control_name
  resource_name         = var.operator_control_assignment_resource_name
  resource_type         = var.operator_control_assignment_resource_type
  state                 = var.operator_control_assignment_state
}

