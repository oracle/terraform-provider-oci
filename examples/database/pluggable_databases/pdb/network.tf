# $Header$
#
# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      network.tf - network / infra related file
#
#    USAGE
#      Use the following path for Example Test & Backward-Compatibility-Test: database/pluggable_databases/pdb
#    NOTES
#      Terraform Example:
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   05/08/2025 - Created


resource "oci_core_virtual_network" "test_vcn" {
  compartment_id = var.compartment_id
  cidr_block = "10.1.0.0/16"
  display_name = "tfVCN"
  dns_label = "tfvcn"
}

resource "oci_core_route_table" "test_route_table" {
  compartment_id = var.compartment_id
  vcn_id = oci_core_virtual_network.test_vcn.id
  route_rules {
    cidr_block = "0.0.0.0/0"
    network_entity_id = oci_core_internet_gateway.test_internet_gateway.id
  }
}
resource "oci_core_internet_gateway" "test_internet_gateway" {
  compartment_id = var.compartment_id
  vcn_id = oci_core_virtual_network.test_vcn.id
  display_name = "tfInternetGateway"
}

resource "oci_core_subnet" "test_subnet" {
  availability_domain = data.oci_identity_availability_domains.test_availability_domain.availability_domains.0.name
  cidr_block          = "10.1.20.0/24"
  display_name        = "tfSubnet"
  compartment_id      = var.compartment_id
  vcn_id              = oci_core_virtual_network.test_vcn.id
  route_table_id      = oci_core_route_table.test_route_table.id
  dhcp_options_id     = oci_core_virtual_network.test_vcn.default_dhcp_options_id
  security_list_ids   = [oci_core_virtual_network.test_vcn.default_security_list_id]
  dns_label           = "tfsubnet"
}

resource "oci_core_network_security_group" "test_network_security_group" {
  compartment_id  = var.compartment_id
  vcn_id            = oci_core_virtual_network.test_vcn.id
  display_name      =  "tfNSG"
}
