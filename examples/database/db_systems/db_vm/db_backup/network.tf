# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      network.tf - Resources file
#
#    USAGE
#      Example & Backward Compatibility Path: database/db_systems/db_vm/db_backup
#    NOTES
#      Terraform Integration Test: TestDatabaseBackupResource_basic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   11/1/2024 - Created


resource "oci_core_vcn" "test_vcn" {
  display_name = "tfVcnForDatabaseBackupExample"
  cidr_block = "10.0.0.0/16"
  compartment_id = var.compartment_id
  dns_label = "tfvcn"
}

resource "oci_core_route_table" "test_route_table" {
  display_name = "tfRouteTable"
  compartment_id = var.compartment_id
  route_rules {
    cidr_block = "0.0.0.0/0"
    description = "Internal traffic for OCI Services"
    network_entity_id = oci_core_internet_gateway.test_internet_gateway.id
  }
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_internet_gateway" "test_internet_gateway" {
  display_name = "tfInternetGateway"
  compartment_id = var.compartment_id
  defined_tags = map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")
  enabled = "true"
  freeform_tags = {
    "Department" = "Finance"
  }
  lifecycle {
    ignore_changes = [defined_tags]
  }
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_subnet" "test_subnet" {
  display_name = "tfPublicSubnet"
  cidr_block = "10.0.0.0/24"
  compartment_id = var.compartment_id
  dhcp_options_id = oci_core_vcn.test_vcn.default_dhcp_options_id
  dns_label = "tfpublicsubnet"
  route_table_id = oci_core_route_table.test_route_table.id
  security_list_ids = [oci_core_vcn.test_vcn.default_security_list_id]
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_security_list" "test_private_subnet_security_list" {
  display_name = "tfRecoveryServiceSecurityList"
  compartment_id = var.compartment_id
  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol = "all"
  }
  ingress_security_rules {
    protocol = "6"
    source = "10.0.0.0/16"
    source_type = "CIDR_BLOCK"
    tcp_options {
      min = "8005"
      max = "8005"
    }
  }
  ingress_security_rules {
    protocol = "6"
    source = "10.0.0.0/16"
    source_type = "CIDR_BLOCK"
    tcp_options {
      min = "2484"
      max = "2484"
    }
  }
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_service_gateway" "test_service_gateway" {
  display_name = "tfRecoveryServiceServiceGateway"
  compartment_id = var.compartment_id
  services {
    service_id = data.oci_core_services.test_services.services.0.id
  }
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_route_table" "test_private_subnet_route_table" {
  display_name = "tfRecoveryServicePrivateSubnetRouteTable"
  compartment_id = var.compartment_id
  route_rules {
    description = "Recovery Service traffic for OCI Services"
    destination = data.oci_core_services.test_services.services[0].cidr_block
    destination_type = "SERVICE_CIDR_BLOCK"
    network_entity_id = oci_core_service_gateway.test_service_gateway.id
  }
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_subnet" "test_private_subnet" {
  display_name = "tfPrivateSubnet"
  cidr_block = "10.0.1.0/24"
  compartment_id = var.compartment_id
  dhcp_options_id = oci_core_vcn.test_vcn.default_dhcp_options_id
  dns_label = "tfprivatesubnet"
  prohibit_public_ip_on_vnic = "true"
  route_table_id = oci_core_route_table.test_private_subnet_route_table.id
  security_list_ids = [oci_core_security_list.test_private_subnet_security_list.id]
  vcn_id = oci_core_vcn.test_vcn.id
}