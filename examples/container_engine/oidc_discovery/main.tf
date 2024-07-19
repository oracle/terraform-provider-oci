// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "compartment_ocid" {
}

variable "region" {
  default = "us-ashburn-1"
}

variable "kms_vault_id" {
}

variable "compartment_id" {
}

variable "cluster_cluster_pod_network_options_cni_type" {
  default = "OCI_VCN_IP_NATIVE"
}

variable "cluster_defined_tags_value" {
  default = "value"
}

variable "cluster_endpoint_config_is_public_ip_enabled" {
  default = false
}

variable "cluster_endpoint_config_nsg_ids" {
  default = []
}

variable "cluster_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "cluster_image_policy_config_is_policy_enabled" {
  default = false
}

variable "cluster_kubernetes_version" {
  default = "v1.30.1"
}

variable "cluster_name" {
  default = "oidc-discovery-example"
}

variable "cluster_options_add_ons_is_kubernetes_dashboard_enabled" {
  default = false
}

variable "cluster_options_add_ons_is_tiller_enabled" {
  default = false
}

variable "cluster_options_admission_controller_options_is_pod_security_policy_enabled" {
  default = false
}

variable "cluster_options_open_id_connect_discovery_is_open_id_connect_discovery_enabled" {
  default = false
}

variable "cluster_options_kubernetes_network_config_pods_cidr" {
  default = "10.1.0.0/16"
}

variable "cluster_options_kubernetes_network_config_services_cidr" {
  default = "10.2.0.0/16"
}

variable "cluster_options_persistent_volume_config_defined_tags_value" {
  default = "value"
}

variable "cluster_options_persistent_volume_config_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "cluster_options_service_lb_config_defined_tags_value" {
  default = "value"
}

variable "cluster_options_service_lb_config_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "cluster_options_service_lb_subnet_ids" {
  default = []
}

variable "cluster_state" {
  default = []
}

// enhanced cluster is required for OIDC Discovery to be enabled
variable "cluster_type" {
  default = "ENHANCED_CLUSTER"
}



provider "oci" {
  region           = var.region
  auth = "SecurityToken"
  config_file_profile = "terraform-federation-test"
}

variable defined_tag_namespace_name {
  default = "test"
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

resource "oci_identity_tag_namespace" "tag-namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description = "example tag namespace"
  name = var.defined_tag_namespace_name != "" ? var.defined_tag_namespace_name : "example-tag-namespace-all"

  is_retired = false
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

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}

data "oci_identity_availability_domain" "ad1" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

data "oci_identity_availability_domain" "ad2" {
  compartment_id = var.tenancy_ocid
  ad_number      = 2
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
  kubernetes_version = var.cluster_kubernetes_version
  name               = "tfTestCluster"
  vcn_id             = oci_core_vcn.test_vcn.id
  type               = var.cluster_type

  #Optional
  #   defined_tags = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.cluster_defined_tags_value)

  freeform_tags = var.cluster_freeform_tags

  options {

    #Optional
    add_ons {

      #Optional
      is_kubernetes_dashboard_enabled = var.cluster_options_add_ons_is_kubernetes_dashboard_enabled
      is_tiller_enabled               = var.cluster_options_add_ons_is_tiller_enabled
    }
    admission_controller_options {

      #Optional
      is_pod_security_policy_enabled = var.cluster_options_admission_controller_options_is_pod_security_policy_enabled
    }
    kubernetes_network_config {

      #Optional
      pods_cidr     = var.cluster_options_kubernetes_network_config_pods_cidr
      services_cidr = var.cluster_options_kubernetes_network_config_services_cidr
    }

    open_id_connect_discovery {
      #Optional
      is_open_id_connect_discovery_enabled = var.cluster_options_open_id_connect_discovery_is_open_id_connect_discovery_enabled
    }
  }
}

data "oci_containerengine_clusters" "test_clusters" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  name  = var.cluster_name
  state = var.cluster_state
}