// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "profile_ocid" {}

variable "audit_profile_access_level" {
  default = "RESTRICTED"
}

variable "audit_profile_audit_collected_volume_greater_than_or_equal_to" {
  default = 10
}

variable "audit_profile_compartment_id_in_subtree" {
  default = false
}

variable "audit_profile_defined_tags_value" {
  default = "value"
}

variable "audit_profile_description" {
  default = "updated-description"
}

variable "audit_profile_display_name" {
  default = "Audit_updated"
}

variable "audit_profile_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "audit_profile_is_override_global_retention_setting" {
  default = false
}

variable "audit_profile_is_paid_usage_enabled" {
  default = false
}

variable "audit_profile_state" {
  default = "AVAILABLE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_audit_profile" "test_audit_profile" {
  #Required
  audit_profile_id = var.profile_ocid

  #Optional
  description           = var.audit_profile_description
  display_name          = var.audit_profile_display_name
  freeform_tags         = var.audit_profile_freeform_tags
  is_paid_usage_enabled = var.audit_profile_is_paid_usage_enabled
}

data "oci_data_safe_audit_profile" "test_audit_profile" {
  #Optional
  audit_profile_id                                = var.profile_ocid
}

