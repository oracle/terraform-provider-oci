resource "oci_core_volume" "test_volume_with_required_parameter_vol" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id = var.compartment_ocid
  display_name = "test_volume_2"
  size_in_gbs  = "50"
}


# Example of querying volumes within a compartment (representing a potential previous query)
data "oci_core_volume" "test_volume_by_compartment_old" {
  volume_id = oci_core_volume.test_volume_with_required_parameter_vol.id
}
