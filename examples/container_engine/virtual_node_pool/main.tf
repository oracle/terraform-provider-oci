// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
variable "tenancy_ocid" {
}
variable "cluster_id" {
}
variable "kubernetes_version" {
  # VNPs are only supported on v1.25+
  default = "v1.25.4"
}
variable "compartment_ocid" {
}
variable "image_id" {
}
variable "virtual_node_pool_state" {
  default = []
}
variable "cluster_endpoint_config_is_public_ip_enabled" {
  default = false
}
data "oci_containerengine_virtual_node_pools" "test_virtual_node_pools" {
  #Required
  compartment_id = var.compartment_ocid
  #Optional
  state      = var.virtual_node_pool_state
}
/*
A complete example to setup a cluster, then configure add-ons, then create node pool.
*/
data "oci_identity_availability_domain" "ad1" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}
resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "tfVcnForClusters"
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
  cidr_block          = "10.0.22.0/24"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id
  # Provider code tries to maintain compatibility with old versions.
  security_list_ids = [oci_core_vcn.test_vcn.default_security_list_id]
  display_name      = "tfSubNet1ForNodePool"
  route_table_id    = oci_core_route_table.test_route_table.id
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
  #manage_default_resource_id = oci_core_vcn.test_vcn.default_security_list_id
}
resource "oci_containerengine_cluster" "test_cluster" {
  #Required
  compartment_id     = var.compartment_ocid
  kubernetes_version = var.kubernetes_version
  name               = "tfTestCluster"
  vcn_id             = oci_core_vcn.test_vcn.id
  type               = "ENHANCED_CLUSTER"
    cluster_pod_network_options {
        # VNPs require cni_type as OCI_VCN_IP_NATIVE
        cni_type = "OCI_VCN_IP_NATIVE"
    }
    endpoint_config {
        #Optional
        is_public_ip_enabled = var.cluster_endpoint_config_is_public_ip_enabled
        nsg_ids              = ["${oci_core_network_security_group.network_security_group_rd.id}"]
        subnet_id            = oci_core_subnet.test_subnet.id
    }
}
resource "oci_core_network_security_group" "network_security_group_rd" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
  display_name   = "displayName"
}
resource "oci_containerengine_virtual_node_pool" "test_virtual_node_pool" {
  #Required
  cluster_id         = oci_containerengine_cluster.test_cluster.id
  compartment_id     = var.compartment_ocid
  display_name       = "tfVirtualNodePool"
  placement_configurations {
    #Required
    availability_domain = data.oci_identity_availability_domain.ad1.name
    subnet_id           = oci_core_subnet.test_subnet.id
    fault_domain        = ["FAULT-DOMAIN-1"]
  }
  #Optional
  initial_virtual_node_labels {
    #Optional
    key   = "key"
    value = "value"
  }
  #Required
  pod_configuration {
    shape = "Pod.Standard.E4.Flex"
    subnet_id = oci_core_subnet.test_subnet.id
    # Optional
    # nsg_ids = ["${oci_core_network_security_group.network_security_group_rd.id}"]
  }
  # Optional - This property cannot be modified at this time, we will fix and update in future releases
  # nsg_ids = ["${oci_core_network_security_group.network_security_group_rd.id}"]
  #Required
  size = 1
  #Optional
  taints {
    #Optional
    key   = "key"
    value = "value"
    # Effect must be one of 'NoSchedule', 'PreferNoSchedule' or 'NoExecute'
    effect = "NoSchedule"
  }
  //use terraform depends_on to enforce cluster->virtual node pool DAG
  depends_on = [oci_containerengine_cluster.test_cluster]
}
data "oci_containerengine_cluster_option" "test_cluster_option" {
  cluster_option_id = "all"
}
data "oci_containerengine_pod_shapes" "test_pod_shapes" {
  compartment_id = var.compartment_ocid
}