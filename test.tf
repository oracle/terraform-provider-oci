provider "baremetal" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  private_key_password = "${var.private_key_password}"
  private_key = ""
}

resource "baremetal_identity_compartment" "test-dec5" {
  compartment_id = "${var.tenancy_ocid}-dec5"
  name = "test_compartment_dec5"
  description = "A special test compartment."
}
