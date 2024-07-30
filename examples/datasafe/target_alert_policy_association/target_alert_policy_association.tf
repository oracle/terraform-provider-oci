// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "data_safe_alert_policy_ocid"{}
variable "data_safe_target_ocid" {}

variable "target_alert_policy_association_description" {
  default = "description"
}

variable "target_alert_policy_association_display_name" {
  default = "displayName"
}

variable "target_alert_policy_association_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "target_alert_policy_association_is_enabled" {
  default = false
}

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_target_alert_policy_association" "test_target_alert_policy_association" {
  #Required
  compartment_id = var.compartment_ocid
  is_enabled     = var.target_alert_policy_association_is_enabled
  policy_id      = var.data_safe_alert_policy_ocid
  target_id      = var.data_safe_target_ocid

  #Optional
  description   = var.target_alert_policy_association_description
  display_name  = var.target_alert_policy_association_display_name
}

data "oci_data_safe_target_alert_policy_associations" "test_target_alert_policy_associations" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  alert_policy_id = var.data_safe_alert_policy_ocid
  target_alert_policy_association_id = oci_data_safe_target_alert_policy_association.test_target_alert_policy_association.id
  target_id = var.data_safe_target_ocid
}
