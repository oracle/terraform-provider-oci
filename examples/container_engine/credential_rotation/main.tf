// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "region" {
  default = "us-ashburn-1"
}
variable "compartment_ocid" {}

variable "cluster_start_credential_rotation_management_auto_completion_delay_duration" {
  default = "P5D"
}

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

resource "oci_containerengine_cluster" "test_cluster" {
  #Required
  compartment_id     = var.compartment_ocid
  kubernetes_version = reverse(data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions)[0]
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
      is_pod_security_policy_enabled = false
    }

    kubernetes_network_config {
      #Optional
      pods_cidr     = "10.1.0.0/16"
      services_cidr = "10.2.0.0/16"
    }
  }
}

// start credential rotation on a cluster
resource "oci_containerengine_cluster_start_credential_rotation_management" "test_cluster_start_credential_rotation_management" {
  #Required
  auto_completion_delay_duration = var.cluster_start_credential_rotation_management_auto_completion_delay_duration
  cluster_id                     = oci_containerengine_cluster.test_cluster.id
}

// get credential rotation status
data "oci_containerengine_cluster_credential_rotation_status" "test_cluster_credential_rotation_status" {
  #Required
  cluster_id = oci_containerengine_cluster.test_cluster.id
}

// complete credential rotation on a cluster
resource "oci_containerengine_cluster_complete_credential_rotation_management" "test_cluster_complete_credential_rotation_management" {
  #Required
  cluster_id = oci_containerengine_cluster.test_cluster.id
  depends_on = [oci_containerengine_cluster_start_credential_rotation_management.test_cluster_start_credential_rotation_management]
}

