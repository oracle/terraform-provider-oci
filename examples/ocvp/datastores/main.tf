// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

# VCN comes with default route table, security list and DHCP options

variable "tenancy_ocid" {
}

variable "compartment_ocid" {
}

variable "region" {
}

variable "config_file_profile" {}

provider "oci" {
  region              = var.region
  tenancy_ocid        = var.tenancy_ocid
  auth                = "SecurityToken"
  config_file_profile = var.config_file_profile
}

data "oci_identity_availability_domains" "ADs" {
  compartment_id = var.compartment_ocid
}

resource "oci_core_volume" "test_block_volume" {
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[1]["name"]
  compartment_id      = var.compartment_ocid
  display_name        = "TestBlockVolume"
  size_in_gbs         = "4096"
}

resource "oci_ocvp_datastore" "test_datastore" {
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[1]["name"]
  block_volume_ids = ["${oci_core_volume.test_block_volume.id}"]
  compartment_id = var.compartment_ocid
  display_name = "displayName"
}

resource "oci_ocvp_datastore_cluster" "test_datastore_cluster" {
  availability_domain = data.oci_identity_availability_domains.ADs.availability_domains[1]["name"]
  compartment_id = var.compartment_ocid
  datastore_cluster_type = "MANAGEMENT"
  datastore_ids = ["${oci_ocvp_datastore.test_datastore.id}"]
  display_name = "displayName"
}

data "oci_ocvp_datastore" "test_datastore_data" {
  datastore_id = oci_ocvp_datastore.test_datastore.id
}

data "oci_ocvp_datastore_cluster" "test_datastore_cluster_data" {
  datastore_cluster_id = oci_ocvp_datastore_cluster.test_datastore_cluster.id
}

data "oci_ocvp_datastores" "test_datastore_datastores" {
  #Optional
  compartment_id = var.compartment_ocid
  datastore_id = oci_ocvp_datastore.test_datastore.id
  display_name = oci_ocvp_datastore.test_datastore.display_name
  #cluster_id = oci_ocvp_cluster
  #state = var.datastore_state
}

data "oci_ocvp_datastore_clusters" "test_datastore_clusters_datastores" {
  #Optional
  compartment_id = var.compartment_ocid
  datastore_cluster_id = oci_ocvp_datastore_cluster.test_datastore_cluster.id
  display_name = oci_ocvp_datastore_cluster.test_datastore_cluster.display_name
  #cluster_id = oci_ocvp_cluster.test_cluster.id
  #state = var.datastore_cluster_state
}