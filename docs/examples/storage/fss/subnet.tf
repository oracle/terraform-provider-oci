resource "oci_core_subnet" "my_subnet" {
  availability_domain = "${var.availability_domain}"
  cidr_block = "${var.my_subnet_cidr}"
  display_name = "mysubnet"
  dns_label = "mysubnet"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${oci_core_virtual_network.my_vcn.id}"
  security_list_ids = ["${oci_core_security_list.my_security_list.id}"]
}