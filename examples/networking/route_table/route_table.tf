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

variable "route_table_route_rules_description" {
  default = "description"
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

resource "oci_core_internet_gateway" "example_ig" {
  compartment_id = var.compartment_ocid
  display_name   = "exampleIG"
  vcn_id         = oci_core_vcn.example_vcn.id
}

resource "oci_core_route_table" "example_route_table" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.example_vcn.id
  display_name   = "exampleRouteTable"

  route_rules {
    description       = var.route_table_route_rules_description
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_internet_gateway.example_ig.id
  }
}

