// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "calculate_audit_volume_available_audit_collection_start_time" {
  default = "2023-05-17T15:05:28Z"
}

variable "calculate_audit_volume_available_database_unique_name" {
  default = "databaseUniqueName"
}

variable "calculate_audit_volume_available_trail_locations" {
  default = []
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_calculate_audit_volume_available" "test_calculate_audit_volume_available" {
  #Required
  audit_profile_id = oci_data_safe_audit_profile_management.test_audit_profile_management.id

  #Optional
  audit_collection_start_time = var.calculate_audit_volume_available_audit_collection_start_time
  database_unique_name        = var.calculate_audit_volume_available_database_unique_name
  trail_locations             = var.calculate_audit_volume_available_trail_locations
}

