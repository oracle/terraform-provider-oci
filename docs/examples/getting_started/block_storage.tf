# This provisions a block storage volume which can later be used by the instances that we create.
resource "oci_core_volume" "MyBlockStorage" {
  availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "MyBlockStorage"
  size_in_gbs = "${var.BlockVolumeSize}"
}
