// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "cluster_id" {

}

variable "kubernetes_version" {

}

variable "compartment_ocid" {
}

variable "image_id" {

}

data "oci_containerengine_addon_options" "all" {
    #Required
    kubernetes_version = var.kubernetes_version
}

data "oci_containerengine_addon_options" "name_filter_example" {
    #Required
    kubernetes_version = var.kubernetes_version
    #Optional, a name uniquely identifies an add-on, see all supported add-on names in data.oci_containerengine_addon_options.all.addon_options
    addon_name = "KubernetesDashboard"
}

resource "oci_containerengine_addon" "addon_resource_example" {
    #Required, a name uniquely identifies an add-on, see all supported add-on names in data.oci_containerengine_addon_options.all.addon_options
    addon_name = "KubernetesDashboard"
    #Required
    cluster_id = var.cluster_id
    #Required, false values keeps installed resources of the addon on deletion. Set to true to fully remove resources
    remove_addon_resources_on_delete = true

    /*
    configurations that are supported by the add-on specified by the addon_name, see all supported configurations in in data.oci_containerengine_addon_options.all.addon_options.
    Unless required by a specific add-on, most of add-ons only have optional configurations that allow customization.
    */
     configurations {

     }
    /*
    Optional, see all supported version in in data.oci_containerengine_addon_options.all.addon_options.
    It is highly recommended to not set this field to let service choose and manage addon version.
    */
    version = "v1.0.0"
}

data "oci_containerengine_addons" "addon_addon_data_source_list_example" {
    #Required
    cluster_id = var.cluster_id
}

data "oci_containerengine_addon" "addon_data_source_singular_example" {
    #Required
    cluster_id = var.cluster_id
    #Required, a name uniquely identifies an add-on, see all supported add-on names in data.oci_containerengine_addon_options.all.addon_options
    addon_name = "KubernetesDashboard"
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
    kubernetes_version = var.kubernetes_version
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
}

resource "oci_containerengine_node_pool" "test_node_pool" {
    #Required
    cluster_id         = oci_containerengine_cluster.test_cluster.id
    compartment_id     = var.compartment_ocid
    kubernetes_version = var.kubernetes_version
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
        image_id    = var.image_id
        source_type = "IMAGE"

        #Optional
        boot_volume_size_in_gbs = "60"
    }

    //use terraform depends_on to enforce cluster->add-on->node pool DAG
    depends_on = [oci_containerengine_addon.dashboard]
}