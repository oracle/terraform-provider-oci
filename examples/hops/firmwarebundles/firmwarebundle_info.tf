
variable "tenancy_ocid" {
}

variable "ssh_public_key" {
}

variable "region" {
  default = "r1"
}

variable "firmware_bundle_id" {
}

data "oci_core_firmware_bundle" "test_firmware_bundle" {
  firmware_bundle_id = var.firmware_bundle_id
}

data "oci_core_firmware_bundles" "test_firmware_bundles" {
  platform = data.oci_core_firmware_bundle.test_firmware_bundle.platforms.0.platform
}

output "firmware_bundle_id_info" {
  value = data.oci_core_firmware_bundle.test_firmware_bundle.*
}

output "firmware_bundles_info" {
  value = data.oci_core_firmware_bundles.test_firmware_bundles.*
}