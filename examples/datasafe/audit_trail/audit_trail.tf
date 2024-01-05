// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "trail_ocid" {}

variable "audit_trail_access_level" {
  default = "RESTRICTED"
}

variable "audit_trail_compartment_id_in_subtree" {
  default = false
}

variable "audit_trail_defined_tags_value" {
  default = "value"
}

variable "audit_trail_description" {
  default = "updated-description"
}

variable "audit_trail_display_name" {
  default = "updated-name"
}

variable "audit_trail_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "audit_trail_is_auto_purge_enabled" {
  default = false
}

variable "audit_trail_state" {
  default = "AVAILABLE"
}

variable "audit_trail_status" {
  default = "STARTING"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_audit_trail" "test_audit_trail" {
  #Required
  audit_trail_id = var.trail_ocid

  #Optional
  description           = var.audit_trail_description
  display_name          = var.audit_trail_display_name
  freeform_tags         = var.audit_trail_freeform_tags
  is_auto_purge_enabled = var.audit_trail_is_auto_purge_enabled
}

data "oci_data_safe_audit_trails" "test_audit_trails" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  audit_trail_id            = oci_data_safe_audit_trail.test_audit_trail.id
}

