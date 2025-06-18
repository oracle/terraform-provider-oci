// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "region" {
  default = "us-ashburn-1"
}

variable "compartment_id" {}
 #OciCacheCluster vars

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


variable "redis_cluster_display_name" {
  type = string
  default = "test-tf-redis-cluster"
}

#OciCacheUser vars
variable "oci_cache_user_description" {
  type = string
  default = "Test Cache user created via Terraform"
}

variable "oci_cache_user_acl_string" {
  type = string
  default = "~* +get +set"
}

variable "oci_cache_user_password_auth_name" {
  type = string
  default = "test-tf-pwd-user"
}

variable "oci_cache_user_iam_auth_name" {
  type = string
  default = "test-tf-iam-user"
}

variable "oci_cache_user_freeform_tags" {
  default = { "department" = "engineering" }
}

variable "oci_cache_user_status" {
  default = "ON"
}

variable "redis_cluster_get_oci_cache_user_display_name" {
  type = string
  default = "test-tf-oci-cache-users"
}

variable "oci_cache_user_get_redis_cluster_display_name" {
  type = string
  default = "test-tf-redis-clusters"
}

# Create a list of user IDs for attaching to the cluster
locals {
  cache_user_ids = [
    oci_redis_oci_cache_user.test_cache_user_password.id,
    oci_redis_oci_cache_user.test_cache_user_iam.id
  ]
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


resource "oci_redis_oci_cache_user" "test_cache_user_password" {
  #Required
  compartment_id = var.compartment_id
  name = var.oci_cache_user_password_auth_name
  description = var.oci_cache_user_description
  acl_string = var.oci_cache_user_acl_string

  authentication_mode {
    authentication_type = "PASSWORD"
    hashed_passwords = ["741f67765bef6f01f37bf5cb1724509a83409324efa6ad2586d27f4e3edea296"]
  }

  #Optional
  freeform_tags = var.oci_cache_user_freeform_tags
  status = var.oci_cache_user_status
  depends_on = [
    oci_redis_redis_cluster.test_redis_cluster
  ]
}

resource "oci_redis_oci_cache_user" "test_cache_user_iam" {
  #Required
  compartment_id = var.compartment_id
  name = var.oci_cache_user_iam_auth_name
  description = "${var.oci_cache_user_description} with IAM authentication"
  acl_string = "~* +@read"

  authentication_mode {
    authentication_type = "IAM"
  }

  #Optional
  freeform_tags = var.oci_cache_user_freeform_tags
  status = var.oci_cache_user_status
  depends_on = [
    oci_redis_redis_cluster.test_redis_cluster
  ]
}

# Attach both users to the Redis cluster
resource "oci_redis_redis_cluster_attach_oci_cache_user" "test_redis_cluster_attach_oci_cache_user" {
  #Required
  oci_cache_users = local.cache_user_ids
  redis_cluster_id = oci_redis_redis_cluster.test_redis_cluster.id

  depends_on = [
    oci_redis_oci_cache_user.test_cache_user_password,
    oci_redis_oci_cache_user.test_cache_user_iam
  ]
}

# Get cache users attached to a specific Redis cluster
resource "oci_redis_redis_cluster_get_oci_cache_user" "test_redis_cluster_get_oci_cache_user" {
  #Required
  redis_cluster_id = oci_redis_redis_cluster.test_redis_cluster.id

  #Optional
  compartment_id = var.compartment_id
  display_name = var.redis_cluster_get_oci_cache_user_display_name

  depends_on = [
    oci_redis_redis_cluster_attach_oci_cache_user.test_redis_cluster_attach_oci_cache_user
  ]
}

# Get Redis clusters to which ociCacheUser is attached
resource "oci_redis_oci_cache_user_get_redis_cluster" "test_oci_cache_user_get_redis_cluster" {
  #Required
  oci_cache_user_id = oci_redis_oci_cache_user.test_cache_user_password.id

  #Optional
  compartment_id = var.compartment_id
  display_name = var.oci_cache_user_get_redis_cluster_display_name

  depends_on = [
    oci_redis_redis_cluster_attach_oci_cache_user.test_redis_cluster_attach_oci_cache_user
  ]
}

# Detach both users from the Redis cluster
resource "oci_redis_redis_cluster_detach_oci_cache_user" "test_redis_cluster_detach_oci_cache_user" {
  #Required
  oci_cache_users = local.cache_user_ids
  redis_cluster_id = oci_redis_redis_cluster.test_redis_cluster.id

  depends_on = [
    oci_redis_redis_cluster_get_oci_cache_user.test_redis_cluster_get_oci_cache_user,
    oci_redis_oci_cache_user_get_redis_cluster.test_oci_cache_user_get_redis_cluster
  ]
}

# Data source for retrieving information about all cache users in a compartment
data "oci_redis_oci_cache_users" "all_cache_users" {
  #Required
  compartment_id = var.compartment_id
}

# Data source for retrieving information about a specific cache user
data "oci_redis_oci_cache_user" "test_cache_user" {
  #Required
  oci_cache_user_id = oci_redis_oci_cache_user.test_cache_user_password.id
}

# Output the details of the created cache users and their relationships
output "password_auth_cache_user_id" {
  value = oci_redis_oci_cache_user.test_cache_user_password.id
}

output "iam_auth_cache_user_id" {
  value = oci_redis_oci_cache_user.test_cache_user_iam.id
}

output "attached_cache_users" {
  value = oci_redis_redis_cluster_get_oci_cache_user.test_redis_cluster_get_oci_cache_user
}

output "attached_redis_clusters" {
  value = oci_redis_oci_cache_user_get_redis_cluster.test_oci_cache_user_get_redis_cluster
}