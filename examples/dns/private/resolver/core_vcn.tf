// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This file demonstrates core vcn creation
 */

resource "oci_core_vcn" "test_vcn" {
  cidr_block = "10.0.0.0/16"
  compartment_id = "${var.compartment_ocid}"
  display_name = "test_vcn"
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block = "10.0.0.0/24"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${oci_core_vcn.test_vcn.id}"
  display_name = "test_vcn"
}

data "oci_core_vcn_dns_resolver_association" "test_vcn_dns_resolver_association" {
  vcn_id = "${oci_core_vcn.test_vcn.id}"
}