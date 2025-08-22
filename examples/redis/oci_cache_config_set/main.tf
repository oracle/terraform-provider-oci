// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "region" {
  default = "us-phoenix-1"
}

variable "compartment_id" {}

variable "redis_cluster_display_name" {
  type = string
  default = "test-redis-cluster"
}

variable "redis_cluster_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "redis_cluster_node_count" {
  default = 5
}

variable "redis_cluster_node_memory_in_gbs" {
  default = 2.0
}

variable "redis_cluster_software_version" {
  default = "REDIS_7_0"
}


# OCI Cache ConfigSet vars
variable "oci_cache_config_set_display_name" {
  type = string
  default = "test-config-set"
}

variable "oci_cache_config_set_software_version" {
  type = string
  default = "REDIS_7_0"
}

variable "oci_cache_config_set_description" {
  type = string
  default = "Test Config Set created via Terraform"
}

variable "oci_cache_config_set_config_key1" {
  type = string
  default = "maxmemory-policy"
}

variable "oci_cache_config_set_config_value1" {
  type = string
  default = "allkeys-random"
}

variable "oci_cache_config_set_config_key2" {
  type = string
  default = "notify-keyspace-events"
}

variable "oci_cache_config_set_config_value2" {
  type = string
  default = "KEA"
}

variable "oci_cache_config_set_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "oci_cache_config_set_state" {
  default = "ACTIVE"
}

provider "oci" {
  auth = "SecurityToken"
  config_file_profile = "DEFAULT"
  region = var.region
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
  oci_cache_config_set_id = oci_redis_oci_cache_config_set.test_oci_cache_config_set_redis.id
  freeform_tags = var.redis_cluster_freeform_tags
}


resource "oci_redis_oci_cache_config_set" "test_oci_cache_config_set_redis" {
  #Required
  compartment_id = var.compartment_id
  configuration_details {
    #Required
    items {
      #Required
      config_key   = var.oci_cache_config_set_config_key1
      config_value = var.oci_cache_config_set_config_value1
    }
    items {
      #Required
      config_key   = var.oci_cache_config_set_config_key2
      config_value = var.oci_cache_config_set_config_value2
    }
  }
  display_name     = var.oci_cache_config_set_display_name
  software_version = var.oci_cache_config_set_software_version

  #Optional
  // TODO defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.oci_cache_config_set_defined_tags_value)
  description   = var.oci_cache_config_set_description
  freeform_tags = var.oci_cache_config_set_freeform_tags
}

data "oci_redis_oci_cache_config_sets" "test_oci_cache_config_sets" {
  #Optional
  compartment_id   = var.compartment_id
  display_name     = var.oci_cache_config_set_display_name
  id = oci_redis_oci_cache_config_set.test_oci_cache_config_set_redis.id
  software_version = var.oci_cache_config_set_software_version
  state = var.oci_cache_config_set_state
}

data "oci_redis_oci_cache_config_set" "test_oci_cache_config_set" {
  oci_cache_config_set_id = oci_redis_oci_cache_config_set.test_oci_cache_config_set_redis.id
}

resource "oci_redis_oci_cache_config_setlist_associated_oci_cache_cluster" "test_oci_cache_config_setlist_associated_oci_cache_cluster" {
  #Required
  oci_cache_config_set_id = oci_redis_oci_cache_config_set.test_oci_cache_config_set_redis.id
}

data "oci_redis_oci_cache_default_config_sets" "test_oci_cache_default_config_sets" {
  #Required
  compartment_id   = var.compartment_id
}

data "oci_redis_oci_cache_default_config_set" "test_oci_cache_default_config_set" {
  #Required
  compartment_id   = var.compartment_id
  oci_cache_default_config_set_id = data.oci_redis_oci_cache_default_config_sets.test_oci_cache_default_config_sets.id
}
