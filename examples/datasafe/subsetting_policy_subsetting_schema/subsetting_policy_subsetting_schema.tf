// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "data_safe_target_ocid" {}

variable "subsetting_schema_name" {
  default = "HR_TEST"
}

variable "subsetting_object_name" {
  default = "EMPLOYEES"
}

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
    schema_source          = "TARGET"
    schemas_for_subsetting = [var.subsetting_schema_name]
    target_id              = var.data_safe_target_ocid
  }
}

data "oci_data_safe_subsetting_policy_subsetting_schemas" "test_subsetting_policy_subsetting_schemas" {
  subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id
  schema_name          = [var.subsetting_schema_name]
}

data "oci_data_safe_subsetting_policy_subsetting_schema_objects" "test_subsetting_policy_subsetting_schema_objects" {
  subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id
  object               = [var.subsetting_object_name]
  schema_name          = [var.subsetting_schema_name]
}
