// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "security_policy_ocid" {}

variable "security_policy_access_level" {
  default = "RESTRICTED"
}

variable "security_policy_compartment_id_in_subtree" {
  default = false
}

variable "security_policy_defined_tags_value" {
  default = "value"
}

variable "security_policy_description" {
  default = "updated-description"
}

variable "security_policy_display_name" {
  default = "updated-name"
}

variable "security_policy_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "security_policy_status" {
  default = "ACTIVE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_security_policy" "test_security_policy" {
  #Required
  security_policy_id = var.security_policy_ocid

  #Optional
  description           = var.security_policy_description
  display_name          = var.security_policy_display_name
  freeform_tags         = var.security_policy_freeform_tags
}

data "oci_data_safe_security_policies" "test_security_policies" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  security_policy_id            = oci_data_safe_security_policy.test_security_policy.id
  access_level                  = var.security_policy_access_level
  compartment_id_in_subtree     = var.security_policy_compartment_id_in_subtree
}

