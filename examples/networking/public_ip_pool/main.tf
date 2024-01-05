// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}
variable "ssh_public_key" {}

variable "byoip_range_id" {}

variable "byoipv6_range_id" {}

variable "public_ip_pool_cidr_block" {}

resource "oci_core_public_ip" "test_public_ip" {
  compartment_id    = "${var.compartment_ocid}"
  lifetime          = "RESERVED"
  public_ip_pool_id = "${oci_core_public_ip_pool_capacity.test_public_ip_pool_capacity.public_ip_pool_id}"
}

resource "oci_core_public_ip_pool_capacity" "test_public_ip_pool_capacity" {
  public_ip_pool_id = "${oci_core_public_ip_pool.test_public_ip_pool.id}"
  cidr_block        = "${var.public_ip_pool_cidr_block}"
  byoip_id          = "${var.byoip_range_id}"
  byoipv6_id        = "${var.byoipv6_range_id}"
}

resource "oci_core_public_ip_pool" "test_public_ip_pool" {
  compartment_id = "${var.compartment_ocid}"
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = "${var.compartment_ocid}"
}

resource "oci_core_nat_gateway" "test_nat_gateway" {
  block_traffic  = "false"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "displayName"
  public_ip_id   = "${oci_core_public_ip.test_public_ip.id}"
  vcn_id         = "${oci_core_vcn.test_vcn.id}"
}

data "oci_core_public_ip_pool" "public_ip_pool" {
  public_ip_pool_id = "${oci_core_public_ip_pool_capacity.test_public_ip_pool_capacity.public_ip_pool_id}"
}

output "public_ip_pool" {
  value = [
    "${data.oci_core_public_ip_pool.public_ip_pool}",
  ]
}
