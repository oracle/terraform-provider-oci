// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

variable "pe_defined_tags_value" {
  default = "pe_tag_value"
}

variable "pe_freeform_tags" {
  default = { "bar-key" = "value" }
}

# Create a new Tag Namespace.
resource "oci_identity_tag_namespace" "tag_namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = "example tag namespace"
  name           = "example-tag-namespace-all"
}

# Create a new Tag definition in the above Tag Namespace.
resource "oci_identity_tag" "tag1" {
  #Required
  description      = "example tag"
  name             = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag_namespace1.id
}

# Create a new Virtual Cloud Network (VCN) resource.
resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_id
  display_name   = "exampleVCN"
  dns_label      = "tfexamplevcn"
}

# Create a Subset in the above VCN.
resource "oci_core_subnet" "test_subnet" {
  cidr_block     = "10.0.0.0/24"
  display_name   = "regionalSubnet"
  dns_label      = "regionalsubnet"
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.test_vcn.id
}

# Create a Network Security Group (NSG) in the above VCN.
resource "oci_core_network_security_group" "test_network_security_group" {
  #Required
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.test_vcn.id
}

# Create a new DB Management Private Endpoint.
resource "oci_database_management_db_management_private_endpoint" "test_db_management_private_endpoint" {
  #Required
  compartment_id = var.compartment_id
  name = var.db_management_private_endpoint_name
  subnet_id = oci_core_subnet.test_subnet.id

  #Optional
  description = var.db_management_private_endpoint_description
  nsg_ids   = [oci_core_network_security_group.test_network_security_group.id]
  is_cluster  = var.db_management_private_endpoint_is_cluster
  defined_tags  = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.pe_defined_tags_value
  }
  freeform_tags = var.pe_freeform_tags
}

# Get DB Management Private Endpoint.
data "oci_database_management_db_management_private_endpoint" "test_db_management_private_endpoint" {
  db_management_private_endpoint_id = oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint.id
}

# List DB Management Private Endpoints.
data "oci_database_management_db_management_private_endpoints" "test_db_management_private_endpoints" {
  #Required
  compartment_id = var.compartment_id
}

# List DB Management Private Endpoints matching the given filter criteria.
data "oci_database_management_db_management_private_endpoints" "test_db_management_private_endpoints_with_name" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  name = var.db_management_private_endpoint_name
  vcn_id = oci_core_vcn.test_vcn.id
  state = var.db_management_private_endpoint_state
  is_cluster = var.db_management_private_endpoint_is_cluster
}
