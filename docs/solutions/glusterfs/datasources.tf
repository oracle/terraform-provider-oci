data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.compartment_ocid}"
}
