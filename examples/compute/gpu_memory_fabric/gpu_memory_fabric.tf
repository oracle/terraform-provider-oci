provider "oci" {
  auth                = var.auth
  config_file_profile = var.config_file_profile
  region              = var.region
}

variable "auth" {}
variable "region" {}
variable "config_file_profile" {}
variable "compartment_ocid" {}
variable "tenancy_ocid" {}

variable "compute_gpu_memory_fabric_compute_gpu_memory_fabric_health" {
  default = "HEALTHY"
}

variable "compute_gpu_memory_fabric_compute_gpu_memory_fabric_lifecycle_state" {
  default = "AVAILABLE"
}

variable "compute_gpu_memory_fabric_freeform_tags" {
  default = { "Department" = "Internal" }
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

output "ad" {
  value = data.oci_identity_availability_domain.ad
}

data "oci_core_compute_gpu_memory_fabrics" "gpu_memory_fabrics" {
  compartment_id = var.tenancy_ocid
  availability_domain = data.oci_identity_availability_domain.ad.name
  compute_gpu_memory_fabric_health = "HEALTHY"
  compute_gpu_memory_fabric_lifecycle_state = "AVAILABLE"
}

output "list_gpu_memory_fabrics" {
  value = data.oci_core_compute_gpu_memory_fabrics.gpu_memory_fabrics
}

data "oci_core_compute_gpu_memory_fabric" "gpu_memory_fabric" {
  compute_gpu_memory_fabric_id = data.oci_core_compute_gpu_memory_fabrics.gpu_memory_fabrics.compute_gpu_memory_fabric_collection[0].items[0].id
}

output "get_gpu_memory_fabric" {
  value = data.oci_core_compute_gpu_memory_fabric.gpu_memory_fabric
}

resource "oci_core_compute_gpu_memory_fabric" "gpu_memory_fabric" {
  compute_gpu_memory_fabric_id = data.oci_core_compute_gpu_memory_fabrics.gpu_memory_fabrics.compute_gpu_memory_fabric_collection[0].items[0].id
  freeform_tags = var.compute_gpu_memory_fabric_freeform_tags
}

output "gpu_memory_fabric" {
  value = oci_core_compute_gpu_memory_fabric.gpu_memory_fabric
}
