provider "oci" {
  auth                = var.auth
  config_file_profile = var.config_file_profile
  region              = var.region
}

variable "auth" {}
variable "config_file_profile" {}
variable "region" {}
variable "compartment_ocid" {}
variable "availability_domain" {}

resource "oci_core_compute_cluster" "test_compute_cluster" {
  availability_domain = var.availability_domain
  compartment_id      = var.compartment_ocid
  display_name        = "TestComputeClusterBasic"
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

output "compute_cluster_id" {
  value = oci_core_compute_cluster.test_compute_cluster.id
}

output "compute_cluster_data_source_id" {
  value = data.oci_core_compute_cluster.test_compute_cluster.id
}

output "compute_clusters_count" {
  value = length(data.oci_core_compute_clusters.test_compute_clusters.compute_cluster_collection[0].items)
}
