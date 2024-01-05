// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_ocid" {
}

variable "kms_vault_id" {
}

variable "region" {
  default = "us-ashburn-1"
}

# Provide the SSH public key to be set on each node in the node pool on launch.
variable "node_pool_ssh_public_key" {
  default = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"
}

variable "node_pool_node_config_details_is_pv_encryption_in_transit_enabled" {
  default = true
}

variable "node_pool_node_config_details_size" {
  default = 1
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

data "oci_identity_availability_domain" "ad1" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

data "oci_identity_availability_domain" "ad2" {
  compartment_id = var.tenancy_ocid
  ad_number      = 2
}

resource "oci_kms_vault" "test_vault" {
  #Required
  compartment_id = var.compartment_ocid
  display_name = "tf_test"
  vault_type     = "DEFAULT"
}

resource "oci_kms_key" "test_key" {
  #Required
  compartment_id = var.compartment_ocid
  display_name = "tf-test-key"
  key_shape {
    #Required
    algorithm = "AES"
    length = 32
  }
  management_endpoint = oci_kms_vault.test_vault.management_endpoint
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

resource "oci_core_subnet" "nodePool_Subnet_1" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad1.name
  cidr_block          = "10.0.22.0/24"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id

  # Provider code tries to maintain compatibility with old versions.
  security_list_ids = [oci_core_vcn.test_vcn.default_security_list_id]
  display_name      = "tfSubNet1ForNodePool"
  route_table_id    = oci_core_route_table.test_route_table.id
}

resource "oci_core_subnet" "clusterSubnet_1" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad1.name
  cidr_block          = "10.0.20.0/24"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id

  # Provider code tries to maintain compatibility with old versions.
  security_list_ids = [oci_core_vcn.test_vcn.default_security_list_id]
  display_name      = "tfSubNet1ForClusters"
  route_table_id    = oci_core_route_table.test_route_table.id
}

resource "oci_core_subnet" "clusterSubnet_2" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad2.name
  cidr_block          = "10.0.21.0/24"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id
  display_name        = "tfSubNet1ForClusters"

  # Provider code tries to maintain compatibility with old versions.
  security_list_ids = [oci_core_vcn.test_vcn.default_security_list_id]
  route_table_id    = oci_core_route_table.test_route_table.id
}

resource "oci_containerengine_cluster" "test_cluster" {
  #Required
  compartment_id     = var.compartment_ocid
  kubernetes_version = "v1.20.11"
  name               = "tfTestCluster"
  vcn_id             = oci_core_vcn.test_vcn.id

  #Optional
  options {
    service_lb_subnet_ids = [oci_core_subnet.clusterSubnet_1.id, oci_core_subnet.clusterSubnet_2.id]

    #Optional
    add_ons {
      #Optional
      is_kubernetes_dashboard_enabled = "true"
      is_tiller_enabled               = "true"
    }

    admission_controller_options {
      #Optional
      is_pod_security_policy_enabled = true
    }

    kubernetes_network_config {
      #Optional
      pods_cidr     = "10.1.0.0/16"
      services_cidr = "10.2.0.0/16"
    }
  }
}

resource "oci_containerengine_node_pool" "test_node_pool" {
  #Required
  cluster_id         = oci_containerengine_cluster.test_cluster.id
  compartment_id     = var.compartment_ocid
  kubernetes_version = "v1.20.11"
  name               = "tfPool"
  node_shape         = "VM.Standard2.1"

  #Optional
  initial_node_labels {
    #Optional
    key   = "key"
    value = "value"
  }

  node_source_details {
    #Required
    image_id    = local.image_id
    source_type = "IMAGE"
  }

  node_config_details {
    #Required
    placement_configs {
      #Required
      availability_domain = data.oci_identity_availability_domain.ad1.name
      subnet_id           = oci_core_subnet.nodePool_Subnet_1.id
    }
    size = var.node_pool_node_config_details_size

    #Optional
    is_pv_encryption_in_transit_enabled = var.node_pool_node_config_details_is_pv_encryption_in_transit_enabled
    kms_key_id                          = oci_kms_key.test_key.id
  }

  ssh_public_key      = var.node_pool_ssh_public_key
}

output "node_pool" {
  value = {
    id                 = oci_containerengine_node_pool.test_node_pool.id
    kubernetes_version = oci_containerengine_node_pool.test_node_pool.kubernetes_version
    name               = oci_containerengine_node_pool.test_node_pool.name
    subnet_ids         = oci_containerengine_node_pool.test_node_pool.subnet_ids
  }
}

data "oci_containerengine_node_pool_option" "test_node_pool_option" {
  node_pool_option_id = "all"
}

data "oci_core_images" "shape_specific_images" {
  #Required
  compartment_id = var.tenancy_ocid
  shape = "VM.Standard2.1"
}

locals {
  all_images = "${data.oci_core_images.shape_specific_images.images}"
  all_sources = "${data.oci_containerengine_node_pool_option.test_node_pool_option.sources}"

  compartment_images = [for image in local.all_images : image.id if length(regexall("Oracle-Linux-[0-9]*.[0-9]*-20[0-9]*",image.display_name)) > 0 ]

  oracle_linux_images = [for source in local.all_sources : source.image_id if length(regexall("Oracle-Linux-[0-9]*.[0-9]*-20[0-9]*",source.source_name)) > 0]

  image_id = tolist(setintersection( toset(local.compartment_images), toset(local.oracle_linux_images)))[0]

}