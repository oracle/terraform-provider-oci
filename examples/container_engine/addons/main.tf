// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
variable "region" {
  default = "us-ashburn-1"
}

variable "tenancy_ocid" {
}

variable "config_file_profile" {
}

variable "compartment_ocid" {
}

provider "oci" {
  region              = var.region
  auth                = "SecurityToken"
  config_file_profile = var.config_file_profile
}

/*
A complete example to setup a cluster, then configure add-ons.
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

resource "oci_core_subnet" "api_endpoint_subnet" {
  #Required
  cidr_block          = "10.0.23.0/24"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id

  # Provider code tries to maintain compatibility with old versions.
  security_list_ids = [oci_core_vcn.test_vcn.default_security_list_id]
  display_name      = "apiEndpointSubnet"
  route_table_id    = oci_core_route_table.test_route_table.id
}

resource "oci_containerengine_cluster" "test_cluster" {
    #Required
    compartment_id     = var.compartment_ocid
    kubernetes_version = reverse(data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions)[0]
    name               = "tfTestCluster"
    vcn_id             = oci_core_vcn.test_vcn.id
    type               = "ENHANCED_CLUSTER"
    endpoint_config {
      subnet_id = oci_core_subnet.api_endpoint_subnet.id
    }
}

resource "oci_containerengine_addon" "dashboard" {
    #Required, a name uniquely identifies an add-on, see all supported add-on names in data.oci_containerengine_addon_options.all.addon_options
    addon_name = "KubernetesDashboard"
    #Required
    cluster_id = oci_containerengine_cluster.test_cluster.id
    #Required, remove the resource on addon deletion
    remove_addon_resources_on_delete = true

    #Optional, will override an existing installation if true and Addon already exists
    override_existing = false

    #Optional
    dynamic configurations {
        for_each = local.addon_mappings

        content {
            key =configurations.value.key
            value = configurations.value.value
            }
        }
}

data "oci_containerengine_cluster_option" "test_cluster_option" {
  cluster_option_id = "all"
}

locals {
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