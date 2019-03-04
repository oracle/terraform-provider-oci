// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

resource "oci_core_virtual_network" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "tfVcnForClusters"
}

resource "oci_core_internet_gateway" "test_ig" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "tfClusterInternetGateway"
  vcn_id         = "${oci_core_virtual_network.test_vcn.id}"
}

resource "oci_core_route_table" "test_route_table" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_virtual_network.test_vcn.id}"
  display_name   = "tfClustersRouteTable"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_internet_gateway.test_ig.id}"
  }
}

resource "oci_core_subnet" "clusterSubnet_1" {
  #Required
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  cidr_block          = "10.0.20.0/24"
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_virtual_network.test_vcn.id}"

  # Provider code tries to maintain compatibility with old versions.
  security_list_ids = ["${oci_core_virtual_network.test_vcn.default_security_list_id}"]
  display_name      = "tfSubNet1ForClusters"
  route_table_id    = "${oci_core_route_table.test_route_table.id}"
}

resource "oci_core_subnet" "clusterSubnet_2" {
  #Required
  availability_domain = "${data.oci_identity_availability_domain.ad2.name}"
  cidr_block          = "10.0.21.0/24"
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_virtual_network.test_vcn.id}"
  display_name        = "tfSubNet1ForClusters"

  # Provider code tries to maintain compatibility with old versions.
  security_list_ids = ["${oci_core_virtual_network.test_vcn.default_security_list_id}"]
  route_table_id    = "${oci_core_route_table.test_route_table.id}"
}

resource "oci_core_subnet" "nodePool_Subnet_1" {
  #Required
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  cidr_block          = "10.0.22.0/24"
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_virtual_network.test_vcn.id}"

  # Provider code tries to maintain compatibility with old versions.
  security_list_ids = ["${oci_core_virtual_network.test_vcn.default_security_list_id}"]
  display_name      = "tfSubNet1ForNodePool"
  route_table_id    = "${oci_core_route_table.test_route_table.id}"
}

resource "oci_core_subnet" "nodePool_Subnet_2" {
  #Required
  availability_domain = "${data.oci_identity_availability_domain.ad2.name}"
  cidr_block          = "10.0.23.0/24"
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_virtual_network.test_vcn.id}"

  # Provider code tries to maintain compatibility with old versions.
  security_list_ids = ["${oci_core_virtual_network.test_vcn.default_security_list_id}"]
  display_name      = "tfSubNet2ForNodePool"
  route_table_id    = "${oci_core_route_table.test_route_table.id}"
}
