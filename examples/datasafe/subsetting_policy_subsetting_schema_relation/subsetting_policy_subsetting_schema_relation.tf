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

variable "subsetting_parent_object_name" {
  default = "EMPLOYEES"
}

variable "subsetting_child_object_name" {
  default = "PROJECTS"
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

resource "oci_data_safe_subsetting_policy_subsetting_schema_relation" "test_subsetting_policy_subsetting_schema_relation" {
  child_columns        = ["START_DATE"]
  child_object_name    = var.subsetting_child_object_name
  child_schema_name    = var.subsetting_schema_name
  parent_columns       = ["HIRE_DATE"]
  parent_object_name   = var.subsetting_parent_object_name
  parent_schema_name   = var.subsetting_schema_name
  subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id
}

data "oci_data_safe_subsetting_policy_subsetting_schema_relations" "test_subsetting_policy_subsetting_schema_relations" {
  subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id
  object               = [var.subsetting_child_object_name]
  relation_type        = "APP_DEFINED"
  schema_name          = [var.subsetting_schema_name]

  filter {
    name   = "key"
    values = [oci_data_safe_subsetting_policy_subsetting_schema_relation.test_subsetting_policy_subsetting_schema_relation.key]
  }
}

data "oci_data_safe_subsetting_policy_subsetting_schema_relation" "test_subsetting_policy_subsetting_schema_relation" {
  subsetting_policy_id           = oci_data_safe_subsetting_policy.test_subsetting_policy.id
  subsetting_schema_relation_key = oci_data_safe_subsetting_policy_subsetting_schema_relation.test_subsetting_policy_subsetting_schema_relation.key
}
