resource "oci_kms_key" "test_key" {
  #Required
  compartment_id      = "${var.compartment_id}"
  display_name        = "${var.key_display_name}"
  management_endpoint = "${data.oci_kms_vault.test_vault.management_endpoint}"

  key_shape {
    #Required
    algorithm = "${var.key_key_shape_algorithm}"
    length    = "${var.key_key_shape_length}"
  }
}

resource "oci_kms_key_version" "test_key_version" {
  #Required
  key_id              = "${oci_kms_key.test_key.id}"
  management_endpoint = "${data.oci_kms_vault.test_vault.management_endpoint}"
}
