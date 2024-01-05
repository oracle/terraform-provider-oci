// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
  default = "us-ashburn-1"
}


variable "tenancy_ocid" {
}

variable "cluster_id" {

}

variable "kubernetes_version" {

}

variable "compartment_ocid" {
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
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

resource "oci_containerengine_cluster" "test_cluster" {
    #Required
    compartment_id     = var.compartment_ocid
    kubernetes_version = reverse(data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions)[0]
    name               = "tfTestCluster"
    vcn_id             = oci_core_vcn.test_vcn.id
    type               = "ENHANCED_CLUSTER"
}

resource "oci_containerengine_addon" "dashboard" {
    #Required, a name uniquely identifies an add-on, see all supported add-on names in data.oci_containerengine_addon_options.all.addon_options
    addon_name = "KubernetesDashboard"
    #Required
    cluster_id = oci_containerengine_cluster.test_cluster.id
    #Required, remove the resource on addon deletion
    remove_addon_resources_on_delete = true
    dynamic configurations {
        for_each = local.addon_mappings

        content {
            key =configurations.value.key
            value = configurations.value.value
            }
        }
}

resource "oci_containerengine_node_pool" "test_node_pool" {
    #Required
    cluster_id         = oci_containerengine_cluster.test_cluster.id
    compartment_id     = var.compartment_ocid
    kubernetes_version = reverse(data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions)[0]
    name               = "tfPool"
    node_shape         = "VM.Standard2.1"

    node_config_details  {
        size = 1
        placement_configs {
            availability_domain = data.oci_identity_availability_domain.ad1.name
            subnet_id           = oci_core_subnet.nodePool_Subnet_1.id
        }
    }

    node_source_details {
        #Required
        image_id    = local.image_id
        source_type = "IMAGE"

        #Optional
        boot_volume_size_in_gbs = "60"
    }

    //use terraform depends_on to enforce cluster->add-on->node pool DAG
    depends_on = [oci_containerengine_addon.dashboard]
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

  addon_mappings = {
        mapping1 = {
            key = "numOfReplicas"
            value = "1"
        }
        mapping2 = {
            key = "nodeSelectors"
            value = "{\"pool\":\"system\"}"
        }
  }
}