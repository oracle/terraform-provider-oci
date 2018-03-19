resource "oci_core_virtual_network" "my_vcn" {
  cidr_block = "${var.my_vcn-cidr}"
  dns_label = "myvcn"
  compartment_id = "${var.compartment_ocid}"
  display_name = "myvcn"
  dns_label = "myvcn"
}
