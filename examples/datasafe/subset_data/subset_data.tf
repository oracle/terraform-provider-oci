// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "data_safe_target_ocid" {}
variable "target_database_user_name" {}

variable "target_database_password" {
  sensitive = true
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

  scope {
    schema_name = "HR_TEST"
    scope_type  = "SPECIFIC"
    object      = "EMPLOYEES"
  }

  subset_rule_entry {
    rule_type = "CONDITION"
    condition = "1 = 1"
  }
}

# Starts the asynchronous subsetting operation. This resource is intentionally non-deletable.
resource "oci_data_safe_subset_data" "test_subset_data" {
  subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id
  target_id            = var.data_safe_target_ocid

  target_credentials {
    user_name = var.target_database_user_name
    password  = var.target_database_password
  }

  depends_on = [oci_data_safe_subsetting_policy_subsetting_rule.test_subsetting_policy_subsetting_rule]
}
