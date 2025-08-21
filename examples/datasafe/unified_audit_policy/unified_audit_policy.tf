// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "security_policy_id" {}
variable "unified_audit_policy_definition_id" {}

variable "unified_audit_policy_access_level" {
  default = "RESTRICTED"
}

variable "unified_audit_policy_compartment_id_in_subtree" {
  default = false
}

variable "unified_audit_policy_defined_tags_value" {
  default = "value"
}

variable "unified_audit_policy_description" {
  default = "Target database for HR and Payroll combined"
}

variable "unified_audit_policy_display_name" {
  default = "AuditPolicy_HRandPayrollTarget"
}

variable "unified_audit_policy_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "unified_audit_policy_state" {
  default = "ACTIVE"
}

variable "unified_audit_policy_status" {
  default = "ENABLED"
}

variable "unified_audit_policy_entity_selection" {
  default = "INCLUDE"
}

variable "unified_audit_policy_entity_type" {
  default = "USER"
}

variable "unified_audit_policy_operation_status" {
  default = "SUCCESS"
}

variable "unified_audit_policy_user_names" {
  default = ["SAMPLE"]
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_unified_audit_policy" "test_unified_audit_policy" {
  #Required
  compartment_id = var.compartment_id
  security_policy_id = var.security_policy_id
  unified_audit_policy_definition_id = var.unified_audit_policy_definition_id
  status = var.unified_audit_policy_status
  conditions {
    entity_selection = var.unified_audit_policy_entity_selection
    entity_type = var.unified_audit_policy_entity_type
    operation_status = var.unified_audit_policy_operation_status
    user_names = var.unified_audit_policy_user_names
  }

  #Optional
  description   = var.unified_audit_policy_description
  display_name  = var.unified_audit_policy_display_name
  freeform_tags = var.unified_audit_policy_freeform_tags
}

data "oci_data_safe_unified_audit_policy" "test_unified_audit_policy" {
  #Optional
  unified_audit_policy_id           = oci_data_safe_unified_audit_policy.test_unified_audit_policy.id
}

