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

variable "database_release" {
  default = "19.0.0.0.0"
}

data "oci_dblm_patch_management" "test_patch_management"{
  #Required
  compartment_id = var.compartment_ocid
  database_release = "19.0.0.0.0"
  time_started_greater_than_or_equal_to = "2006-01-02T15:04:05Z"
  time_started_less_than = "2026-01-02T15:04:05Z"
}

data "oci_dblm_patch_management_databases" "test_dblm_patch_management_databases"{
  #Optional
  compartment_id = var.compartment_ocid
  database_release = var.database_release
  database_type = "SI"
}