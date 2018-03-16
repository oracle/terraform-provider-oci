resource "oci_core_virtual_network" "vcn" {
  cidr_block     = "${var.vcn_cidr_block}"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "${var.vcn_dns_name}.vcn"
  dns_label      = "${var.vcn_dns_name}"
}

resource "oci_core_internet_gateway" "ig" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "${var.label_prefix}${var.vcn_dns_name}.ig"
  vcn_id         = "${oci_core_virtual_network.vcn.id}"
}

resource "oci_core_route_table" "rt" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "${var.vcn_dns_name}.rt"
  vcn_id         = "${oci_core_virtual_network.vcn.id}"

  route_rules {
    cidr_block        = "0.0.0.0/0"
    network_entity_id = "${oci_core_internet_gateway.ig.id}"
  }
}

resource "oci_core_dhcp_options" "custom_dhcp_options" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "${var.vcn_dns_name}-custom.dhcp"

  options {
    type               = "DomainNameServer"
    server_type        = "CustomDnsServer"
    custom_dns_servers = ["${var.local_dns_server}"]
  }

  vcn_id = "${oci_core_virtual_network.vcn.id}"
}

resource "oci_core_dhcp_options" "internet_dhcp_options" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "${var.vcn_dns_name}-internet.dhcp"

  options {
    type        = "DomainNameServer"
    server_type = "VcnLocalPlusInternet"
  }

  vcn_id = "${oci_core_virtual_network.vcn.id}"
}
