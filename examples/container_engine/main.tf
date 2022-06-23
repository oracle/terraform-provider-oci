// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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

variable "region" {
  default = "us-ashburn-1"
}

# Provide the SSH public key to be set on each node in the node pool on launch.
variable "node_pool_ssh_public_key" {

}

variable "kms_vault_id" {

}

variable "node_pool_node_eviction_node_pool_settings_eviction_grace_duration" {
  default = "PT50M"
}

variable "node_pool_node_eviction_node_pool_settings_is_force_delete_after_grace_duration" {
  default = false
}

variable "node_pool_state" {
  default = []
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

data "oci_kms_vault" "test_vault" {
  #Required
  vault_id = var.kms_vault_id
}

data "oci_kms_keys" "test_keys_dependency_RSA" {
  #Required
  compartment_id = var.tenancy_ocid
  management_endpoint = data.oci_kms_vault.test_vault.management_endpoint
  algorithm = "RSA"

  filter {
    name = "state"
    values = ["ENABLED", "UPDATING"]
  }
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

resource "oci_core_subnet" "nodePool_Subnet_2" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad2.name
  cidr_block          = "10.0.23.0/24"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id

  # Provider code tries to maintain compatibility with old versions.
  security_list_ids = [oci_core_vcn.test_vcn.default_security_list_id]
  display_name      = "tfSubNet2ForNodePool"
  route_table_id    = oci_core_route_table.test_route_table.id
}

resource "oci_containerengine_cluster" "test_cluster" {
  #Required
  compartment_id     = var.compartment_ocid
  kubernetes_version = data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions[0]
  name               = "tfTestCluster"
  vcn_id             = oci_core_vcn.test_vcn.id

  #Optional
  image_policy_config {
    is_policy_enabled = "true"
    key_details {
      kms_key_id = data.oci_kms_keys.test_keys_dependency_RSA.keys[0].id
    }
  }

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
  kubernetes_version = data.oci_containerengine_node_pool_option.test_node_pool_option.kubernetes_versions[0]
  name               = "tfPool"
  node_shape         = "VM.Standard2.1"
  subnet_ids         = [oci_core_subnet.nodePool_Subnet_1.id, oci_core_subnet.nodePool_Subnet_2.id]

  #Optional
  initial_node_labels {
    #Optional
    key   = "key"
    value = "value"
  }

  node_eviction_node_pool_settings {
    #Optional
    eviction_grace_duration              = var.node_pool_node_eviction_node_pool_settings_eviction_grace_duration
    is_force_delete_after_grace_duration = var.node_pool_node_eviction_node_pool_settings_is_force_delete_after_grace_duration
  }

  node_source_details {
    #Required
    image_id    = local.image_id
    source_type = "IMAGE"

    #Optional
    boot_volume_size_in_gbs = "60"
  }

  quantity_per_subnet = 2
  ssh_public_key      = var.node_pool_ssh_public_key
}

resource "oci_containerengine_node_pool" "test_flex_shape_node_pool" {
  #Required
  cluster_id         = oci_containerengine_cluster.test_cluster.id
  compartment_id     = var.compartment_ocid
  kubernetes_version = data.oci_containerengine_node_pool_option.test_node_pool_option.kubernetes_versions[0]
  name               = "flexShapePool"
  node_shape         = "VM.Standard.E3.Flex"
  subnet_ids         = [oci_core_subnet.nodePool_Subnet_1.id, oci_core_subnet.nodePool_Subnet_2.id]

  node_source_details {
    #Required
    image_id    = local.oracle_linux_images.0
    source_type = "IMAGE"
  }

  node_shape_config {
    ocpus = 2
    memory_in_gbs = 40
  }

  quantity_per_subnet = 2
  ssh_public_key      = var.node_pool_ssh_public_key
}

output "cluster" {
  value = {
    id                 = oci_containerengine_cluster.test_cluster.id
    kubernetes_version = oci_containerengine_cluster.test_cluster.kubernetes_version
    name               = oci_containerengine_cluster.test_cluster.name
  }
}

output "node_pool" {
  value = {
    id                 = oci_containerengine_node_pool.test_node_pool.id
    kubernetes_version = oci_containerengine_node_pool.test_node_pool.kubernetes_version
    name               = oci_containerengine_node_pool.test_node_pool.name
    subnet_ids         = oci_containerengine_node_pool.test_node_pool.subnet_ids
  }
}

output "flex_node_pool" {
  value = {
    id                 = oci_containerengine_node_pool.test_flex_shape_node_pool.id
    kubernetes_version = oci_containerengine_node_pool.test_flex_shape_node_pool.kubernetes_version
    name               = oci_containerengine_node_pool.test_flex_shape_node_pool.name
  }
}

data "oci_containerengine_cluster_option" "test_cluster_option" {
  cluster_option_id = "all"
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

output "cluster_kubernetes_versions" {
  value = [data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions]
}

output "node_pool_kubernetes_version" {
  value = [data.oci_containerengine_node_pool_option.test_node_pool_option.kubernetes_versions]
}

data "oci_containerengine_cluster_kube_config" "test_cluster_kube_config" {
  #Required
  cluster_id = oci_containerengine_cluster.test_cluster.id

  #Optional
  token_version = "2.0.0"
}

resource "local_file" "test_cluster_kube_config_file" {
  content  = data.oci_containerengine_cluster_kube_config.test_cluster_kube_config.content
  filename = "${path.module}/test_cluster_kubeconfig"
}

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}

data "oci_containerengine_node_pools" "test_node_pools" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  state      = var.node_pool_state
}

variable "InstanceImageOCID" {
  type = map(string)

  default = {
    // See https://docs.us-phoenix-1.oraclecloud.com/images/
    // Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
    us-phoenix-1   = "ocid1.image.oc1.phx.aaaaaaaadjnj3da72bztpxinmqpih62c2woscbp6l3wjn36by2cvmdhjub6a"
    us-ashburn-1   = "ocid1.image.oc1.iad.aaaaaaaawufnve5jxze4xf7orejupw5iq3pms6cuadzjc7klojix6vmk42va"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaagbrvhganmn7awcr7plaaf5vhabmzhx763z5afiitswjwmzh7upna"
    uk-london-1    = "ocid1.image.oc1.uk-london-1.aaaaaaaajwtut4l7fo3cvyraate6erdkyf2wdk5vpk6fp6ycng3dv2y3ymvq"
  }
}

