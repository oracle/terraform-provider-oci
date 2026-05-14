provider "oci" {
  auth                = var.auth
  config_file_profile = var.config_file_profile
  region              = var.region
}

variable "auth" {}
variable "config_file_profile" {}
variable "region" {}
variable "compartment_ocid" {}
variable "tenancy_ocid" {}
variable "availability_domain" {}

variable "logical_placement_constraint" {
  default = "PACKED_DISTRIBUTION_MULTI_BLOCK"
}

data "oci_core_compute_capacity_topologies" "test_compute_capacity_topologies" {
  compartment_id      = var.tenancy_ocid
  availability_domain = var.availability_domain
}

data "oci_core_compute_capacity_topology_compute_hpc_islands" "test_compute_hpc_islands" {
  compute_capacity_topology_id = data.oci_core_compute_capacity_topologies.test_compute_capacity_topologies.compute_capacity_topology_collection[0].items[0].id
  compartment_id               = var.tenancy_ocid
  availability_domain          = var.availability_domain
}

data "oci_core_compute_capacity_topology_compute_network_blocks" "test_compute_network_blocks" {
  compute_capacity_topology_id = data.oci_core_compute_capacity_topologies.test_compute_capacity_topologies.compute_capacity_topology_collection[0].items[0].id
  compute_hpc_island_id        = data.oci_core_compute_capacity_topology_compute_hpc_islands.test_compute_hpc_islands.compute_hpc_island_collection[0].items[0].id
  compartment_id               = var.tenancy_ocid
  availability_domain          = var.availability_domain
}

resource "oci_core_compute_cluster" "test_compute_cluster" {
  availability_domain = var.availability_domain
  compartment_id      = var.compartment_ocid
  display_name        = "TestComputeClusterPlacementConstraint"

  placement_constraint_details {
    type                         = "COMPUTE_CLUSTER"
    hpc_island_id                = data.oci_core_compute_capacity_topology_compute_hpc_islands.test_compute_hpc_islands.compute_hpc_island_collection[0].items[0].id
    logical_placement_constraint = var.logical_placement_constraint
    target_network_block_ids     = [data.oci_core_compute_capacity_topology_compute_network_blocks.test_compute_network_blocks.compute_network_block_collection[0].items[0].id]
  }
}

data "oci_core_compute_cluster" "test_compute_cluster" {
  compute_cluster_id = oci_core_compute_cluster.test_compute_cluster.id
}

data "oci_core_compute_clusters" "test_compute_clusters" {
  compartment_id      = var.compartment_ocid
  availability_domain = var.availability_domain

  filter {
    name   = "id"
    values = [oci_core_compute_cluster.test_compute_cluster.id]
  }
}

output "compute_cluster" {
  value = oci_core_compute_cluster.test_compute_cluster
}

output "compute_cluster_data_source" {
  value = data.oci_core_compute_cluster.test_compute_cluster
}

output "compute_clusters_data_source" {
  value = data.oci_core_compute_clusters.test_compute_clusters
}