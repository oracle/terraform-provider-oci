
variable "source_availability_domain" {
  default = ""
}

resource "oci_core_volume" "test_volume_with_required_parameter" {
  availability_domain = var.source_availability_domain
  compartment_id      = var.compartment_ocid
  size_in_gbs         = "50"
}

variable "kms_key_ocid_cross_region" {
  default = ""
}

variable "replica_availability_domain" {
  default = ""
}

resource "oci_core_volume" "test_volume_with_optional_parameter" {
  availability_domain = var.source_availability_domain
  compartment_id      = var.compartment_ocid
  display_name        = "test_volume"
  size_in_gbs         = "50"

  // please find allowed other region's availability_domain and hardcode here
  block_volume_replicas {
    availability_domain = var.replica_availability_domain
    display_name        = "test_replicas"
    xrr_kms_key_id      = var.kms_key_ocid_cross_region
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
