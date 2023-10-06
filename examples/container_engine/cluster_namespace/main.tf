// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "region" {}
variable "compartment_ocid" {}
variable "auth" {}
variable "config_file_profile" {}

provider "oci" {
  region              = var.region
  tenancy_ocid        = var.tenancy_ocid
  auth                = var.auth
  config_file_profile = var.config_file_profile
}

resource "oci_core_vcn" "test_vcn" {
  compartment_id = var.compartment_ocid
  display_name   = "test_vcn"
  cidr_block     = "10.0.0.0/16"
}

resource "oci_core_security_list" "test_security_list" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
  display_name   = "test_security_list"
  ingress_security_rules {
    source           = "0.0.0.0/0"
    source_type      = "CIDR_BLOCK"
    protocol         = 6 # local.TCP
    stateless        = false
    tcp_options {
      max = "6443"
      min = "6443"
    }
  }
}

resource "oci_core_internet_gateway" "test_internet_gateway" {
  compartment_id = var.compartment_ocid
  display_name   = "test_internet_gateway"
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_core_route_table" "test_route_table" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
  display_name   = "test_route_table"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_internet_gateway.test_internet_gateway.id
  }
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block          = "10.0.20.0/24"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id

  display_name      = "test_subnet"
  security_list_ids = [oci_core_security_list.test_security_list.id]
  route_table_id    = oci_core_route_table.test_route_table.id
}

resource "oci_containerengine_cluster" "test_cluster" {
  compartment_id     = var.compartment_ocid
  kubernetes_version = "v1.27.2"
  name               = "test_cluster"
  vcn_id             = oci_core_vcn.test_vcn.id
  type               = "ENHANCED_CLUSTER"

  cluster_pod_network_options {
    cni_type = "OCI_VCN_IP_NATIVE"
  }

  endpoint_config {
    is_public_ip_enabled = true
    subnet_id = oci_core_subnet.test_subnet.id
  }
}

resource "oci_containerengine_cluster_namespace_profile" "test_cluster_namespace_profile" {
  compartment_id = var.compartment_ocid
  display_name   = "test_cluster_namespace_profile"
  namespace_suffix = "tf-test"
}

resource "oci_containerengine_cluster_attachment" "test_cluster_attachment" {
  cluster_id                   = oci_containerengine_cluster.test_cluster.id
  cluster_namespace_profile_id = oci_containerengine_cluster_namespace_profile.test_cluster_namespace_profile.id
  compartment_id               = var.compartment_ocid
  display_name                 = "test_cluster_attachment"
  depends_on                   = [oci_containerengine_cluster.test_cluster]
}

resource "oci_containerengine_cluster_namespace_profile_version" "test_cluster_namespace_profile_version" {
  admin_cluster_role_name      = "cluster-admin"
  cluster_namespace_profile_id = oci_containerengine_cluster_namespace_profile.test_cluster_namespace_profile.id
  compartment_id               = var.compartment_ocid
  name                         = "test_cluster_namespace_profile_version"
  depends_on                   = [oci_containerengine_cluster.test_cluster]
}

resource "oci_containerengine_cluster_namespace" "test_cluster_namespace" {
  cluster_namespace_profile_version_id = oci_containerengine_cluster_namespace_profile_version.test_cluster_namespace_profile_version.id
  compartment_id                       = var.compartment_ocid
  name                                 = "test-cluster-namespace"
  description                          = "test-cluster-namespace"
}