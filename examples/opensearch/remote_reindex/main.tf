// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// tfworkflow supplies required_providers_block.tf. See README.md for direct Terraform use.

variable "region" {
  type = string
}

variable "config_file_profile" {
  type = string
}

variable "compartment_id" {
  type = string
}

variable "display_name" {
  type    = string
  default = "opensearch-remote-reindex-example"
}

variable "software_version" {
  type    = string
  default = "2.15.0"
}

variable "security_master_user_name" {
  type = string
}

variable "security_master_user_password_hash" {
  type        = string
  description = "OpenSearch-compatible password hash for the cluster's master user, not a plaintext password."
  sensitive   = true
}

variable "remote_reindex_ca_bundle_pem_path" {
  type        = string
  description = "Path to a PEM file containing the CA certificates that validate the remote source's TLS certificate."
}

variable "remote_reindex_endpoint" {
  type        = string
  description = "HTTPS URL of the remote OpenSearch or Elasticsearch source, including its port if needed."
}

variable "remote_reindex_endpoint_resolution_ip" {
  type        = string
  description = "Optional IP address used to resolve the remote endpoint hostname from the cluster. The hostname is still used for TLS verification."
  default     = null
}

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = var.config_file_profile
  region              = var.region
}

resource "oci_core_vcn" "opensearch" {
  compartment_id = var.compartment_id
  cidr_block     = "10.0.0.0/16"
  display_name   = "${var.display_name}-vcn"
}

resource "oci_core_subnet" "opensearch" {
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.opensearch.id
  cidr_block     = "10.0.0.0/24"
  display_name   = "${var.display_name}-subnet"
}

resource "oci_certificates_management_ca_bundle" "remote_reindex" {
  compartment_id = var.compartment_id
  name           = "${var.display_name}-ca-bundle"
  ca_bundle_pem  = file(pathexpand(var.remote_reindex_ca_bundle_pem_path))
}

resource "oci_opensearch_opensearch_cluster" "remote_reindex" {
  compartment_id                     = var.compartment_id
  display_name                       = var.display_name
  software_version                   = var.software_version
  vcn_id                             = oci_core_vcn.opensearch.id
  vcn_compartment_id                 = var.compartment_id
  subnet_id                          = oci_core_subnet.opensearch.id
  subnet_compartment_id              = var.compartment_id
  master_node_count                  = 1
  master_node_host_type              = "FLEX"
  master_node_host_ocpu_count        = 1
  master_node_host_memory_gb         = 20
  data_node_count                    = 1
  data_node_host_type                = "FLEX"
  data_node_host_ocpu_count          = 2
  data_node_host_memory_gb           = 20
  data_node_storage_gb               = 50
  coordinator_node_count             = 1
  coordinator_node_host_type         = "FLEX"
  coordinator_node_host_shape        = "VM.Standard.E3.Flex"
  coordinator_node_host_ocpu_count   = 2
  coordinator_node_host_memory_gb    = 20
  opendashboard_node_count           = 1
  opendashboard_node_host_ocpu_count = 2
  opendashboard_node_host_memory_gb  = 10
  security_mode                      = "ENFORCING"
  security_master_user_name          = var.security_master_user_name
  security_master_user_password_hash = var.security_master_user_password_hash

  // This reference creates the CA bundle before the cluster and destroys it after the cluster.
  remote_reindex_trust_config {
    remote_sources {
      ca_bundle_id           = oci_certificates_management_ca_bundle.remote_reindex.id
      endpoint               = var.remote_reindex_endpoint
      endpoint_resolution_ip = var.remote_reindex_endpoint_resolution_ip
    }
  }
}

output "cluster_id" {
  value = oci_opensearch_opensearch_cluster.remote_reindex.id
}

output "ca_bundle_id" {
  value = oci_certificates_management_ca_bundle.remote_reindex.id
}
