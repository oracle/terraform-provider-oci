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

variable "sensitive_type_group_display_name" {
  default = "displayName"
}

variable "sensitive_type_group_description" {
  default = "description"
}

variable "sensitive_type_group_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "sensitive_type_group_access_level" {
  default = "ACCESSIBLE"
}

variable "sensitive_type_group_compartment_id_in_subtree" {
  default = true
}

variable "sensitive_type_group_state" {
  default = "ACTIVE"
}

variable "sensitive_type_group_time_created_less_than" {
  default = "2038-01-01T00:00:00.000Z"
}

resource "oci_data_safe_sensitive_type_group" "test_sensitive_type_group" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = var.sensitive_type_group_display_name
  description = var.sensitive_type_group_description
  freeform_tags = var.sensitive_type_group_freeform_tags
}

data "oci_data_safe_sensitive_type_groups" "test_sensitive_type_groups" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  access_level = var.sensitive_type_group_access_level
  compartment_id_in_subtree = var.sensitive_type_group_compartment_id_in_subtree
  display_name = var.sensitive_type_group_display_name
  state = var.sensitive_type_group_state
  time_created_less_than = var.sensitive_type_group_time_created_less_than
}

data "oci_data_safe_sensitive_type_group" "test_sensitive_type_group" {
  sensitive_type_group_id = oci_data_safe_sensitive_type_group.test_sensitive_type_group.id
}