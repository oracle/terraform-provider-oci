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
    schema_source = "TARGET"
    target_id     = var.data_safe_target_ocid
  }
}

resource "oci_data_safe_subsetting_policy_subsetting_rule" "test_subsetting_policy_subsetting_rule" {
  subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id

  scope {
    schema_name = var.subsetting_schema_name
    scope_type  = "SPECIFIC"
    object      = var.subsetting_object_name
  }

  subset_rule_entry {
    rule_type = "PERCENT"
    percent   = 10
  }

  peer_tables_action         = "MINIMUM_ROWS"
  related_tables_propagation = "DESCENDANTS"
}

data "oci_data_safe_subsetting_policy_subsetting_rule_processing_chain_objects" "test_subsetting_policy_subsetting_rule_processing_chain_objects" {
  subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id
  subsetting_rule_key  = oci_data_safe_subsetting_policy_subsetting_rule.test_subsetting_policy_subsetting_rule.key
  filter {
    name   = "propagation_impact"
    values = ["PARENT_CHILD_SUBSET"]
  }
}

# The processing-chain object key is obtained from the policy after the rule is created.
resource "oci_data_safe_subsetting_policy_subsetting_rule_processing_chain_object" "test_subsetting_policy_subsetting_rule_processing_chain_object" {
  processing_chain_object_key = data.oci_data_safe_subsetting_policy_subsetting_rule_processing_chain_objects.test_subsetting_policy_subsetting_rule_processing_chain_objects.subsetting_rule_processing_chain_objects_collection[0].items[0].key
  subsetting_policy_id        = oci_data_safe_subsetting_policy.test_subsetting_policy.id
  subsetting_rule_key         = oci_data_safe_subsetting_policy_subsetting_rule.test_subsetting_policy_subsetting_rule.key
  is_enabled_for_processing   = false
}
