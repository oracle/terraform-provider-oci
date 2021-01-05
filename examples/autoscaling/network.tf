// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_core_vcn" "ExampleVCN" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "TFExampleVCN"
  dns_label      = "tfexamplevcn"
}

resource "oci_core_subnet" "ExampleSubnet" {
  availability_domain = data.oci_identity_availability_domain.AD.name
  cidr_block          = "10.1.20.0/24"
  display_name        = "TFExampleSubnet"
  dns_label           = "tfexamplesubnet"
  security_list_ids   = [oci_core_vcn.ExampleVCN.default_security_list_id]
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.ExampleVCN.id
  route_table_id      = oci_core_route_table.ExampleRT.id
  dhcp_options_id     = oci_core_vcn.ExampleVCN.default_dhcp_options_id
}

resource "oci_core_internet_gateway" "ExampleIG" {
  compartment_id = var.compartment_ocid
  display_name   = "TFExampleIG"
  vcn_id         = oci_core_vcn.ExampleVCN.id
}

resource "oci_core_route_table" "ExampleRT" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.ExampleVCN.id
  display_name   = "TFExampleRouteTable"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_internet_gateway.ExampleIG.id
  }
}

