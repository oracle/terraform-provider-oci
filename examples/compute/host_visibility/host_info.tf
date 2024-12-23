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

data "oci_core_compute_hosts" "test_compute_hosts" {
  compartment_id = "${var.compartment_ocid}"
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

output "compute_host_values" {
  value = {
    for key, value in data.oci_core_compute_host.test_compute_host :
    key => {
      name = value.display_name
      fd = value.fault_domain
    }
  }
}