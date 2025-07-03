# $Header$
#
# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      network.tf - network file
#
#    USAGE
#      Use the following path for Example Test & Backward-Compatibility-Test: database/pluggable_databases/remote_clone
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
    destination = "0.0.0.0/0"
    destination_type = "CIDR_BLOCK"
    network_entity_id = oci_core_internet_gateway.test_internet_gateway.id
  }
}

resource "oci_core_default_security_list" "test_default_security_list" {
  manage_default_resource_id = oci_core_virtual_network.test_vcn.default_security_list_id

  ingress_security_rules {
    source      = "0.0.0.0/0"
    source_type = "CIDR_BLOCK"
    protocol    = "all"
  }

  egress_security_rules {
    destination = "0.0.0.0/0"
    destination_type = "CIDR_BLOCK"
    protocol    = "all"
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

# resource "oci_core_network_security_group" "test_network_security_group" {
#   compartment_id  = var.compartment_id
#   vcn_id            = oci_core_virtual_network.test_vcn.id
#   display_name      =  "tfNSG"
# }

# REMOTE NETWORK
# resource "oci_core_virtual_network" "test_vcn_remote" {
#   compartment_id = var.compartment_id
#   cidr_block = "10.1.0.0/16"
#   display_name = "tfVCNRemote"
#   dns_label = "tfvcnremote"
# }
#
# resource "oci_core_route_table" "test_route_table_remote" {
#   compartment_id = var.compartment_id
#   vcn_id = oci_core_virtual_network.test_vcn_remote.id
#   route_rules {
#     destination = "0.0.0.0/0"
#     destination_type = "CIDR_BLOCK"
#     network_entity_id = oci_core_internet_gateway.test_internet_gateway_remote.id
#   }
# }
#
# resource "oci_core_default_security_list" "test_default_security_list_remote" {
#   manage_default_resource_id = oci_core_virtual_network.test_vcn_remote.default_security_list_id
#
#   ingress_security_rules {
#     source      = "0.0.0.0/0"
#     source_type = "CIDR_BLOCK"
#     protocol    = "all"
#   }
#
#   egress_security_rules {
#     destination = "0.0.0.0/0"
#     destination_type = "CIDR_BLOCK"
#     protocol    = "all"
#   }
# }
#
# resource "oci_core_internet_gateway" "test_internet_gateway_remote" {
#   compartment_id = var.compartment_id
#   vcn_id = oci_core_virtual_network.test_vcn_remote.id
#   display_name = "tfInternetGateway"
# }
#
# resource "oci_core_subnet" "test_subnet_remote" {
#   availability_domain = data.oci_identity_availability_domains.test_availability_domain.availability_domains.0.name
#   cidr_block          = "10.1.20.0/24"
#   display_name        = "tfSubnet"
#   compartment_id      = var.compartment_id
#   vcn_id              = oci_core_virtual_network.test_vcn_remote.id
#   route_table_id      = oci_core_route_table.test_route_table_remote.id
#   dhcp_options_id     = oci_core_virtual_network.test_vcn_remote.default_dhcp_options_id
#   security_list_ids   = [oci_core_virtual_network.test_vcn_remote.default_security_list_id]
#   dns_label           = "tfsubnet"
# }

#
# resource "oci_core_network_security_group" "test_network_security_group_remote" {
#   compartment_id  = var.compartment_id
#   vcn_id            = oci_core_virtual_network.test_vcn_remote.id
#   display_name      =  "tfNSG"
# }