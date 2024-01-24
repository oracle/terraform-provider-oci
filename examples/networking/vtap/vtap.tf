// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

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

resource "oci_core_capture_filter" "example_capture_filter" {
  compartment_id = var.compartment_ocid
  display_name   = "exampleCaptureFilter"
  filter_type    = "VTAP"
  vtap_capture_filter_rules {
    traffic_direction = "INGRESS"
  }
}

// A regional subnet will not specify an Availability Domain
resource "oci_core_subnet" "regional_subnet" {
  cidr_block        = "10.1.1.0/24"
  display_name      = "regionalSubnet"
  dns_label         = "regionalsubnet"
  compartment_id    = var.compartment_ocid
  vcn_id            = oci_core_vcn.example_vcn.id
  security_list_ids = [oci_core_vcn.example_vcn.default_security_list_id]
  route_table_id    = oci_core_vcn.example_vcn.default_route_table_id
  dhcp_options_id   = oci_core_vcn.example_vcn.default_dhcp_options_id
}

resource "oci_core_network_security_group" "test_network_security_group" {
  #Required
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.example_vcn.id
}

resource "oci_load_balancer" "lb1" {
  shape          = "100Mbps"
  compartment_id = var.compartment_ocid

  subnet_ids = [
    oci_core_subnet.regional_subnet.id,
  ]

  display_name               = "lb1"
  is_private                 = true
  network_security_group_ids = [oci_core_network_security_group.test_network_security_group.id]
}

resource "oci_core_vtap" "example_vtap" {
  compartment_id    = var.compartment_ocid
  vcn_id            = oci_core_vcn.example_vcn.id
  display_name      = "exampleVtap"
  capture_filter_id = oci_core_capture_filter.example_capture_filter.id
  source_id         = oci_load_balancer.lb1.id
  source_type       = "LOAD_BALANCER"
  target_ip         = "1.1.1.1"
  is_vtap_enabled   = true
}

