// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "region" {
  default = "us-ashburn-1"
}
variable "compartment_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  auth             = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region           = var.region
}

data "oci_identity_availability_domain" "ad1" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

data "oci_identity_availability_domain" "ad2" {
  compartment_id = var.tenancy_ocid
  ad_number      = 2
}

data "oci_containerengine_cluster_option" "test_cluster_option" {
  cluster_option_id = "all"
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "tfDualStackVcnForClusters"
  is_ipv6enabled = true
  is_oracle_gua_allocation_enabled = true
}


resource "oci_core_internet_gateway" "test_ig" {
  compartment_id = var.compartment_ocid
  display_name   = "tfClusterInternetGateway"
  vcn_id         = oci_core_vcn.test_vcn.id
}
resource "oci_core_route_table" "test_route_table" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
  display_name   = "tfClustersRouteTable"
  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_internet_gateway.test_ig.id
  }
}
resource "oci_core_subnet" "test_subnet" {
  #Required
  cidr_block          = "10.0.20.0/24"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id
  # Provider code tries to maintain compatibility with old versions.
  security_list_ids = [oci_core_vcn.test_vcn.default_security_list_id]
  display_name      = "tfSubNet1ForNodePool"
  route_table_id    = oci_core_route_table.test_route_table.id
  ipv6cidr_block   = cidrsubnet(oci_core_vcn.test_vcn.ipv6cidr_blocks[0], 8, 1)  # Creating a /64 subnet from /56
}
resource "oci_core_security_list" "test_security_list" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
  display_name   = "Default Security List for virtual node pool"
  egress_security_rules {
    destination      = "0.0.0.0/0"
    destination_type = "CIDR_BLOCK"
    protocol         = "all"
    stateless        = false
    description      = "Allowing egress to all via all protocols."
  }
  ingress_security_rules {
    source           = "10.0.0.0/16"
    source_type      = "CIDR_BLOCK"
    protocol         = "all"
    stateless        = false
  }
  ingress_security_rules {
    protocol    = 6 # local.TCP
    source      = "0.0.0.0/0"
    source_type = "CIDR_BLOCK"
    stateless   = false
    description = "Allowing ingress to all via TCP"
    # Optional
    tcp_options {
      max = "6443"
      min = "6443"
      source_port_range {
        max = "1521"
        min = "1521"
      }
    }
  }
  ingress_security_rules {
    # Optional
    icmp_options {
      code = "4"
      type = "3"
    }
    protocol    = 1 # local.ICMP
    source      = "0.0.0.0/0"
    source_type = "CIDR_BLOCK"
    stateless   = false
    description = "Allowing ingress to all via ICMP"
  }
  ingress_security_rules {
    # Optional
    icmp_options {
      code = "-1"
      type = "3"
    }
    protocol    = 1 # local.ICMP
    source      = "10.0.0.0/16"
    source_type = "CIDR_BLOCK"
    stateless   = false
  }
}

# Create a dual stack cluster
resource "oci_containerengine_cluster" "test_cluster" {
  #Required
  compartment_id     = var.compartment_ocid
  kubernetes_version = reverse(data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions)[0]
  name               = "tfTestClusterDualStack"
  vcn_id             = oci_core_vcn.test_vcn.id
  type               = "ENHANCED_CLUSTER"
  cluster_pod_network_options {
    # VNPs require cni_type as OCI_VCN_IP_NATIVE
    cni_type = "OCI_VCN_IP_NATIVE"
  }
  options {
    ip_families = ["IPv4", "IPv6"]
  }
  endpoint_config {
    #Optional
    is_public_ip_enabled = true
    nsg_ids              = ["${oci_core_network_security_group.network_security_group_rd.id}"]
    subnet_id            = oci_core_subnet.test_subnet.id
  }
}
resource "oci_core_network_security_group" "network_security_group_rd" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
  display_name   = "displayName"
}