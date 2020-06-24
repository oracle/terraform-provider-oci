// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_core_vcn" "vcn_rd" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "vcnRD"
  dns_label      = "vcnrd"
}

resource "oci_core_internet_gateway" "internetgateway_rd" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "internetgatewayRD"
  vcn_id         = "${oci_core_vcn.vcn_rd.id}"
}

resource "oci_core_route_table" "routetable_rd" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.vcn_rd.id}"
  display_name   = "routetableRD"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_internet_gateway.internetgateway_rd.id}"
  }
}

resource "oci_core_network_security_group" "test_network_security_group" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.vcn_rd.id}"
  display_name   = "displayName"
}

resource "oci_core_network_security_group" "test_network_security_group_backup" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.vcn_rd.id}"
  display_name   = "displayName"
}
