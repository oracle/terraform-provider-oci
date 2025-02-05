// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "sensitive_types_export_is_include_all_sensitive_types" {
  default = true
}

variable "sensitive_types_export_display_name" {
  default = "displayName"
}

variable "sensitive_types_export_description" {
  default = "description"
}

variable "sensitive_types_export_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "sensitive_types_export_access_level" {
  default = "ACCESSIBLE"
}

variable "sensitive_types_export_compartment_id_in_subtree" {
  default = true
}

variable "sensitive_types_export_state" {
  default = "ACTIVE"
}

variable "sensitive_types_export_time_created_less_than" {
  default = "2038-01-01T00:00:00.000Z"
}

resource "oci_data_safe_sensitive_types_export" "test_sensitive_types_export" {
  #Required
  compartment_id = var.compartment_ocid
  is_include_all_sensitive_types = var.sensitive_types_export_is_include_all_sensitive_types

  #Optional
  display_name = var.sensitive_types_export_display_name
  description = var.sensitive_types_export_description
  freeform_tags = var.sensitive_types_export_freeform_tags
}

data "oci_data_safe_sensitive_types_exports" "test_sensitive_types_exports" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  access_level = var.sensitive_types_export_access_level
  compartment_id_in_subtree = var.sensitive_types_export_compartment_id_in_subtree
  display_name = var.sensitive_types_export_display_name
  state = var.sensitive_types_export_state
  time_created_less_than = var.sensitive_types_export_time_created_less_than
}

data "oci_data_safe_sensitive_types_export" "test_sensitive_types_export" {
  sensitive_types_export_id = oci_data_safe_sensitive_types_export.test_sensitive_types_export.id
}