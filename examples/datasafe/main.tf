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

variable "region" {
}

variable "compartment_ocid" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "exampleVCN"
  dns_label      = "tfexamplevcn"
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block     = "10.0.1.0/24"
  display_name   = "regionalSubnet"
  dns_label      = "regionalsubnet"
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_data_safe_data_safe_configuration" "test_data_safe_configuration" {
  is_enabled = "true"
}

resource "oci_data_safe_data_safe_private_endpoint" "test_data_safe_private_endpoint" {
  compartment_id = var.compartment_ocid
  display_name   = "PE2"
  subnet_id      = oci_core_subnet.test_subnet.id
  vcn_id         = oci_core_vcn.test_vcn.id
}

data "oci_data_safe_data_safe_private_endpoint" "test_data_safe_private_endpoint" {
  data_safe_private_endpoint_id = oci_data_safe_data_safe_private_endpoint.test_data_safe_private_endpoint.id
}

data "oci_data_safe_data_safe_private_endpoints" "test_data_safe_private_endpoints" {
  compartment_id = var.compartment_ocid
}

