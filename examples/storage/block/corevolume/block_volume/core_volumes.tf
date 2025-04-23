resource "oci_core_volume" "test_volume_with_required_parameter_vg" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id = var.compartment_ocid
  display_name = "test_volume_1"
  size_in_gbs         = "50"
}

# Example of listing volumes in a compartment (representing a potential previous query)
data "oci_core_volumes" "test_volumes_in_compartment_old" {
  compartment_id = var.compartment_ocid
}

# Example of listing volumes with a specific display name (if your changes affected this)
data "oci_core_volumes" "test_volumes_by_display_name_old" {
  compartment_id = var.compartment_ocid
  filter {
    name   = "display_name"
    values = ["test_volume_1"]
  }
}
