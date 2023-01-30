// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_id
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block     = "10.0.0.0/24"
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.test_vcn.id
}

variable "opensearch_cluster_data_node_count" {
  default = 1
}

variable "opensearch_cluster_data_node_host_bare_metal_shape" {
  default = "dataNodeHostBareMetalShape"
}

variable "opensearch_cluster_data_node_host_memory_gb" {
  default = 10
}

variable "opensearch_cluster_data_node_host_ocpu_count" {
  default = 2
}

variable "opensearch_cluster_data_node_host_type" {
  default = "FLEX"
}

variable "opensearch_cluster_data_node_storage_gb" {
  default = 50
}

variable "opensearch_cluster_display_name" {
  default = "OpensearchExampleCluster"
}

variable "opensearch_cluster_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "opensearch_cluster_id" {
  default = "id"
}

variable "opensearch_cluster_master_node_count" {
  default = 1
}

variable "opensearch_cluster_master_node_host_bare_metal_shape" {
  default = "masterNodeHostBareMetalShape"
}

variable "opensearch_cluster_master_node_host_memory_gb" {
  default = 16
}

variable "opensearch_cluster_master_node_host_ocpu_count" {
  default = 1
}

variable "opensearch_cluster_master_node_host_type" {
  default = "FLEX"
}

variable "opensearch_cluster_opendashboard_node_count" {
  default = 1
}

variable "opensearch_cluster_opendashboard_node_host_memory_gb" {
  default = 10
}

variable "opensearch_cluster_opendashboard_node_host_ocpu_count" {
  default = 2
}

variable "opensearch_cluster_software_version" {
  default = "1.2.4"
}

variable "opensearch_cluster_state" {
  default = "ACTIVE"
}

variable "opensearch_cluster_system_tags" {
  default = { }
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_opensearch_opensearch_cluster" "test_opensearch_cluster" {
  #Required
  compartment_id                     = var.compartment_id
  data_node_count                    = var.opensearch_cluster_data_node_count
  data_node_host_memory_gb           = var.opensearch_cluster_data_node_host_memory_gb
  data_node_host_ocpu_count          = var.opensearch_cluster_data_node_host_ocpu_count
  data_node_host_type                = var.opensearch_cluster_data_node_host_type
  data_node_storage_gb               = var.opensearch_cluster_data_node_storage_gb
  display_name                       = var.opensearch_cluster_display_name
  master_node_count                  = var.opensearch_cluster_master_node_count
  master_node_host_memory_gb         = var.opensearch_cluster_master_node_host_memory_gb
  master_node_host_ocpu_count        = var.opensearch_cluster_master_node_host_ocpu_count
  master_node_host_type              = var.opensearch_cluster_master_node_host_type
  opendashboard_node_count           = var.opensearch_cluster_opendashboard_node_count
  opendashboard_node_host_memory_gb  = var.opensearch_cluster_opendashboard_node_host_memory_gb
  opendashboard_node_host_ocpu_count = var.opensearch_cluster_opendashboard_node_host_ocpu_count
  software_version                   = var.opensearch_cluster_software_version
  subnet_compartment_id              = var.compartment_id
  subnet_id                          = oci_core_subnet.test_subnet.id
  vcn_compartment_id                 = var.compartment_id
  vcn_id                             = oci_core_vcn.test_vcn.id

  #Optional
  data_node_host_bare_metal_shape   = var.opensearch_cluster_data_node_host_bare_metal_shape
  #  defined_tags                      = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.opensearch_cluster_defined_tags_value)
  freeform_tags                     = var.opensearch_cluster_freeform_tags
  master_node_host_bare_metal_shape = var.opensearch_cluster_master_node_host_bare_metal_shape
  #  system_tags                       = var.opensearch_cluster_system_tags
}

data "oci_opensearch_opensearch_clusters" "test_opensearch_clusters" {
  #Required
  compartment_id = var.compartment_id
  #Optional
  #  display_name = var.opensearch_cluster_display_name
  #  id           = var.opensearch_cluster_id
  #  state        = var.opensearch_cluster_state
}