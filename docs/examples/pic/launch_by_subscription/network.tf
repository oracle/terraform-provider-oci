resource "oci_core_virtual_network" "pic_vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "PICVcn"
  dns_label      = "PICVcn"
}

resource "oci_core_subnet" "pic_subnet" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  cidr_block          = "10.1.20.0/24"
  display_name        = "PICSubnet"
  dns_label           = "PICSubnet"
  security_list_ids   = ["${oci_core_virtual_network.pic_vcn.default_security_list_id}"]
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_virtual_network.pic_vcn.id}"
  route_table_id      = "${oci_core_route_table.pic_rt.id}"
  dhcp_options_id     = "${oci_core_virtual_network.pic_vcn.default_dhcp_options_id}"
}

resource "oci_core_internet_gateway" "pic_ig" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "PICIG"
  vcn_id         = "${oci_core_virtual_network.pic_vcn.id}"
}

resource "oci_core_route_table" "pic_rt" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_virtual_network.pic_vcn.id}"
  display_name   = "PICRT"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_internet_gateway.pic_ig.id}"
  }
}
