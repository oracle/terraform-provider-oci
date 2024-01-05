// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "report_ocid" {}

variable "report_access_level" {
  default = "RESTRICTED"
}

variable "report_compartment_id_in_subtree" {
  default = false
}

variable "report_display_name" {
  default = "displayName"
}

variable "report_state" {
  default = "ACTIVE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_data_safe_reports" "test_reports" {
  #Required
  compartment_id = var.compartment_ocid
}

