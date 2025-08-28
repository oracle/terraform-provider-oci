// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "auditProfileId" {
  description = "OCID of the Data Safe Audit Profile"
  default = "<ocid>"
}

variable "workId" {
  description = "OCID of the Work Request from Data Safe Audit Profile creation"
  default = "<ocid>"
}


data "oci_data_safe_audit_profile_collected_audit_volumes" "test_audit_profile_collected_audit_volumes" {
  audit_profile_id = var.auditProfileId
  work_request_id  = var.workId

  # Optional filters
  # month_in_consideration_greater_than = "2025-01-01T00:00:00.000Z"
  # month_in_consideration_less_than    = "2025-12-01T00:00:00.000Z"
}


data "oci_data_safe_audit_profile_collected_audit_volume" "test_audit_profile_collected_audit_volume" {
  audit_profile_id = var.auditProfileId
  work_request_id  = var.workId

  # Optional filters
  # month_in_consideration_greater_than = "2025-01-01T00:00:00.000Z"
  # month_in_consideration_less_than    = "2025-12-01T00:00:00.000Z"
}



output "plural_collected_audit_volumes" {
  value = data.oci_data_safe_audit_profile_collected_audit_volumes.test_audit_profile_collected_audit_volumes.collected_audit_volume_collection
}

output "singular_collected_audit_volumes" {
  value = data.oci_data_safe_audit_profile_collected_audit_volume.test_audit_profile_collected_audit_volume.items
}