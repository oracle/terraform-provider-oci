resource "oci_identity_compartment" "compartment" {
  provider    = "oci.home"
  name        = "${var.compartment_name}"
  description = "${var.compartment_description}"
}
