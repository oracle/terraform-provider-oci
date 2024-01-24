// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "region" {
  default = "us-ashburn-1"
}

variable "compartment_id" {
  default = "ocid1.compartment.oc1..aaaaaaaarnj7d6pfp2c5op4oct6qfh7553xzxezc2afmxxboqdov23a7cp2a"
}

variable "redis_cluster_display_name" {
  type = string
  default = "test-tf-redis-example"
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
  default = "V7_0_5"
}

variable "redis_cluster_state" {
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
  //  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.redis_cluster_defined_tags_value)
  freeform_tags = var.redis_cluster_freeform_tags
}

data "oci_redis_redis_clusters" "test_redis_clusters" {

  #Optional
  compartment_id = var.compartment_id
  display_name   = var.redis_cluster_display_name
  id             = oci_redis_redis_cluster.test_redis_cluster.id
  state          = var.redis_cluster_state
}
