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

resource "oci_data_safe_subsetting_policy_subsetting_rule" "test_subsetting_rule" {
  subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id

  scope {
    schema_name = "HR_TEST"
    scope_type  = "SPECIFIC"
    object      = "EMPLOYEES"
  }

  subset_rule_entry {
    rule_type = "PERCENT"
    percent   = 10
  }
}

# Retrieves the generated health report and supports its lifecycle management.
resource "oci_data_safe_subsetting_policy_health_report_management" "test_subsetting_policy_health_report_management" {
  target_id            = var.data_safe_target_ocid
  subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id

  target_credentials {
    user_name = var.target_database_user_name
    password  = var.target_database_password
  }

  depends_on = [oci_data_safe_subsetting_policy_subsetting_rule.test_subsetting_rule]
}

data "oci_data_safe_subsetting_policy_health_report" "test_subsetting_policy_health_report" {
  subsetting_policy_health_report_id = oci_data_safe_subsetting_policy_health_report_management.test_subsetting_policy_health_report_management.id
}

data "oci_data_safe_subsetting_policy_health_reports" "test_subsetting_policy_health_reports" {
  compartment_id       = var.compartment_ocid
  subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id
 depends_on = [oci_data_safe_subsetting_policy_health_report_management.test_subsetting_policy_health_report_management]
}

data "oci_data_safe_subsetting_policy_health_report_logs" "test_subsetting_policy_health_report_logs" {
  subsetting_policy_health_report_id = oci_data_safe_subsetting_policy_health_report_management.test_subsetting_policy_health_report_management.id
}
