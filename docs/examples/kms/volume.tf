resource "oci_core_volume" "my_volume" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  compartment_id      = "${var.compartment_id}"
  display_name        = "-tf-volume"
  size_in_gbs         = "${var.volume_size}"
  kms_key_id          = "${oci_kms_key.test_key.id}"
}
