// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "sql_firewall_policy_ocid" {}

variable "sql_firewall_policy_access_level" {
  default = "ACCESSIBLE"
}

variable "sql_firewall_policy_compartment_id_in_subtree" {
  default = true
}

variable "sql_firewall_policy_defined_tags_value" {
  default = "value"
}

variable "sql_firewall_policy_description" {
  default = "updated-description"
}

variable "sql_firewall_policydisplay_name" {
  default = "updated-name"
}

variable "sql_firewall_policy_freeform_tags" {
  default = { "Department" = "Finance" }
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_sql_firewall_policy" "test_sql_firewall_policy" {
  #Required
  sql_firewall_policy_id = var.sql_firewall_policy_ocid

  #Optional
  description           = var.sql_firewall_policy_description
  display_name          = var.sql_firewall_policydisplay_name
  freeform_tags         = var.sql_firewall_policy_freeform_tags
}

data "oci_data_safe_sql_firewall_policies" "test_sql_firewall_policies" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  sql_firewall_policy_id = var.sql_firewall_policy_ocid
}

