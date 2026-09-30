// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "data_safe_subsetting_report_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_data_safe_subsetting_report" "test_subsetting_report" {
  subsetting_report_id = var.data_safe_subsetting_report_ocid
}

data "oci_data_safe_subsetting_reports" "test_subsetting_reports" {
  compartment_id = var.compartment_ocid
}

data "oci_data_safe_subsetting_report_subsetted_objects" "test_subsetting_report_subsetted_objects" {
  subsetting_report_id = var.data_safe_subsetting_report_ocid
}

data "oci_data_safe_subsetting_report_subsetting_errors" "test_subsetting_report_subsetting_errors" {
  subsetting_report_id = var.data_safe_subsetting_report_ocid
}
