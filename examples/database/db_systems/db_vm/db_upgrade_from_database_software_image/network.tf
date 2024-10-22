# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      network.tf - Upstream dependencies
#
#    USAGE
#      Use the following path for Example Test & Backward Compatibility Test: database/db_systems/db_vm/db_upgrade_from_database_software_image
#
#    NOTES
#      Terraform Example: TestDatabaseDatabaseUpgradeResource_DbSoftwareImage
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   10/23/2024 - Created


resource "oci_core_vcn" "test_vcn" {
  cidr_block = "10.1.0.0/16"
  compartment_id = var.compartment_id
  display_name = "tfVcn"
  dns_label = "tfvcn"
}

resource "oci_core_route_table" "test_route_table" {
  compartment_id = var.compartment_id
  display_name = "tfRouteTable"
  route_rules {
    cidr_block = "0.0.0.0/0"
    description = "Internal traffic for OCI Services"
    network_entity_id = oci_core_internet_gateway.test_internet_gateway.id
  }
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_internet_gateway" "test_internet_gateway" {
  compartment_id = var.compartment_id
  display_name = "tfInternetGateway"
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block = "10.1.20.0/24"
  compartment_id = var.compartment_id
  dhcp_options_id = oci_core_vcn.test_vcn.default_dhcp_options_id
  display_name = "tfSubnet"
  dns_label = "tfsubnet"
  route_table_id = oci_core_route_table.test_route_table.id
  security_list_ids = [oci_core_vcn.test_vcn.default_security_list_id]
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_network_security_group" "test_nsg" {
  compartment_id = var.compartment_id
  display_name   = "tfNsg"
  vcn_id         = oci_core_vcn.test_vcn.id
}
