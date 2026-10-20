// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License v2.0

variable "auth" {
  description = "OCI authentication method."
  type        = string
  default     = "SecurityToken"
}

variable "config_file_profile" {
  description = "OCI CLI configuration profile used by the provider."
  type        = string
  default     = "DEFAULT"
}

variable "region" {
  description = "Secondary cluster region."
  type        = string
  default     = "us-phoenix-1"
}

variable "compartment_id" {
  description = "Compartment OCID in which to create the secondary cluster and its network resources."
  type        = string
}

variable "primary_cluster_id" {
  description = "Active primary Redis cluster OCID in the paired primary region."
  type        = string
}

variable "secondary_oci_cache_config_set_id" {
  description = "OCI Cache config set OCID in the secondary region. It must be CRR-compatible with the primary cluster's config set."
  type        = string
}

variable "display_name" {
  description = "Display name for the secondary cluster."
  type        = string
  default     = "redis-crr-secondary-example"
}

variable "node_count" {
  description = "Number of nodes in the non-sharded secondary cluster."
  type        = number
  default     = 1
}

variable "node_memory_in_gbs" {
  description = "Memory, in GB, allocated to each node."
  type        = number
  default     = 8
}

variable "software_version" {
  description = "Redis software version for the secondary cluster."
  type        = string
  default     = "REDIS_7_0"
}

variable "freeform_tags" {
  description = "Optional freeform tags for the secondary cluster."
  type        = map(string)
  default = {
    "example" = "redis-crr"
  }
}

provider "oci" {
  auth                = var.auth
  config_file_profile = var.config_file_profile
  region              = var.region
}

resource "oci_core_vcn" "secondary" {
  compartment_id = var.compartment_id
  cidr_block     = "10.0.0.0/16"
  display_name   = "redis-crr-secondary-vcn"
}

resource "oci_core_security_list" "secondary" {
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.secondary.id
  display_name   = "redis-security-list"

  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol    = "all"
  }
}

resource "oci_core_subnet" "secondary" {
  compartment_id    = var.compartment_id
  vcn_id            = oci_core_vcn.secondary.id
  cidr_block        = "10.0.0.0/24"
  security_list_ids = [oci_core_security_list.secondary.id]
  display_name      = "redis-crr-secondary-subnet"
}

# The external primary cluster is not created or managed by this example.
# Supplying primary_cluster_id creates this cluster as a CRR secondary.
resource "oci_redis_redis_cluster" "secondary" {
  compartment_id          = var.compartment_id
  display_name            = var.display_name
  node_count              = var.node_count
  node_memory_in_gbs      = var.node_memory_in_gbs
  software_version        = var.software_version
  subnet_id               = oci_core_subnet.secondary.id
  cluster_mode            = "NONSHARDED"
  primary_cluster_id      = var.primary_cluster_id
  oci_cache_config_set_id = var.secondary_oci_cache_config_set_id
  freeform_tags           = var.freeform_tags
}

output "secondary_cluster_id" {
  description = "OCID of the CRR secondary cluster created by this example."
  value       = oci_redis_redis_cluster.secondary.id
}

output "secondary_cluster_role" {
  description = "Role reported by OCI for the created cluster."
  value       = oci_redis_redis_cluster.secondary.cluster_role
}
