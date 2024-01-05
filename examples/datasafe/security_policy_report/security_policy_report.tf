// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "security_policy_report_ocid" {}
variable "data_safe_target_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_data_safe_security_policy_reports" "test_security_policy_reports" {
  #Required
  compartment_id = var.compartment_ocid
  access_level = "ACCESSIBLE"
  compartment_id_in_subtree = true
}

data "oci_data_safe_security_policy_report" "test_security_policy_report" {
  #Required
  security_policy_report_id = var.security_policy_report_ocid  
}
