// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

provider "oci" {
  tenancy_ocid = var.tenancy_ocid
  user_ocid = var.user_ocid
  fingerprint = var.fingerprint
  private_key_path = var.private_key_path
  region = var.region
}

variable "compartment_id" {  
  default = "<compartment.ocid>"
}

variable "db_management_private_endpoint_name" {
  default = "TestPrivateEndpoint"
}

variable "db_management_private_endpoint_description" {
  default = "Test private endpoint"
}

variable "db_management_private_endpoint_state" {
  default = "ACTIVE"
}

variable "db_management_private_endpoint_is_cluster" {
  default = false
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_id
  display_name   = "exampleVCN"
  dns_label      = "tfexamplevcn"
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block     = "10.0.0.0/24"
  display_name   = "regionalSubnet"
  dns_label      = "regionalsubnet"
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_core_network_security_group" "test_network_security_group" {
  #Required
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.test_vcn.id
}



resource "oci_database_management_db_management_private_endpoint" "test_db_management_private_endpoint" {
  #Required
  compartment_id = var.compartment_id
  name = var.db_management_private_endpoint_name
  subnet_id = oci_core_subnet.test_subnet.id

  #Optional
  description = var.db_management_private_endpoint_description
  nsg_ids   = [oci_core_network_security_group.test_network_security_group.id]
  is_cluster  = var.db_management_private_endpoint_is_cluster
}

data "oci_database_management_db_management_private_endpoint" "test_db_management_private_endpoint" {
  db_management_private_endpoint_id = oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint.id
}

data "oci_database_management_db_management_private_endpoints" "test_db_management_private_endpoints" {
  #Required
  compartment_id = var.compartment_id
}

data "oci_database_management_db_management_private_endpoints" "test_db_management_private_endpoints_with_name" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  name = var.db_management_private_endpoint_name
  vcn_id = oci_core_vcn.test_vcn.id
  state = var.db_management_private_endpoint_state
  is_cluster = var.db_management_private_endpoint_is_cluster
}