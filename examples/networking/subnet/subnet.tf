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

variable "compartment_ocid" {
}

variable "region" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

resource "oci_core_vcn" "vcn1" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "exampleVCN"
  dns_label      = "tfexamplevcn"
}

// A regional subnet will not specify an Availability Domain
resource "oci_core_subnet" "regional_subnet" {
  cidr_block        = "10.0.1.0/24"
  display_name      = "regionalSubnet"
  dns_label         = "regionalsubnet"
  compartment_id    = var.compartment_ocid
  vcn_id            = oci_core_vcn.vcn1.id
  security_list_ids = [oci_core_vcn.vcn1.default_security_list_id]
  route_table_id    = oci_core_vcn.vcn1.default_route_table_id
  dhcp_options_id   = oci_core_vcn.vcn1.default_dhcp_options_id
}

// An AD based subnet will supply an Availability Domain
resource "oci_core_subnet" "ad_subnet" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.0.2.0/24"
  display_name        = "ADSubnet"
  dns_label           = "adsubnet"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.vcn1.id
  security_list_ids   = [oci_core_vcn.vcn1.default_security_list_id]
  route_table_id      = oci_core_vcn.vcn1.default_route_table_id
  dhcp_options_id     = oci_core_vcn.vcn1.default_dhcp_options_id
}

