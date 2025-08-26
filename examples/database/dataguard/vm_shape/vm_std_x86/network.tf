# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      variables.tf - Shepherd Data Source file
#
#    USAGE
#
#    NOTES
#      Terraform Example: TestDatabaseDataGuardAssociationResourceVmStdx86_basic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    aavadhan   08/18/2025 - Created

resource "oci_core_vcn" "test_vcn" {
  cidr_block = "10.0.0.0/16"
  compartment_id = var.compartment_id
  display_name = "displayName"
  dns_label = "dnslabel"
}

resource "oci_core_network_security_group" "test_network_security_group" {
  compartment_id = var.compartment_id
  display_name = "displayName2"
  freeform_tags = {
    "Department" = "Accounting"
  }

  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_service_gateway" "test_service_gateway" {
  compartment_id = var.compartment_id
  display_name = "test_service_gateway"
  services {
    service_id = data.oci_core_services.test_services.services.0.id
  }
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_default_route_table" "test_vcn_default_route_table" {
  manage_default_resource_id = oci_core_vcn.test_vcn.default_route_table_id
  route_rules {
    description = "Internal traffic for OCI Services"
    destination = data.oci_core_services.test_services.services[0].cidr_block
    destination_type = "SERVICE_CIDR_BLOCK"
    network_entity_id = oci_core_service_gateway.test_service_gateway.id
  }
}

resource "oci_core_route_table" "test_route_table" {
  compartment_id = var.compartment_id
  display_name = "test_subnet_rt"
  route_rules {
    description = "Internal traffic for OCI Services"
    destination = data.oci_core_services.test_services.services[0].cidr_block
    destination_type = "SERVICE_CIDR_BLOCK"
    network_entity_id = oci_core_service_gateway.test_service_gateway.id
  }
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_security_list" "test_security_list" {
  compartment_id = var.compartment_id
  display_name = "test_security_list"
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

resource "oci_core_subnet" "test_subnet" {
  cidr_block = "10.0.2.0/24"
  compartment_id = var.compartment_id
  dhcp_options_id = oci_core_vcn.test_vcn.default_dhcp_options_id
  display_name = "test_subnet"
  dns_label = "tftestsubnet"
  prohibit_public_ip_on_vnic = "true"
  route_table_id = oci_core_route_table.test_route_table.id
  security_list_ids = [oci_core_security_list.test_security_list.id]
  vcn_id = oci_core_vcn.test_vcn.id
}