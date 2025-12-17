
variable "tenancy_ocid" {
}

variable "region" {
}

variable "compartment_ocid" {
}

variable "ssh_public_key" {
}

variable "config_file_profile" {
}

variable "shape" {
}

# provider "oci" {
#   region              = var.region
#   auth                = "SecurityToken"
#   config_file_profile = var.config_file_profile
#   version             = "7.20.0"
#  }

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}

data "oci_core_shapes" "test_shapes" {
  compartment_id = var.compartment_ocid
  shape = var.shape
}

output "oci_core_shapes" {
  value = data.oci_core_shapes.test_shapes.*
}
