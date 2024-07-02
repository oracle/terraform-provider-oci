// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_core_virtual_network" "exadbxs_vcn" {
  compartment_id = var.compartment_ocid
  cidr_blocks    = ["10.1.0.0/16"]
  display_name   = "exadbxs-tf-vcn"
  dns_label      = "tfvcn"
}

resource "oci_core_internet_gateway" "exadbxs_igw" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_virtual_network.exadbxs_vcn.id
  display_name   = "exadbxs-tf-igw"
}

resource "oci_core_route_table" "exadbxs_rt" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_virtual_network.exadbxs_vcn.id
  display_name   = "exadbxs-tf-route-table"
  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_internet_gateway.exadbxs_igw.id
  }
}

resource "oci_core_subnet" "exadbxs_client_subnet" {
  cidr_block        = "10.1.20.0/24"
  compartment_id  = var.compartment_ocid
  vcn_id            = oci_core_virtual_network.exadbxs_vcn.id
  route_table_id    = oci_core_route_table.exadbxs_rt.id
  security_list_ids = [
    oci_core_virtual_network.exadbxs_vcn.default_security_list_id,
    oci_core_security_list.exadbxs_security_list.id
  ]
  dns_label    = "tfclientsub"
  display_name = "exadbxs-tf-client-subnet"
}

resource "oci_core_subnet" "exadbxs_backup_subnet" {
  cidr_block     = "10.1.21.0/24"
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_virtual_network.exadbxs_vcn.id
  route_table_id = oci_core_route_table.exadbxs_rt.id
  dns_label      = "tfbackupsub"
  display_name   = "exadbxs-tf-backup-subnet"
}

resource "oci_core_security_list" "exadbxs_security_list" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_virtual_network.exadbxs_vcn.id
  display_name   = "exadbxs-security-list"

  ingress_security_rules {
    source   = "10.1.22.0/24"
    protocol = "6"
  }

  ingress_security_rules {
    source   = "10.1.22.0/24"
    protocol = "1"
  }

  ingress_security_rules {
    source   = "10.1.22.0/24"
    protocol = "all"
  }

  egress_security_rules {
    destination = "10.1.22.0/24"
    protocol    = "6"
  }

  egress_security_rules {
    destination = "10.1.22.0/24"
    protocol    = "1"
  }
}