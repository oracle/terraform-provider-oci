provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}

data "oci_core_compute_hosts" "test_compute_hosts" {
  compartment_id = "${var.compartment_ocid}"
  compute_host_in_subtree = true
}

locals {
  core_compute_hosts_ids = [
    for host in data.oci_core_compute_hosts.test_compute_hosts.compute_host_collection :
    [ for item in host.items : item.id ]
  ]
}

data "oci_core_compute_host" "test_compute_host" {
  for_each = toset(flatten(local.core_compute_hosts_ids))

  compute_host_id = each.key
}

resource "oci_core_compute_host_group" "best_compute_host_group" {
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains[0].name
  compartment_id = var.compartment_ocid
  display_name = "BestComputeHostGroup"
  is_targeted_placement_required = false
}

resource "oci_core_compute_host" "best_compute_host" {
  compute_host_id = keys(data.oci_core_compute_host.test_compute_host)[0]
  compute_host_group_id = oci_core_compute_host_group.best_compute_host_group.id
}

output "compute_host_values" {
  value = {
    for key, value in data.oci_core_compute_host.test_compute_host :
    key => {
      name = value.display_name
      fd = value.fault_domain
    }
  }
}


output "compute_hosts_list" {
  value = data.oci_core_compute_hosts.test_compute_hosts
}

output "compute_host_resource" {
  value = oci_core_compute_host.best_compute_host
}