# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      network.tf - Network Infra file
#
#    USAGE
#      Use the following path for Example and Backward-Compatibility Tests: database/db_systems/db_vm/db_management
#    NOTES
#      Associated Integration Test: TestDatabaseCloudDatabaseManagementResource_basic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   06/23/2025 - Created


# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      network.tf - Network Infra file
#
#    USAGE
#      Use the following path for Example and Backward-Compatibility Tests: database/db_systems/db_vm/db_management
#    NOTES
#      Associated Integration Test: TestDatabaseCloudDatabaseManagementResource_basic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   06/23/2025 - Created


resource "oci_core_vcn" "test_vcn" {
  cidr_block = "10.1.0.0/16"
  compartment_id = var.compartment_id
  display_name = "tfVcn"
  dns_label = "tfvcn"
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block = "10.1.20.0/24"
  compartment_id = var.compartment_id
  dhcp_options_id = oci_core_vcn.test_vcn.default_dhcp_options_id
  display_name = "tfSubnet"
  dns_label = "tfsubnet"
  route_table_id = oci_core_route_table.test_route_table.id
  security_list_ids = [oci_core_security_list.test_security_list.id]
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_internet_gateway" "test_internet_gateway" {
  compartment_id = var.compartment_id
  display_name = "tfInternetGateway"
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_route_table" "test_route_table" {
  compartment_id = var.compartment_id
  display_name = "tfRouteTable"
  route_rules {
    description = "Internal traffic for OCI Services"
    destination = "0.0.0.0/0"
    destination_type = "CIDR_BLOCK"
    network_entity_id = oci_core_internet_gateway.test_internet_gateway.id
  }
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_security_list" "test_security_list" {
  compartment_id = var.compartment_id
  display_name = "tfSecurityList"
  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol = "6"
  }
  ingress_security_rules {
    protocol = "6"
    source = "0.0.0.0/0"
  }
  vcn_id = oci_core_vcn.test_vcn.id
}