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

variable "unified_audit_policy_definition_access_level" {
  default = "RESTRICTED"
}

variable "unified_audit_policy_definition_compartment_id_in_subtree" {
  default = false
}

variable "unified_audit_policy_definition_defined_tags_value" {
  default = "value"
}

variable "unified_audit_policy_definition_description" {
  default = "Target database for HR and Payroll combined"
}

variable "unified_audit_policy_definition_display_name" {
  default = "AuditPolicy_HRandPayrollTarget"
}

variable "unified_audit_policy_definition_freeform_tags" {
  default = { "Department" = "Finance" }
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_unified_audit_policy_definition" "test_unified_audit_policy_definition" {
  #Required
  unified_audit_policy_definition_id = var.unified_audit_policy_definition_id

  #Optional
  description   = var.unified_audit_policy_definition_description
  display_name  = var.unified_audit_policy_definition_display_name
  freeform_tags = var.unified_audit_policy_definition_freeform_tags
}

data "oci_data_safe_unified_audit_policy_definition" "test_unified_audit_policy_definition" {
  #Optional
  unified_audit_policy_definition_id           = oci_data_safe_unified_audit_policy_definition.test_unified_audit_policy_definition.id
}

