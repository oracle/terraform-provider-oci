variable "VPC-CIDR" {
  default = "10.0.0.0/16"
}

resource "oci_core_virtual_network" "cloudera_vcn" {
  cidr_block = "${var.VPC-CIDR}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "cloudera_vcn"
  dns_label = "cdhvcn"
}

resource "oci_core_internet_gateway" "cloudera_internet_gateway" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "cloudera_internet_gateway"
    vcn_id = "${oci_core_virtual_network.cloudera_vcn.id}"
}

resource "oci_core_route_table" "RouteForComplete" {
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.cloudera_vcn.id}"
    display_name = "RouteTableForComplete"
    route_rules {
        cidr_block = "0.0.0.0/0"
        network_entity_id = "${oci_core_internet_gateway.cloudera_internet_gateway.id}"
    }
}

resource "oci_core_security_list" "PublicSubnet" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "Public Subnet"
    vcn_id = "${oci_core_virtual_network.cloudera_vcn.id}"
    egress_security_rules = [{
        destination = "0.0.0.0/0"
        protocol = "6"
    }]
    ingress_security_rules = [{
        tcp_options {
            "max" = 7180
            "min" = 7180
        }
        protocol = "6"
        source = "0.0.0.0/0"
    }]
    ingress_security_rules = [{
        tcp_options {
            "max" = 8088
            "min" = 8088
        }
        protocol = "6"
        source = "0.0.0.0/0"
    }]
    ingress_security_rules = [{
        tcp_options {
            "max" = 19888
            "min" = 19888
        }
        protocol = "6"
        source = "0.0.0.0/0"
    }]
    ingress_security_rules = [{
        tcp_options {
            "max" = 22
            "min" = 22
        }
        protocol = "6"
        source = "0.0.0.0/0"
    }]
    ingress_security_rules = [{
        protocol = "6"
	source = "${var.VPC-CIDR}"
    }]


}

resource "oci_core_security_list" "PrivateSubnet" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "Private"
    vcn_id = "${oci_core_virtual_network.cloudera_vcn.id}"
    egress_security_rules = [{
        destination = "0.0.0.0/0"
        protocol = "6"
    }]
    egress_security_rules = [{
	protocol = "6"
	destination = "${var.VPC-CIDR}"
    }]
    ingress_security_rules = [{
        protocol = "6"
        source = "${var.VPC-CIDR}"
    }]
}

resource "oci_core_security_list" "BastionSubnet" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "Bastion"
    vcn_id = "${oci_core_virtual_network.cloudera_vcn.id}"
    egress_security_rules = [{
	protocol = "6"
        destination = "0.0.0.0/0"
    }]
    ingress_security_rules = [{
        tcp_options {
            "max" = 22
            "min" = 22
        }
        protocol = "6"
        source = "0.0.0.0/0"
    },
	{
	protocol = "6"
        source = "${var.VPC-CIDR}"
    }]	
}

## Publicly Accessable Subnet Setup

resource "oci_core_subnet" "public" {
  count = "3"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[count.index],"name")}"
  cidr_block = "${cidrsubnet(var.VPC-CIDR, 8, count.index)}"
  display_name = "public_${count.index}"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${oci_core_virtual_network.cloudera_vcn.id}"
  route_table_id = "${oci_core_route_table.RouteForComplete.id}"
  security_list_ids = ["${oci_core_security_list.PublicSubnet.id}"]
  dhcp_options_id = "${oci_core_virtual_network.cloudera_vcn.default_dhcp_options_id}"
  dns_label = "public${count.index}"
}

## Private Subnet Setup

resource "oci_core_subnet" "private" {
  count = "3"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[count.index],"name")}"
  cidr_block = "${cidrsubnet(var.VPC-CIDR, 8, count.index+3)}"
  display_name = "private_ad${count.index}"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${oci_core_virtual_network.cloudera_vcn.id}"
  route_table_id = "${oci_core_route_table.RouteForComplete.id}"
  security_list_ids = ["${oci_core_security_list.PrivateSubnet.id}"]
  dhcp_options_id = "${oci_core_virtual_network.cloudera_vcn.default_dhcp_options_id}"
  #prohibit_public_ip_on_vnic = "true"
  dns_label = "private${count.index}"
}
## Bastion Subnet Setup

resource "oci_core_subnet" "bastion" {
  count = "3"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[count.index],"name")}"
  cidr_block = "${cidrsubnet(var.VPC-CIDR, 8, count.index+6)}" 
  display_name = "bastion_ad${count.index}"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${oci_core_virtual_network.cloudera_vcn.id}"
  route_table_id = "${oci_core_route_table.RouteForComplete.id}"
  security_list_ids = ["${oci_core_security_list.BastionSubnet.id}"]
  dhcp_options_id = "${oci_core_virtual_network.cloudera_vcn.default_dhcp_options_id}"
  dns_label = "bastion${count.index}"
}
