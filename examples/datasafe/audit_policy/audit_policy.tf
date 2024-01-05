// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "policy_ocid" {}

variable "audit_policy_access_level" {
  default = "RESTRICTED"
}

variable "audit_policy_compartment_id_in_subtree" {
  default = false
}

variable "audit_policy_defined_tags_value" {
  default = "value"
}

variable "audit_policy_description" {
  default = "Target database for HR and Payroll combined"
}

variable "audit_policy_display_name" {
  default = "AuditPolicy_HRandPayrollTarget"
}

variable "audit_policy_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "audit_policy_state" {
  default = "AVAILABLE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_audit_policy" "test_audit_policy" {
  #Required
  audit_policy_id = var.policy_ocid

  #Optional
  description   = var.audit_policy_description
  display_name  = var.audit_policy_display_name
  freeform_tags = var.audit_policy_freeform_tags
}

data "oci_data_safe_audit_policy" "test_audit_policy" {
  #Optional
  audit_policy_id           = oci_data_safe_audit_policy.test_audit_policy.id
}

