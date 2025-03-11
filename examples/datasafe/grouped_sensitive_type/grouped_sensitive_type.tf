// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "data_safe_target_ocid" {}

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

variable "sensitive_type_entity_type" {
  default = "SENSITIVE_TYPE"
}

variable "sensitive_type_display_name" {
  default = "displayName"
}

resource "oci_data_safe_sensitive_type_group" "test_sensitive_type_group" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = var.sensitive_type_group_display_name
  description = var.sensitive_type_group_description
}

resource "oci_data_safe_sensitive_type" "test_sensitive_type" {
  #Required
  compartment_id = var.compartment_ocid
  entity_type    = var.sensitive_type_entity_type

  #Optional
  display_name              = var.sensitive_type_display_name
}

resource "oci_data_safe_sensitive_type_group_grouped_sensitive_type" "test_sensitive_type_group_grouped_sensitive_type" {
  #Required
  sensitive_type_group_id = oci_data_safe_sensitive_type_group.test_sensitive_type_group.id

  #Optional
  patch_operations {
    #Required
    operation = "INSERT"
    selection = "items"
    value = {
      sensitiveTypeId = oci_data_safe_sensitive_type.test_sensitive_type.id
    }
  }
}

data "oci_data_safe_sensitive_type_group_grouped_sensitive_types" "test_sensitive_type_group_grouped_sensitive_types" {
  #Required
  sensitive_type_group_id = oci_data_safe_sensitive_type_group.test_sensitive_type_group.id

  #Optional
  sensitive_type_id = oci_data_safe_sensitive_type.test_sensitive_type.id
}

