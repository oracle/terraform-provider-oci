variable "tenancy_ocid" {
}

variable "auth" {
}

variable "config_file_profile" {

}

variable "region" {
}

variable "compartment_ocid" {
}

provider "oci" {
  auth                = var.auth
  config_file_profile = var.config_file_profile
  region              = var.region
  version = "6.35.0"
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

resource "oci_core_volume" "test_volume_with_required_parameter" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id = var.compartment_ocid
  size_in_gbs         = "50"
}

resource "oci_core_volume" "test_volume_with_optional_parameter" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id = var.compartment_ocid
  size_in_gbs         = "50"
  display_name = "test_volume"
}

output "volume" {
  value = {
    test_volume_with_required_parameter = oci_core_volume.test_volume_with_required_parameter.id
    test_volume_with_optional_parameter = oci_core_volume.test_volume_with_optional_parameter.id
  }
}


#Path - storage/block/corevolume/block_volume
