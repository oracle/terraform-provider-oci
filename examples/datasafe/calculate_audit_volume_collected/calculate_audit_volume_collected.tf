// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "calculate_audit_volume_collected_time_from_month" {
  default = "2024-05-17T15:05:28Z"
}

variable "calculate_audit_volume_collected_time_to_month" {
  default = "2024-05-17T17:05:28Z"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}


resource "oci_data_safe_calculate_audit_volume_collected" "test_calculate_audit_volume_collected" {
  #Required
  audit_profile_id = oci_data_safe_audit_profile_management.test_audit_profile_management.id
  time_from_month  = var.calculate_audit_volume_collected_time_from_month

  #Optional
  time_to_month = var.calculate_audit_volume_collected_time_to_month
}
