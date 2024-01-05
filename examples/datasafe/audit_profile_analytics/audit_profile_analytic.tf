// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "audit_profile_analytic_access_level" {
  default = "RESTRICTED"
}

variable "audit_profile_analytic_compartment_id_in_subtree" {
  default = false
}

variable "audit_profile_analytic_group_by" {
  default = []
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_data_safe_audit_profile_analytic" "test_audit_profile_analytic" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  group_by                  = var.audit_profile_analytic_group_by
}

