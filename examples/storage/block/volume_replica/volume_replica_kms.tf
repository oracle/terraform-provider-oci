
resource "oci_core_volume" "test_volume_with_required_parameter" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id = var.compartment_ocid
}

variable "kms_key_ocid_cross_region" {
  default = ""
}
resource "oci_core_volume" "test_volume_with_optional_parameter" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id = var.compartment_ocid
  display_name = "test_volume"

  // please find allowed other region's availability_domain and hardcode here
  block_volume_replicas {
    availability_domain = data.oci_identity_availability_domain.ad.name
    display_name = "test_replicas"
  }

  // if you want delete volume and this volume has replicas, please disable replicas at first, set this "block_volume_replicas_deletion" to true
   block_volume_replicas_deletion = true

}

output "volume" {
  value = {
    test_volume_with_required_parameter = oci_core_volume.test_volume_with_required_parameter.id
    test_volume_with_optional_parameter = oci_core_volume.test_volume_with_optional_parameter.id
  }
}
