// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# This example demonstrates the various dhcp option configurations.

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

resource "oci_core_vcn" "example_vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "exampleVCN"
  dns_label      = "tfexamplevcn"
}

resource "oci_core_dhcp_options" "example_dhcp_options_1" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.example_vcn.id
  display_name   = "exampleDhcpOptions1"

  // required
  options {
    type        = "DomainNameServer"
    server_type = "VcnLocalPlusInternet"
  }

  // optional
  options {
    type                = "SearchDomain"
    search_domain_names = ["test.com"]
  }
}

resource "oci_core_dhcp_options" "example_dhcp_options_2" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.example_vcn.id
  display_name   = "exampleDhcpOptions2"

  // required
  options {
    type               = "DomainNameServer"
    server_type        = "CustomDnsServer"
    custom_dns_servers = ["8.8.4.4", "8.8.8.8"]
  }

  // optional
  options {
    type                = "SearchDomain"
    search_domain_names = ["test.com"]
  }
}

