// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_id
  display_name   = "exampleVCN"
  dns_label      = "tfexamplevcn"
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block        = "10.0.0.0/24"
  display_name      = "tfexampleSubnet"
  dns_label         = "tfexampleSubnet"
  compartment_id    = var.compartment_id
  vcn_id            = oci_core_vcn.test_vcn.id
  security_list_ids = [oci_core_vcn.test_vcn.default_security_list_id]
  route_table_id    = oci_core_vcn.test_vcn.default_route_table_id
  dhcp_options_id   = oci_core_vcn.test_vcn.default_dhcp_options_id
}
