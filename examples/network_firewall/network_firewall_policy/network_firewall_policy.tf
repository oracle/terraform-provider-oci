// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
  This example shows how to create a firewall policy.
*/
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {

}

variable "network_firewall_policy_defined_tags_value" {
  default = "value"
}

variable "network_firewall_policy_display_name" {
  default = "network_firewall_policy"
}

variable "network_firewall_policy_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "network_firewall_policy_id" {
  default = "id"
}

variable "network_firewall_policy_state" {
  default = "ACTIVE"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}



resource "oci_network_firewall_network_firewall_policy" "test_network_firewall_policy" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  #defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.network_firewall_policy_defined_tags_value)
  display_name  = var.network_firewall_policy_display_name
  freeform_tags = var.network_firewall_policy_freeform_tags
}

data "oci_network_firewall_network_firewall_policies" "test_network_firewall_policies" {
  #Required
  compartment_id = var.compartment_id

  #Optional
#  display_name = var.network_firewall_policy_display_name
#  id           = var.network_firewall_policy_id
#  state        = var.network_firewall_policy_state
}

data "oci_network_firewall_network_firewall_policy" "test_network_firewall_policy" {
  #Required
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
}
