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

variable "subsetting_rule_display_name" {
  default = "example-subsetting-rule"
}

variable "subsetting_rule_description" {
  default = "Example percent subsetting rule"
}

variable "subsetting_rule_percent" {
  default = 10
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
    schema_source = "TARGET"
    target_id     = var.data_safe_target_ocid
  }
}

resource "oci_data_safe_subsetting_policy_subsetting_rule" "test_subsetting_policy_subsetting_rule" {
  subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id
  display_name         = var.subsetting_rule_display_name
  description          = var.subsetting_rule_description

  scope {
    schema_name = var.subsetting_schema_name
    scope_type  = "SPECIFIC"
    object      = var.subsetting_object_name
  }

  subset_rule_entry {
    rule_type = "PERCENT"
    percent   = var.subsetting_rule_percent
  }
}

resource "oci_data_safe_subsetting_policy_subsetting_rule" "test_partition_subsetting_rule" {
  subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id

  scope {
    schema_name = "HR_TEST"
    scope_type  = "SPECIFIC"
    object      = "PARTITION_TEST"
  }

  subset_rule_entry {
    rule_type           = "PARTITION"
    sub_partitions_list = ["P2025.P2025_HR", "P2026.P2026_HR", "P2026.P2026_IT"]
  }

  depends_on = [oci_data_safe_subsetting_policy_subsetting_rule.test_subsetting_policy_subsetting_rule]
}

data "oci_data_safe_subsetting_policy_subsetting_rules" "test_subsetting_policy_subsetting_rules" {
  subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id
  object               = [var.subsetting_object_name]
  schema_name          = [var.subsetting_schema_name]
}

data "oci_data_safe_subsetting_policy_subsetting_rule" "test_subsetting_policy_subsetting_rule" {
  subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id
  subsetting_rule_key  = oci_data_safe_subsetting_policy_subsetting_rule.test_subsetting_policy_subsetting_rule.key
}
