// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "data_safe_target_ocid" {}
variable "masking_health_report_id" {}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "access_level" {
  default = "RESTRICTED"
}

variable "report_compartment_id_in_subtree" {
  default = false
}
data "oci_data_safe_masking_policy_health_reports" "test_health_reports" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  access_level = var.access_level
  compartment_id_in_subtree =  var.report_compartment_id_in_subtree
}
data "oci_data_safe_masking_policy_health_report" "test_health_report" {
  #Required
  masking_policy_health_report_id = var.masking_health_report_id
}
