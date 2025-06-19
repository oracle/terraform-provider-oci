// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "region" {
  default = "us-ashburn-1"
}

variable "compartment_id" {
  default  = "compartment.r47eayq"
}
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

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}

variable "redis_cluster_create_identity_token_defined_tags_value" {
  default = "definedTags"
}


variable "redis_cluster_create_identity_token_public_key" {
  default = "ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"
}

variable "redis_cluster_create_identity_token_redis_user" {
  default = "OCI_REDIS_OWNER"
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
  //  defined_tags  = 
}



resource "oci_redis_redis_cluster_create_identity_token" "test_redis_cluster_create_identity_token" {
  #Required
  public_key       = var.redis_cluster_create_identity_token_public_key
  redis_cluster_id = oci_redis_redis_cluster.test_redis_cluster.id
  redis_user       = var.redis_cluster_create_identity_token_redis_user

  #Optional
}


