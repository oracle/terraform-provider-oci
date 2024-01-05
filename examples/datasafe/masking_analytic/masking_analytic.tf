// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "data_safe_target_ocid" {}

variable "masking_analytic_compartment_id_in_subtree" {
  default = true
}

variable "masking_analytic_group_by" {
  default = "targetId"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "masking_policy_column_source_column_source" {
  default = "TARGET"
}

variable "masking_policy_display_name" {
  default = "displayName"
}

resource "oci_data_safe_masking_policy" "test_masking_policy" {
  #Required
  column_source {
    #Required
    column_source = var.masking_policy_column_source_column_source

    #Optional
    target_id               = var.data_safe_target_ocid
  }
  compartment_id = var.compartment_ocid

  #Optional
  display_name                = var.masking_policy_display_name
}

data "oci_data_safe_masking_analytics" "test_masking_analytics" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  compartment_id_in_subtree = var.masking_analytic_compartment_id_in_subtree
  group_by                  = var.masking_analytic_group_by
  masking_policy_id         = oci_data_safe_masking_policy.test_masking_policy.id
}

