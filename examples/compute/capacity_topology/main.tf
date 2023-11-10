variable "profile" {
}

variable "region" {
}

variable "availability_domain" {
}

variable "tenant_id" {
}

provider "oci" {
  auth = "SecurityToken"
  config_file_profile = var.profile
  region = var.region
}

resource "oci_core_compute_capacity_topology" "test_compute_capacity_topology" {
  availability_domain = var.availability_domain
  capacity_source {
    capacity_type = "DEDICATED"
    compartment_id = var.tenant_id
  }
  compartment_id = var.tenant_id
  display_name = "test-dedicated-capacity-topology"
  freeform_tags = {
    "department" = "Finance"
  }
}

output "output_compute_capacity_topology" {
  value = oci_core_compute_capacity_topology.test_compute_capacity_topology
}

data "oci_core_compute_capacity_topologies" "test_compute_capacity_topologies" {
  compartment_id = var.tenant_id
}

output "output_compute_capacity_topologies" {
  value = data.oci_core_compute_capacity_topologies.test_compute_capacity_topologies
}

data "oci_core_compute_capacity_topology_compute_bare_metal_hosts" "test_compute_bare_metal_hosts" {
  compute_capacity_topology_id = oci_core_compute_capacity_topology.test_compute_capacity_topology.id
}

output "output_compute_bare_metal_hosts" {
  value = data.oci_core_compute_capacity_topology_compute_bare_metal_hosts.test_compute_bare_metal_hosts
}

data "oci_core_compute_capacity_topology_compute_hpc_islands" "test_compute_hpc_islands" {
  compute_capacity_topology_id = oci_core_compute_capacity_topology.test_compute_capacity_topology.id
}

output "output_compute_hpc_islands" {
  value = data.oci_core_compute_capacity_topology_compute_hpc_islands.test_compute_hpc_islands
}

data "oci_core_compute_capacity_topology_compute_network_blocks" "test_compute_network_blocks" {
  compute_capacity_topology_id = oci_core_compute_capacity_topology.test_compute_capacity_topology.id
}

output "output_compute_network_blocks" {
  value = data.oci_core_compute_capacity_topology_compute_network_blocks.test_compute_network_blocks
}

