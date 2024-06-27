// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "test_notification_topic_id"{
}

variable "delegation_control_defined_tags_value" {
  default = "value"
}

variable "delegation_control_delegation_subscription_ids" {
}

variable "delegation_control_description" {
  default = "description"
}

variable "delegation_control_display_name" {
  default = "displayName"
}

variable "delegation_control_freeform_tags" {
  default = { "Department" = "Finance2" }
}

variable "delegation_control_is_auto_approve_during_maintenance" {
  default = false
}

variable "delegation_control_notification_message_format" {
  default = "JSON"
}

variable "delegation_control_num_approvals_required" {
  default = 1
}

variable "delegation_control_pre_approved_service_provider_action_names" {
  default = ["DLGT_MGMT_SYS_DIAG", "DLGT_MGMT_DBAAS_API_ACCESS"]
}

variable "delegation_control_resource_ids" {
}

variable "delegation_control_resource_type" {
  default = "VMCLUSTER"
}

variable "delegation_control_state" {
  default = "AVAILABLE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_delegation_management_delegation_control" "test_delegation_control" {
  #Required
  compartment_id              = var.compartment_id
  delegation_subscription_ids = var.delegation_control_delegation_subscription_ids
  display_name                = var.delegation_control_display_name
  notification_message_format = var.delegation_control_notification_message_format
  notification_topic_id       = var.test_notification_topic_id
  resource_ids                = var.delegation_control_resource_ids
  resource_type               = var.delegation_control_resource_type

  #Optional
  #defined_tags                               = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.delegation_control_defined_tags_value)
  description                                = var.delegation_control_description
  freeform_tags                              = var.delegation_control_freeform_tags
  is_auto_approve_during_maintenance         = var.delegation_control_is_auto_approve_during_maintenance
  num_approvals_required                     = var.delegation_control_num_approvals_required
  pre_approved_service_provider_action_names = var.delegation_control_pre_approved_service_provider_action_names
  #vault_id                                   = oci_kms_vault.test_vault.id
  #vault_key_id                               = oci_kms_key.test_key.id
}

data "oci_delegation_management_delegation_controls" "test_delegation_controls" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  #display_name  = var.delegation_control_display_name
  #resource_id   = oci_usage_proxy_resource.test_resource.id
  #resource_type = var.delegation_control_resource_type
  #state         = var.delegation_control_state
}