// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

locals {
  policy_b_nat_rules = [
    for idx in range(16) : {
      drg_nat_rule_priority = idx + 1
      original_source       = "10.0.${floor(idx/256)}.${idx%256}/32"
      translated_source     = "192.168.${floor(idx/256)}.${idx%256}/32"
    }
  ]
}

provider "oci" {
  tenancy_ocid = var.tenancy_ocid
  user_ocid = var.user_ocid
  fingerprint = var.fingerprint
  private_key_path = var.private_key_path
  region = var.region
}


resource "oci_core_vcn" "test_vcn" {
  // Required
  cidr_block = "10.0.0.0/16"
  compartment_id = var.compartment_ocid

  // Optional
  display_name = "MyTestVcnA"
  dns_label = "dnslabelA"
}


resource "oci_core_drg" "test_vcn_drg" {
  // Required
  compartment_id = var.compartment_ocid

  // Optional
  display_name = "MyTestVcnDrg"
}

resource "oci_core_drg_nat_policy" "test_drg_nat_policy_a" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name  = "TestPolicy"
  freeform_tags = { "Department" = "Finance" }
}

resource "oci_core_drg_nat_policy" "test_drg_nat_policy_b" {
  #Required
  compartment_id = var.compartment_ocid
}

resource "oci_core_drg_attachment" "test_vcn_drg_attachment_a" {
  // Required
  drg_id = oci_core_drg.test_vcn_drg.id
  vcn_id = oci_core_vcn.test_vcn.id

  // Optional
  drg_nat_policy_id  = oci_core_drg_nat_policy.test_drg_nat_policy_a.id

}

resource "oci_core_drg_nat_policy_drg_nat_rule" "test_drg_nat_rule_a" {
  drg_nat_policy_id      = oci_core_drg_nat_policy.test_drg_nat_policy_a.id
  drg_nat_rule_priority  = 1

  original_source        = "10.0.0.0/24"
  translated_source      = "192.168.0.0/24"
  original_destination   = "172.16.1.0/24"
  translated_destination = "10.1.1.0/24"
}

resource "oci_core_drg_nat_policy_drg_nat_rule" "test_drg_nat_rule_b" {
  drg_nat_policy_id      = oci_core_drg_nat_policy.test_drg_nat_policy_a.id
  drg_nat_rule_priority  = 2

  original_destination        = "192.168.1.0/24"
  translated_destination      = "10.0.0.0/24"
}

resource "oci_core_drg_nat_policy_drg_nat_rule" "policy_b_nat_rules" {
  for_each = { for idx, rule in local.policy_b_nat_rules : idx => rule }

  drg_nat_policy_id      = oci_core_drg_nat_policy.test_drg_nat_policy_b.id
  drg_nat_rule_priority  = each.value.drg_nat_rule_priority
  original_source        = lookup(each.value, "original_source", null)
  translated_source      = lookup(each.value, "translated_source", null)
  original_destination   = lookup(each.value, "original_destination", null)
  translated_destination = lookup(each.value, "translated_destination", null)
}

data "oci_core_drg_nat_policy" "test_drg_nat_policy" {
  drg_nat_policy_id      = oci_core_drg_nat_policy.test_drg_nat_policy_a.id
}

data "oci_core_drg_nat_policies" "test_drg_nat_policies" {
  #Required
  compartment_id = var.compartment_ocid
}

data "oci_core_drg_nat_policy_drg_nat_rules" "test_drg_nat_policy_drg_nat_rules" {
  #Required
  drg_nat_policy_id = oci_core_drg_nat_policy.test_drg_nat_policy_b.id
}