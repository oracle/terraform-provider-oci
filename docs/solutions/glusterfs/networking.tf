resource "oci_core_virtual_network" "CompleteVCN" {
  cidr_block = "10.0.0.0/16"
  compartment_id = "${var.compartment_ocid}"
  dns_label = "${var.DnsLabel}"
  display_name = "CompleteVCN"
}

resource "oci_core_internet_gateway" "CompleteIG" {
  compartment_id = "${var.compartment_ocid}"
  display_name = "CompleteIG"
  vcn_id = "${oci_core_virtual_network.CompleteVCN.id}"
}

resource "oci_core_route_table" "RouteForComplete" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${oci_core_virtual_network.CompleteVCN.id}"
  display_name = "RouteTableForComplete"
  route_rules {
    cidr_block = "0.0.0.0/0"
    network_entity_id = "${oci_core_internet_gateway.CompleteIG.id}"
  }
}

resource "oci_core_security_list" "Subnet" {
  compartment_id = "${var.compartment_ocid}"
  display_name = "Subnet"
  vcn_id = "${oci_core_virtual_network.CompleteVCN.id}"
  egress_security_rules = [{
    protocol = "all"
    destination = "0.0.0.0/0"
  }]
  ingress_security_rules = [{
    protocol = "all"
    source = "0.0.0.0/0"
  }]
}

resource "oci_core_subnet" "SubnetAD1" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
  cidr_block = "10.0.4.0/24"
  display_name = "SubnetAD1"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${oci_core_virtual_network.CompleteVCN.id}"
  route_table_id = "${oci_core_route_table.RouteForComplete.id}"
  security_list_ids = ["${oci_core_security_list.Subnet.id}"]
  dhcp_options_id = "${oci_core_virtual_network.CompleteVCN.default_dhcp_options_id}"
  dns_label = "publicsubnetad1"
}

resource "oci_core_subnet" "SubnetAD2" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[1],"name")}"
  cidr_block = "10.0.5.0/24"
  display_name = "SubnetAD2"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${oci_core_virtual_network.CompleteVCN.id}"
  route_table_id = "${oci_core_route_table.RouteForComplete.id}"
  security_list_ids = ["${oci_core_security_list.Subnet.id}"]
  dhcp_options_id = "${oci_core_virtual_network.CompleteVCN.default_dhcp_options_id}"
  dns_label = "publicsubnetad2"
}

resource "oci_core_subnet" "SubnetAD3" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[2],"name")}"
  cidr_block = "10.0.6.0/24"
  display_name = "SubnetAD3"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${oci_core_virtual_network.CompleteVCN.id}"
  route_table_id = "${oci_core_route_table.RouteForComplete.id}"
  security_list_ids = ["${oci_core_security_list.Subnet.id}"]
  dhcp_options_id = "${oci_core_virtual_network.CompleteVCN.default_dhcp_options_id}"
  dns_label = "publicsubnetad3"
}

