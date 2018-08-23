resource "oci_core_virtual_network" "my_vcn" {
  cidr_block     = "${var.my_vcn-cidr}"
  dns_label      = "myvcn"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "myvcn"
  dns_label      = "myvcn"
}

resource "oci_core_internet_gateway" "my_internet_gateway" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "my internet gateway"
  vcn_id         = "${oci_core_virtual_network.my_vcn.id}"
}

resource "oci_core_route_table" "my_route_table" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_virtual_network.my_vcn.id}"
  display_name   = "my route table"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_internet_gateway.my_internet_gateway.id}"
  }
}

resource "oci_core_subnet" "my_subnet" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  cidr_block          = "${var.my_subnet_cidr}"
  display_name        = "mysubnet"
  dns_label           = "mysubnet"
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_virtual_network.my_vcn.id}"
  security_list_ids   = ["${oci_core_security_list.my_security_list.id}"]
  route_table_id      = "${oci_core_route_table.my_route_table.id}"
}
