// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0



resource "oci_core_vcn" "test_vcn_a" {
  // Required
  cidr_block = "10.0.0.0/16"
  compartment_id = var.compartment_ocid

  // Optional
  display_name = "MyTestVcnA"
  dns_label = "dnslabelA"
}


resource "oci_core_vcn" "test_vcn_b" {
  // Required
  cidr_block = "20.0.0.0/16"
  compartment_id = var.compartment_ocid

  // Optional
  display_name = "MyTestVcnB"
  dns_label = "dnslabelB"
}


resource "oci_core_drg" "test_vcn_drg" {
  // Required
  compartment_id = var.compartment_ocid

  // Optional
  display_name = "MyTestVcnDrg"
}

resource "oci_core_drg_route_distribution" "test_vcn_drg_route_distribution" {
  // Required
  drg_id = oci_core_drg.test_vcn_drg.id
  distribution_type = "IMPORT"

  // optional
  display_name = "MyTestVcnDrgRouteDistribution"

}

resource "oci_core_drg_attachment" "test_vcn_drg_attachment_a" {
  // Required
  drg_id = oci_core_drg.test_vcn_drg.id
  vcn_id = oci_core_vcn.test_vcn_a.id

  // Optional
  drg_route_table_id = oci_core_drg_route_table.test_vcn_drg_route_table.id

}

resource "oci_core_drg_attachment" "test_vcn_drg_attachment_b" {
  // Required
  drg_id = oci_core_drg.test_vcn_drg.id
  vcn_id = oci_core_vcn.test_vcn_b.id

  // Optional
  drg_route_table_id = oci_core_drg_route_table.test_vcn_drg_route_table.id

}

resource "oci_core_drg_route_distribution_statement" "test_vcn_drg_route_distribution_statements" {
  // Required
  drg_route_distribution_id = oci_core_drg_route_distribution.test_vcn_drg_route_distribution.id
  action = "ACCEPT"

  match_criteria {
    match_type= "DRG_ATTACHMENT_TYPE"
    attachment_type = "VCN"
  }

  priority = 10


}

resource "oci_core_drg_route_distribution_statement" "test_vcn_drg_route_distribution_statements_empty" {
  // Required
  drg_route_distribution_id = oci_core_drg_route_distribution.test_vcn_drg_route_distribution.id
  action = "ACCEPT"

  match_criteria {
  }

  priority = 50


}

data "oci_core_drg_route_distribution" "test_vcn_drg_route_distribution_data" {
  // Required
  drg_route_distribution_id = oci_core_drg_route_distribution.test_vcn_drg_route_distribution.id
}

resource "oci_core_drg_route_table" "test_vcn_drg_route_table" {
  drg_id = oci_core_drg.test_vcn_drg.id

  // Optional
  import_drg_route_distribution_id = oci_core_drg_route_distribution.test_vcn_drg_route_distribution.id
  display_name = "MyTestVcnDrgRouteTable"
}

resource "oci_core_drg_route_table_route_rule" "test_vcn_drg_route_table_route_rule" {
  // Required
  drg_route_table_id = oci_core_drg_route_table.test_vcn_drg_route_table.id
  destination                = "10.0.0.0/8"
  destination_type           = "CIDR_BLOCK"
  next_hop_drg_attachment_id = oci_core_drg_attachment.test_vcn_drg_attachment_a.id
}

