// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "region" {
  default = "ap-mumbai-1"
}
variable "compartment_id" {}

variable "oci_cache_backup_export_to_object_storage_bucket" {
  default = "test-bucket"
}

variable "oci_cache_backup_export_to_object_storage_namespace" {
  default = "id5p2j2htymo"
}

variable "redis_cluster_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "redis_cluster_node_count" {
  default = 2
}

variable "redis_cluster_node_memory_in_gbs" {
  default = 2.0
}

variable "redis_cluster_software_version" {
  default = "VALKEY_7_2"
}

variable "oci_cache_backup_backup_source" {
  default = "REPLICA"
}

variable "oci_cache_backup_defined_tags_value" {
  default = "value"
}

variable "oci_cache_backup_description" {
  default = "description"
}

variable "oci_cache_backup_display_name" {
  default = "displayName"
}

variable "oci_cache_backup_display_name_sharded" {
  default = "displayNameSharded"
}

variable "oci_cache_backup_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "oci_cache_backup_retention_period_in_days" {
  default = 10
}

variable "oci_cache_backup_state" {
  default = "AVAILABLE"
}

provider "oci" {
  auth = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region = var.region
}

variable "redis_cluster_display_name" {
  type = string
  default = "test-tf-redis-cluster"
}

variable "redis_sharded_cluster_display_name" {
  type = string
  default = "test-tf-redis-sharded_example"
}

variable "oci_cache_backup_export_to_object_storage_prefix" {
  default = "prefix"
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_id
}

resource "oci_core_security_list" "test_security_list" {
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.test_vcn.id
  display_name   = "redis-security-list"

  // allow outbound udp traffic on a port range
  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol    = "17" // udp
    stateless   = true
  }

  // allow inbound ssh traffic from a specific port
  ingress_security_rules {
    protocol  = "6" // tcp
    source    = "0.0.0.0/0"
    stateless = false
  }
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block     = "10.0.0.0/24"
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.test_vcn.id
  security_list_ids = [oci_core_security_list.test_security_list.id]
}

resource "oci_redis_redis_cluster" "test_redis_cluster" {
  #Required
  compartment_id     = var.compartment_id
  display_name       = var.redis_cluster_display_name
  node_count         = var.redis_cluster_node_count
  node_memory_in_gbs = var.redis_cluster_node_memory_in_gbs
  software_version   = var.redis_cluster_software_version
  subnet_id          = oci_core_subnet.test_subnet.id

  #Optional
  //  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.redis_cluster_defined_tags_value)
  freeform_tags = var.redis_cluster_freeform_tags
}

// create a 3 shard 2 node per shard
resource "oci_redis_redis_cluster" "test_redis_sharded_cluster" {
  #Required
  compartment_id     = var.compartment_id
  display_name       = var.redis_sharded_cluster_display_name
  node_count         = 2
  node_memory_in_gbs = var.redis_cluster_node_memory_in_gbs
  software_version   = var.redis_cluster_software_version
  subnet_id          = oci_core_subnet.test_subnet.id

  #Optional
  cluster_mode       = "SHARDED"
  shard_count        = 3
  freeform_tags = var.redis_cluster_freeform_tags
}

resource "oci_redis_oci_cache_backup" "test_oci_cache_backup" {
  compartment_id    = var.compartment_id
  display_name      = var.oci_cache_backup_display_name
  source_cluster_id = oci_redis_redis_cluster.test_redis_cluster.id
  backup_source            = var.oci_cache_backup_backup_source
  description              = var.oci_cache_backup_description
  freeform_tags            = var.oci_cache_backup_freeform_tags
  retention_period_in_days = var.oci_cache_backup_retention_period_in_days
}

resource "oci_redis_oci_cache_backup" "test_oci_cache_backup_sharded" {
  compartment_id    = var.compartment_id
  display_name      = var.oci_cache_backup_display_name_sharded
  source_cluster_id = oci_redis_redis_cluster.test_redis_sharded_cluster.id
  backup_source            = var.oci_cache_backup_backup_source
  description              = var.oci_cache_backup_description
  freeform_tags            = var.oci_cache_backup_freeform_tags
  retention_period_in_days = var.oci_cache_backup_retention_period_in_days
}

resource "oci_redis_oci_cache_backup_export_to_object_storage" "test_oci_cache_backup_export_to_object_storage" {
  #Required
  bucket              = var.oci_cache_backup_export_to_object_storage_bucket
  namespace           = var.oci_cache_backup_export_to_object_storage_namespace
  oci_cache_backup_id = oci_redis_oci_cache_backup.test_oci_cache_backup.id

  #Optional
  prefix = var.oci_cache_backup_export_to_object_storage_prefix
}