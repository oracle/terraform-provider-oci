// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "operator_control_approver_groups_list" {
  default = []
}

variable "operator_control_approvers_list" {
  default = []
}

variable "operator_control_defined_tags_value" {
  default = "definedTags"
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

resource "oci_operator_access_control_operator_control" "test_operator_control" {
  #Required
  compartment_id        = var.compartment_id
  operator_control_name = "tfexample-opctl51"

  #Optional
  approver_groups_list        = [data.oci_identity_groups.get_admin_approver_group.groups[0].id]
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

data "oci_operator_access_control_operator_controls" "test_operator_controls" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name  = var.operator_control_display_name
  resource_type = var.operator_control_resource_type
  state         = var.operator_control_state
}
