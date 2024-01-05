// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {
}

variable "network_firewall_availability_domain" {
  default = ""
}

variable "network_firewall_defined_tags_value" {
  default = "value"
}

variable "network_firewall_display_name" {
  default = "displayName"
}

variable "network_firewall_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "network_firewall_id" {
  default = ""
}

variable "network_firewall_ipv4address" {
  default = ""
}

variable "network_firewall_ipv6address" {
  default = ""
}

variable "network_firewall_network_security_group_ids" {
  default = []
}

variable "network_firewall_state" {
  default = ""
}

variable "network_firewall_network_firewall_policy_id" {
  default = ""
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
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
  compartment_id             = var.compartment_id
}

resource "oci_network_firewall_network_firewall" "test_network_firewall" {
  #Required
  compartment_id             = var.compartment_id
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  subnet_id                  = oci_core_subnet.test_subnet.id

  #Optional
  availability_domain        = data.oci_identity_availability_domain.ad.name
  #defined_tags               = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.network_firewall_defined_tags_value)
  display_name               = var.network_firewall_display_name
  freeform_tags              = var.network_firewall_freeform_tags
  ipv4address                = var.network_firewall_ipv4address
  ipv6address                = var.network_firewall_ipv6address
  network_security_group_ids = var.network_firewall_network_security_group_ids
}

resource "oci_core_subnet" "test_subnet" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.0.2.0/24"
  display_name        = "TestSubnet"
  dns_label           = "testsubnet"
  security_list_ids   = [oci_core_vcn.vcn1.default_security_list_id]
  compartment_id      = var.compartment_id
  vcn_id              = oci_core_vcn.vcn1.id
  route_table_id      = oci_core_vcn.vcn1.default_route_table_id
  dhcp_options_id     = oci_core_vcn.vcn1.default_dhcp_options_id
}

resource "oci_core_vcn" "vcn1" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_id
  display_name   = "firewallVCN"
  dns_label      = "VcnFw"
}

data "oci_network_firewall_network_firewalls" "test_network_firewalls" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  availability_domain        = var.network_firewall_availability_domain
  display_name               = var.network_firewall_display_name
  id                         = var.network_firewall_id
  network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
  state                      = var.network_firewall_state
}