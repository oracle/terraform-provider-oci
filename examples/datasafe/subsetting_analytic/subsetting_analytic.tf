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

resource "oci_data_safe_subsetting_policy" "test_subsetting_policy" {
  compartment_id = var.compartment_ocid

  schema_source {
    schema_source = "TARGET"
    target_id     = var.data_safe_target_ocid
  }
}

data "oci_data_safe_subsetting_analytics" "test_subsetting_analytics" {
  compartment_id       = var.compartment_ocid
  subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id
  target_id            = var.data_safe_target_ocid
}
